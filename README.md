# Go Service Library for Merrymake

This is the official Go service library for Merrymake. It defines all the basic functions needed to work with Merrymake.

## Usage

Here is the most basic example of how to use this library: 

```go
package main

import (
	"fmt"
	"os"
	m "github.com/merrymake/go-service-library"
	l "github.com/merrymake/go-service-library/lib"
)

func handleHello(payloadBytes []byte, envelope l.Envelope) {
	payload := string(payloadBytes)

	m.ReplyStringToOrigin(fmt.Sprintf("Hello, %s!", payload), l.GetMimeType("txt"))
}

func main() {
	args := os.Args[1:]
	new(m.Merrymake).Service(args).
		Handle("handleHello", handleHello)
}
```

## Tutorials and templates

For more information check out our tutorials at [merrymake.dev](https://merrymake.dev).

All templates are available through our CLI and on our [GitHub](https://github.com/merrymake).
