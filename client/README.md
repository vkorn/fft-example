Client is responsible for capturing a short audio sample using [PortAudio](http://www.portaudio.com) and sending it through MQTT topic to analyzer service.

### Configuration params

* `log_level` -- [logrus](https://github.com/sirupsen/logrus) log level.
* `sample_rate` -- audio capturing frame rate. If you're not sure which rate to use, start with `debug` log level and check console output.
* `record_frame` -- number of seconds to record.


#### Dev notes for Mac
1. Install pkg-config

```bash
brew install pkg-config 
```

Make sure `/usr/local/bin` is in the `$PATH` otherwise override env variable `PKG_CONFIG=/usr/local/bin/pkg-config`.

2. Install PortAudio

```bash
brew install portaudio
```
