module github.com/co-in/hr32/tests

go 1.25

replace github.com/co-in/hr32 => ../

require (
	github.com/btcsuite/btcd/btcutil v1.1.6
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/co-in/hr32 v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
