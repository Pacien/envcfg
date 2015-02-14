envcfg [![Build Status](https://travis-ci.org/Pacien/envcfg.svg)](//travis-ci.org/Pacien/envcfg)
======

Package [envcfg](//github.com/Pacien/envcfg) provides environment variable mapping to structs.

It can be used to read configuration parameters from the environment.

Fields for which environment variables can be found are overwritten, otherwise they are left to their previous value.

This package can be used, for example, after [gcfg](//code.google.com/p/gcfg/) to override settings provided in a
configuration file.


```Go
import "github.com/Pacien/envcfg"
```


Documentation
-------------

Package documentation can be found on [GoDoc](//godoc.org/github.com/Pacien/envcfg)


Usage example
-------------

Set environment variables:

```Bash
export PORT=8080
export USER_PASSWORD="S3cUrE"
```


Create a struct optionally with tagged fields and/or already set values, then call the `ReadInto` function to read
the values set in the environment variables.

```Go
package main

import (
	"fmt"
	"github.com/Pacien/envcfg"
)

type Config struct {
	Server struct {
		Port        int `env:"PORT" absenv:"true"`
	}
	User struct {
		Username    string
		Password    string
	}
}

var cnf Config

func (c *Config) setDefaults() *Config {
	c.User.Username = "root"
	c.User.Password = "password"
	return c
}

func init() {
	cnf.setDefaults()

	_, errs := envcfg.ReadInto(&cnf)
	if len(errs) != 0 {
		fmt.Println(errs)
	}
}

func main() {
	fmt.Println(cnf)
}
```


Output of the previous program:

```Bash
{{8080} {root S3cUrE}}
```


See tests for other examples.


License
-------

This program is published under the MIT License.
See the LICENSE.txt file.
