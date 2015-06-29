# stash 

[![GoDoc](https://godoc.org/github.com/cogger/stash?status.png)](http://godoc.org/github.com/cogger/stash)  
[![Build Status](https://travis-ci.org/cogger/stash.svg?branch=master)](https://travis-ci.org/cogger/stash)  
[![Coverage Status](https://coveralls.io/repos/cogger/stash/badge.svg?branch=master)](https://coveralls.io/r/cogger/stash?branch=master)  
[![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)

stash adds generic data to contexts
## Installation

The import path for the package is *gopkg.in/cogger/stash.v1*.

To install it, run:

    go get gopkg.in/cogger/stash.v1

## Usage
~~~ go
// main.go
package main

import (
	"gopkg.in/cogger/stash"
	"golang.org/x/net/context"
)

func main() {
	ctx := stash.Set(context.Background(),"something","somethingelse")
	ctx = stash.Set(ctx,"anotherKey",132.2)

	value, ok := stash.Get(ctx,"something").(string)
}

~~~

