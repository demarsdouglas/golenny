# golenny ( ͡° ͜ʖ ͡°)

Heavily inspired by [RESTful-lenny](https://github.com/LennyToday/RESTful-lenny), golenny aims to deliver similar Lenny
functionality in a module form, with the end goal being improved performance and less reliance on external systems by
installing this module directly into your project dependencies.

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

```
( ͡° ͜ʖ ͡°)
```

`LoadDefault()` will load the old-fashioned lenny face struct into the var, and calling `.Render()` from
that type it will render the face into a string format.

It is of note that sometimes a Lenny will _not_ look as good in monospaced fonts as he does in other spaces (emails,
messages, forum posts, documents, etc).

You may also want to generate a randomized Lenny, and this can be done in the same way:

```go
func main() {
    l := lenny.LoadRandom()
    fmt.Println(l.Render())
}
```

Outputs:

```
ᕕ(ᗒ෴ᗕ)ᕗ
```

These Lennys will be entirely randomized off a few preset lists of possible eyes, mouths, and ends. Some of them are
fantastic, like the one above. Some of them are absolute abominations. They do _not_ come with a `Mood` or `name`
attribute in their struct, but you're more than able to assign those values if needed.

We have three other functions you can call as well; `LoadPredefinedByName(name string)`, `LoadRandomPredefined()`,
and `LoadRandomPredefinedByMood(mood Mood)`. A list of named Lennys is coming soon if you don't want to peruse the
source code. The `Mood`s you're allowed to pass into the latter are `Happy`, `Sad`, `Strange`, `Angry`, and `Surprised`.

```go
func main() {
    l1 := lenny.LoadPredefinedByName("Lenny")
    l2 := lenny.LoadRandomPredefinedByMood(lenny.Strange)
    l3 := lenny.LoadRandomPredefined()
    
    fmt.Println(l1.Render())
    fmt.Println(l2.Render())
    fmt.Println(l3.Render())
}
```

Outputs:

```
( ͡° ͜ʖ ͡°)
ᕦ⊙෴⊙ᕤ
(◞д◟)
```

## TO-DOs:

This project is far from complete!

- I need a larger list of predefined Lennys
- "Getting started" and related documentation
