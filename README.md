# Dummy log generator

A simple go program to generate dummy logs for apache or nginx.

## Install

```sh
go build -o dummy-log-generator .
```

## Help

```sh
./dummy-log-generator --help
```

## Run

```sh
./dummy-log-generator --format apache --path apache.log --interval 1s
```

or

```sh
./dummy-log-generator --format nginx --path nginx.log --interval 1s
```

This creates a log in the current path, and continues to create logs until you exist the program.

## Warning

If you just let this program continue forever, your log file will get gigantic.
