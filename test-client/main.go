package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"go-fft-client/shared"
	"io"
	"io/ioutil"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/zenwerk/go-wave"
)

func main() {
	ch := make(chan *shared.Response, 1)

	config := NewConfig()
	logrus.SetLevel(logrus.InfoLevel)

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
	})

	c := shared.NewMQTTClient(config.MQTT)
	c.Subscribe(config.MQTT.ResultTopic, func(bytes []byte) {
		r, err := shared.FromJson(bytes)
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.Infof("Triggered: %t", r.Trigger)
		ch <- r
	})

	files, err := ioutil.ReadDir("./sets")

	if err != nil {
		logrus.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		logrus.Infof("Using %s sets", f.Name())

		good, err := getSample(fmt.Sprintf("./sets/%s/good.wav", f.Name()))

		if err != nil {
			logrus.Fatal(err)
		}

		bad, err := getSample(fmt.Sprintf("./sets/%s/bad.wav", f.Name()))

		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Info("Sending initial good")
		err = sendSample(c, good, false, ch)
		if err != nil {
			logrus.Error(err)
			continue
		}

		logrus.Info("Sending second good")

		err = sendSample(c, good, false, ch)
		if err != nil {
			logrus.Error(err)
			continue
		}

		logrus.Info("Sending first bad")

		err = sendSample(c, bad, true, ch)
		if err != nil {
			logrus.Error(err)
			continue
		}

		logrus.Info("Sending second bad")

		err = sendSample(c, bad, false, ch)
		if err != nil {
			logrus.Error(err)
			continue
		}

		logrus.Info("Sending last good")

		err = sendSample(c, good, true, ch)
		if err != nil {
			logrus.Error(err)
			continue
		}

		logrus.Info("All passed")
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	for range sig {
		c.Close()
		os.Exit(0)
	}
}

func getSample(fName string) ([]byte, error) {
	goodReader, err := wave.NewReader(fName)

	if err != nil {
		return nil, err
	}

	buf := make([]byte, 0)

	for {
		b, err := goodReader.ReadSample()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		buffer := make([]byte, 4*len(b))

		for i, x := range b {
			binary.BigEndian.PutUint32(buffer[i*4:], math.Float32bits(float32(x)))
		}

		buf = append(buf, buffer...)
	}

	return buf, nil
}

func sendSample(client shared.MQTTClient, sample []byte, expectedResult bool, ch chan *shared.Response) error {
	err := client.SendSample(sample)
	if err != nil {
		logrus.Fatal(err)
	}

	select {
	case r := <-ch:
		if r.Trigger != expectedResult {
			return errors.New("got unexpected trigger response")
		}
		break

	case <-time.After(5 * time.Second):
		return errors.New("got timeout while waiting on response")
	}

	return nil
}
