# stash 

**Documentation:** [![GoDoc](https://godoc.org/github.com/cogger/stash?status.png)](http://godoc.org/github.com/cogger/stash)  
**Build Status:** [![Build Status](https://travis-ci.org/cogger/stash.svg?branch=master)](https://travis-ci.org/cogger/stash)  
**Test Coverage:** [![Coverage Status](https://coveralls.io/repos/cogger/stash/badge.svg?branch=master)](https://coveralls.io/r/cogger/stash?branch=master)  
**License:**       [![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)


stash adds generic data to contexts

## Usage
~~~ go
// main.go
package main

import (
	"github.com/cogger/stash"
	"golang.org/x/net/context"
)

func main() {
	ctx := stash.Set(context.Background(),"something","somethingelse")
	ctx = stash.Set(ctx,"anotherKey",132.2)

	value, ok := stash.Get(ctx,"something").(string)
}

~~~

