# moon

[![GoDoc](https://godoc.org/github.com/liamg/moon?status.svg)](<https://godoc.org/github.com/liamg/moon>)

Need to warn against deploys on a full moon?

Client needs a report generating every [waxing gibbous](https://en.wikipedia.org/wiki/Waxing_gibbous)?

No problem! This is a module in pure Go for calculating moon phase.

## Usage

```go
package main

import (
 "fmt"

 "github.com/liamg/moon"
)

func main() {
 phase := moon.GetPhase()
 fmt.Printf("The moon phase is currently %s - %s\n", phase, phase.Emoji())
 // Output example: The moon phase is currently Waxing Crescent - 🌒
}
```

---
