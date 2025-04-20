# go-nsid

Package **nsid** provides an implementation of **BlueSky**'s **AT-Protocol**'s NSID (Namespaced Identifier), for the Go programming language.

NSIDs are used by the Bluesky network and its AT-protocol, and are defined at:
https://atproto.com/specs/nsid

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-nsid

[![GoDoc](https://godoc.org/github.com/reiver/go-nsid?status.svg)](https://godoc.org/github.com/reiver/go-nsid)

## NSID

In plain language, an NSID (Namespaced Identifier) is a combination of **Internet domain-name** and a **name**.

### NSID Example 1

For example, if the **Internet domain-name** was:

	example.com

And the **name** was:

	fooBar

Then the resulting NSID (Namespaced Identifier) would be:

	com.example.fooBar

Note that the **Internet domain-name** (`example.com`) was written in reverse-order (`com.example`) before constructing the NSID (Namespaced Identifier).

### NSID Example 2

Here is another example.
If the **Internet domain-name** was:

	video.archive.org

Note that we are using a sub-domain of `example.com` here.

And if the **name** was:

	clipVideo

Then the resulting NSID (Namespaced Identifier) would be:

	org.archive.video.clipVideo

Again note that the **Internet domain-name** (`video.archive.org`) was written in reverse-order (`org.archive.video`) before constructing the NSID (Namespaced Identifier).

## Examples

To split an NSID into its _domain-authority_ and _name_:

```golang
import "github.com/reiver/go-nsid"

// ...

domainAuthority, name := nsid.Split("com.example.fooBar")

// domainAuthority == "com.example"
// name == "fooBar"

```

To normalize an NSID:

```golang
import "github.com/reiver/go-nsid"

// ...

value := nsid.Normalize("COM.Example.fooBar")

// value == "com.example.fooBar"
```

To validate an NSID:

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
