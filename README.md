# floatcompare

go linter to search for float comparison.

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

TODO
