## cuid

Collision-resistant hashes for the cloud, in Go.

The ```cuid`` package provides collision-resistant ids optimized for horizontal scaling and sequential lookup performance. This README file is just going to cover the basics and the Go-specific implementation details.

**Please refer to the [main project site](http://usecuid.org) for the full rationale behind CUIDs.**

## Sample CUID

ch72gsb320000udocl363eofy

# Installation

**HEAD:**

```go get github.com/lucsky/cuid```

**v1.0.0:**

```go get gopkg.in/lucsky/cuid.v1```

## Example usage

```Go
package main

import fmt
import "gopkg.in/lucsky/cuid.v1"

func main() {
    fmt.Println(cuid.New())
}
```

## Go package specific features

The Go cuid package provides APIs to specify a custom random source as well as a custom counter. A custom counter implementation could provide a centralized Redis base counter, for example.

## Contributors

- Luc Heinrich (lucsky, author)
- Thomas Hopkins (hopkinsth)
