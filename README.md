# fft-example

This repo contains an example of FFT processing. It consists of two parts: client and server. 

Client is responsible for the audio capturing, server -- for its analyzing. 

MQTT broker used for data transfer. 

Both applications share the same MQTT config. All other configuration options described in the corresponding readme. 

### Configuration params

* `broker` -- MQTT broker location.
* `client` -- client ID to use. Depending on broker configuration, you'll need to use different IDs.
* `username` -- login to use.
* `password` -- password to use.
* `topic` -- topic name for the audio samples. 
* `result_topic` -- topic name for trigger alerts. If left empty, server will output to console only.
* `qos` -- quality of service to use with messages.

