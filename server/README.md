Server is responsible for analyzing received audio sample and sending a trigger notification if thresholds were met. 

### Configuration params

* `log_level` - [logrus](https://github.com/sirupsen/logrus) log level.
* `sample_rate` - frame rate used for received audio sample.
* `nfft` - number of data points for FFT.
* `peaks_limit` -- peaks limit.
* `peaks_threshold` -- peaks threshold.
* `freqs_threshold` -- frequencies threshold.
