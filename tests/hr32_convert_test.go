package tests

import (
	"testing"

	"github.com/co-in/hr32"
	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	res, err := hr32.Default.Convert("am0RAVW2CG4TQ3RH7K8YMDAEVNGQ4FTQV28T8EFKE", func(_ string, version int) *hr32.Config {
		cfg, err := hr32.BIP173.Clone(
			hr32.WithPrefix("bc"),
			hr32.WithVersion(version),
		)
		assert.NoError(t, err)

		return cfg
	})
	assert.NoError(t, err)
	assert.Equal(t, "bc15qe6rzgtkj85fhvmawyq9esgjtxkjermajh0u6", res)
}
