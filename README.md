# Dummy log generator

A simple go program to generate dummy logs for apache or nginx.
And yes, the output is the same if you use apache or nginx, both
are using the Combined log format. The idea here is to maybe add more
options later on. With variations for apache and nginx. So for now,
it doesn't make any difference if you use nginx or apache format.

## Use case?

If you need some logs for testing functionality of reading real time logs,
are just for having a log file to experiment with.

## Install

```sh
go build -o dummy-log-generator .
```

## Help

```sh
./dummy-log-generator --help
```

## Run

With default settings:

```sh
./dummy-log-generator
```

Or with override:

```sh
./dummy-log-generator --format apache --path apache.log --interval 1s
```

This creates a log in the current path, and continues to create logs until you exist the program.

## Warning

If you just let this program continue forever, your log file will get gigantic.

## License

Copyright (C) 2024 Digitalist Open Cloud <cloud@digitalist.com>

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program.  If not, see <https://www.gnu.org/licenses/>
