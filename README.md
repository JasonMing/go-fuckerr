# go-fuckerr

It is annoying that `err != nil` checks are everywhere in your Go code
when you don't need to know about what exactly it is, and just want to
propagate them to caller.

## Quick Start

```go
package main

import "github.com/jasonming/go-fuckerr/try"

func ExposedFunc() (string, error) {
    return try.Catching(func() string {
        return internalCall() // no more explicit err != nil
    })
}

func internalCall() string {
    if !satisfied {
        // just throw it
        try.Throw("illegal argument")
    }

    // r, err := somelib.Call()
    // if err != nil {
    //     return "", err
    // }

    // Simplify error check by Must
    r := try.Must(somelib.Call())

    // Or，propagate error after partial check
    r, err := somelib.Call()
    switch err {
    case somelib.ErrEmpty:
        return ""
    default:
        try.Check(err)
    }

    return r
}
```