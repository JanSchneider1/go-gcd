# About

A small command line tool written in go to calculate the
'greatest common divisor'.

- [Installation](#Installation)
- [Usage](#Usage)

# Installation

```shell script
$ go get github.com/JanSchneider1/go-gcd
```

# Usage

```shell_script
$ go-gcd <options> <integer> <integer>
```

## Options

`-s` prints each step of calculation

## Examples

```shell script
$ go-gcd 1 1
--> 1

$ go-gcd 128 68
--> 4

$ go-gcd 68 128
--> 4

$ go-gcd 68 0
--> not defined for 0

$ go-gcd 68 -128
--> not defined for negative numbers

$ go-gcd -s 100 51
--> 100   mod    51 =    49
    51    mod    49 =     2
    49    mod     2 =     1

```
