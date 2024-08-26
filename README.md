# go-nsid

Package **nsid** provides an implementation of **BlueSky**'s **AT-Protocol**'s NSID, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-nsid

[![GoDoc](https://godoc.org/github.com/reiver/go-nsid?status.svg)](https://godoc.org/github.com/reiver/go-nsid)

## Examples

Split an NSID into its _domain-authority_ and _name_:

```golang
import "github.com/reiver/go-nsid"

// ...

domainAuthority, name := nsid.Split("com.example.fooBar")

// domainAuthority == "com.example"
// name == "fooBar"

```

Normalize an NSID into its _domain-authority_ and _name_:

```golang
import "github.com/reiver/go-nsid"

// ...

value := nsid.Normalize("COM.Example.fooBar")

// value == "com.example.fooBar"
```

Validate an NSID:

```golang
import "github.com/reiver/go-nsid"

// ...

err := nsid.Validate("com.example.fooBar")
if nil != err {
	return err
}

// value == "com.example.fooBar"
```

## Import

To import package **nsid** use `import` code like the follownig:
```
import "github.com/reiver/go-nsid"
```

## Installation

To install package **nsid** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-nsid
```

## Author

Package **nsid** was written by [Charles Iliya Krempeaux](http://reiver.link)
