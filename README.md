# go-personnummer [![Build Status](https://travis-ci.org/frozzare/go-personnummer.svg?branch=master)](https://travis-ci.org/frozzare/go-personnummer)

 Validate Swedish personal identity numbers.

 View the [docs](http://godoc.org/github.com/frozzare/go-personnummer).

## Installation

```
$ go get github.com/frozzare/go-personnummer
```

## Example

```go
package main

import (
	"github.com/frozzare/go-personnummer"
)

func main() {
	personnummer.Test(6403273813)
	//=> true
	
	personnummer.Test("19900101-0017")
	//=> true
	
	// works with co-ordination numbers
	personnummer.Test("701063-2391")
	//=> true

	personnummer.Test("510818-916")
	//=> false
}
```

# License

 MIT
