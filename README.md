Romanumeral
===========

Golang application for converting number into roman numerals.

Usage
-----

The generator can be used in two ways, as a library or executable binary (one is included in the project already).


Binary Usage - `./main <integer>` or if you want to compile yourself then `go build cmd/romanumeral/main.go` followed by `./main <integer>` or
whatever name you compiled to.


Library Usage - `go get github.com/chvck/romanumeral` and then the easiest way to use is

```
generator := romanumeral.NewRomanNumeralGenerator()
roman, err := generator.Generate(<integer>)
```

Testing
-------
`go test`