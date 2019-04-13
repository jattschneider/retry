# Go Retry

Retry is a Go library that implements a simple retry.

## Getting Started

Just a quick example how to use the retry library:

#### main.go
```
package main

import (
	"flag"
	"fmt"

	"github.com/jattschneider/retry"
)

func init() {
	flag.Parse()
}

func main() {
	callCount := 0
	err := retry.With(func(attempts int) error {
		fmt.Printf("Attempt: %d\n", attempts)
		callCount++
		return nil
	}, 5)
}
```

```
$ go run main.go
```

### Installing

```
go get -v github.com/jattschneider/retry
```

## Built With

* [Go](https://golang.org/) - The Go Programming Language
* [dep](https://golang.github.io/dep/) - Dependency management for Go

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/jattschneider/config/tags). 

## Authors

* **José Augusto Schneider** - *Initial work* - [jattschneider](https://github.com/jattschneider)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
