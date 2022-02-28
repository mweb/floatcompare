[![Go Report Card](https://goreportcard.com/badge/github.com/mweb/floatcompare)](https://goreportcard.com/report/github.com/mweb/floatcompare)
[![Build Status](https://github.com/mweb/floatcompare/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/mweb/floatcompare/actions)
[![Coverage Status](https://coveralls.io/repos/github/mweb/floatcompare/badge.svg?branch=main)](https://coveralls.io/github/mweb/floatcompare?branch=main)
[![License](https://img.shields.io/github/license/mweb/floatcompare)](/LICENSE)

# floatcompare

Go linter to search for float comparison.

## Why are float comparisons problematic?

The [Floating-Point Guide](https://floating-point-gui.de) has a good
[description](https://floating-point-gui.de/errors/comparison/) why comparison
of floats is problematic.

In short, since floating-point numbers do have some issues with rounding, a
comparison of float might not return the expected result. Therefore, I wrote
this linter which helps to find possible bugs with floating point number
comparison.

## Example

    package main

    import "fmt"

    func main() {
        a := 0.15
        b := 0.15

        c := 0.2
        d := 0.1

        // comparison of this two calculation, we would expect them to be true
        if (a + b) == (c + d) {
            fmt.Println("I would expect this is true")
        } else {
            fmt.Println("But it is not.")
        }

        // Or we would expect this to be true
        if (a + b) >= (c + d) {
            fmt.Println("I would expect this is true")
        } else {
            fmt.Println("But it is not.")
        }

        // But I would not expect this to be true
        if (a + b) < (c + d) {
            fmt.Println("I would not expect this is true")
        }
    }

This example would have the following output:

    But it is not.
    But it is not.
    I would not expect this is true

## What it detected?

This linter detects all comparisons of floats within the checked code. This
includes switch case statements.

It provides flags to reduce the test to equal and not equal comparison and to
omit test files.

## How to use the linter

### Install

    go install github.com/mweb/floatcompare/cmd/floatcompare@latest

### Usage

Call this within your project.

    floatcompare ./...

There are two parameters available to disable certain checks:

    --skipTests (to skip all *_test.go files)
    --equalOnly (only warn for == and != comparison)

## Credits

This tool is built on top of the excellent go/analysis package that makes it
easy to write custom analyzers in Go. And the blog post "[Using go/analsysis to
write a custom linter](https://arslan.io/2019/06/13/using-go-analysis-to-write-a-custom-linter/)"
helped me to get started with this linter.


Thanks to [Fatih Arslan](https://github.com/fatih)
