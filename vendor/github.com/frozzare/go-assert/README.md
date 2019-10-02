# go-assert [![Build Status](https://travis-ci.org/frozzare/go-assert.svg?branch=master)](https://travis-ci.org/frozzare/go-assert)

Asserts to Go testing

View the [docs](http://godoc.org/github.com/frozzare/go-assert).

## Installation

```
$ go get github.com/frozzare/go-assert
```

## Example

```go
package mypack

import (
  "testing"
  "github.com/frozzare/go-assert"
)

func TestName(t *testing.T) {
  assert.Equal(t, "Foo", "Foo")
}
```

# License

MIT © [Fredrik Forsmo](https://github.com/frozzare)
