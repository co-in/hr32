### Converter for better Human Readable Representation of data

Motivation:
- 

Examples:

```go
package main

import (
	"encoding/hex"
	"fmt"

	"github.com/co-in/hr32"
)

var data, _ = hex.DecodeString("a033a1890bb48f44dd9beb8802e60892cd69647b")

func main() {
	cfg, err := hr32.BIP173.Clone(
		hr32.WithPrefix("bc"),
	)
	if err != nil {
		panic(err)
	}

	res, err := cfg.Encode(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("BIP173", res)

	cfg, err = hr32.Default.Clone(
		hr32.WithPrefix("pub"),
	)
	if err != nil {
		panic(err)
	}

	res, err = cfg.Encode(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Default", res)
}
```

Pros:

- Configurability
- Consistency check using BCH checksum
- Alphabet of 32 symbols (5 bits per symbol)
- Ability to compose parts of the representation
