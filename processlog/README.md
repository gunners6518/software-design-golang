## processlog command

By writing it before the program to be executed, the status code and processing time will be output to standard output. It also supports writing log files to a folder with command line options.

## sample
```zsh
// オプションなし
$ ./processlog ./someprogram

exit status 0
wall clock time=4.030729012s system time =6.777ms user time=7.529ms


// オプションあり(-logdir)
// ./tmp/logにログファイル書き出し
$ ./processlog -logdir ./tmp/log ./someprogram

```
