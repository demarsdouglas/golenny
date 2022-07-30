# golenny ( ͡° ͜ʖ ͡°)

Heavily inspired by [RESTful-lenny](https://github.com/LennyToday/RESTful-lenny),
golenny aims to deliver similar Lenny functionality in a module form, with the end goal
being improved performance and less reliance on external systems by installing this module
directly into your project dependencies.

## Install

In a terminal:
```sh
$ cd <project path>
$ go get github.com/demarsdouglas/golenny
```

## Example use

In your project somewhere:
```go
func main() {
    l := lenny.LoadDefault()
    fmt.Println(l.Render())
}
```

Outputs:
```go
( ͡° ͜ʖ ͡°)
```

`LoadDefault()` will load the old-fashioned lenny face struct into the var, and calling `.Render()` from
that type it will render the face into a string format.

It is of note that Lenny does _not_ look as good in monospaced fonts as he does in other
spaces (emails, messages, forum posts, documents, etc).

## TO-DOs:

This project is far from complete!

- I need a large list of predefined Lennys
- I need a large list of potential eyes, ends, and mouths
- I want to be able to return a predefined Lenny by its name
- I want to be able to return a random predefined Lenny
- I want to be able to return a new randomized Lenny
- I want to be able to return a Lenny based off its mood
- "Getting started" and related documentation
