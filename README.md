# About

A small command line tool written in go to calculate the
greatest command divisor.

# Usage

```shell_script
$ go-gcd <integer> <integer>
```

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
```