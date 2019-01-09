# Horde

Masterless node tracker.

### What is this for?

To have nodes talk to eachother and be aware of eachother in realtime.

Replace need for centralized cache. Automatically shard data throughout a cluster.

### TODO

All of it :pray:

### Example (very early API)

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/selfup/horde"
)

func main() {
	hordeManager := horde.Boot()

	res := hordeManager.Ping()

	fmt.Println(res)

	http.ListenAndServe(":8080", nil)
}

```
