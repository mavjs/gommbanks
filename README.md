About
=====
A simple golang cli app to show exchange rates from Myanmar bank(s).

Supported banks:
* [Central Bank of Myanmar](http://forex.cbm.gov.mm/)

Compiling
=========
first get the dependency
```
go get github.com/codegangsta/cli
```

then build the codez
```
go build
```

run
```
NAME:
   gommbanks - Get latest exchange rates from different banks in Myanmar.

USAGE:
   gommbanks [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR(S): 
   Ye Myat Kaung (Maverick) <mavjs01@gmail.com> 
   
COMMANDS:
   bank, b  bank to get exchange rate from. (currently supported CBM)
   help, h  Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h       show help
   --version, -v    print the version
```

TODO
====
* auto calculate based on input
```
    % ./gommbanks calculate 100000 MMK

    USD: 96.81
    EUR: 89.09
```
* add more support for other banks
