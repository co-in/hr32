## Human Readable Representation with BCH checksum

### Motivation:

I liked the idea in [BIP173](https://github.com/bitcoin/bips/blob/master/bip-0173.mediawiki), but I didn't like the
existing implementations, especially regarding the fixed alphabet and
the version for the checksum. Therefore, I decided to make it more flexible.

### Pros:

- Configurability
- Consistency check using BCH checksum
- Alphabet of 32 symbols (5 bits per symbol)

### Examples:

```go
package main

import (
	"encoding/hex"
	"fmt"

	"github.com/co-in/hr32"
)

var data, _ = hex.DecodeString("a033a1890bb48f44dd9beb8802e60892cd69647b")

func main() {
	//BIP173 backward compatible
	cfg, err := hr32.BIP173.Clone(
		hr32.WithPrefix("bc"),
	)
	if err != nil {
		panic(err)
	}

	encoded, err := cfg.Encode(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("BIP173\t", encoded)

	//Another preset
	cfg, err = hr32.Default.Clone(
		hr32.WithPrefix("addr"),
	)
	if err != nil {
		panic(err)
	}

	encoded, err = cfg.Encode(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Default\t", encoded)

	//Decode example
	fmt.Println()

	prefix, version, decoded, err := cfg.Decode(encoded)
	if err != nil {
		panic(err)
	}

	fmt.Println("Prefix:\t", prefix)
	fmt.Printf("Version: 0x%0X\n", version)
	fmt.Printf("Decoded: %0x\n", decoded)
}

```
