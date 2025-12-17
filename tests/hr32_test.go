package tests_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/btcutil/base58"
	btcutil "github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/co-in/hr32"
	"github.com/stretchr/testify/assert"
)

var (
	data, _ = hex.DecodeString("a033a1890bb48f44dd9beb8802e60892cd69647b")
	hrp     = "bc"
)

func TestBackCompatibilityEncode(t *testing.T) {
	conv, err := btcutil.ConvertBits(data, 8, 5, true)
	assert.NoError(t, err)
	encoded, err := btcutil.Encode(hrp, conv)
	assert.NoError(t, err)
	conv, err = btcutil.ConvertBits(data, 8, 5, true)
	assert.NoError(t, err)
	encodedM, err := btcutil.EncodeM(hrp, conv)
	assert.NoError(t, err)

	b32m, err := hr32.BIP173.Clone(hr32.WithPrefix(hrp))
	assert.NoError(t, err)
	b32, err := b32m.Clone(hr32.WithVersion(1))
	assert.NoError(t, err)
	encodedM, err = b32m.Encode(data)
	assert.NoError(t, err)
	encoded, err = b32.Encode(data)
	assert.NoError(t, err)

	dLen, err := b32m.Len(data)
	assert.NoError(t, err)
	assert.Equal(t, dLen, len(encodedM))

	rPrefix, rVersion, rData, err := b32.Decode(encoded)
	assert.NoError(t, err)
	assert.Equal(t, data, rData)
	assert.Equal(t, 1, rVersion)
	assert.Equal(t, hrp, rPrefix)

	rPrefix, rVersion, rData, err = b32.Decode(encodedM)
	assert.NoError(t, err)
	assert.Equal(t, data, rData)
	assert.Equal(t, 0x2bc830a3, rVersion)
	assert.Equal(t, hrp, rPrefix)

	b32d, err := hr32.Default.Clone(hr32.WithPrefix("am"))
	assert.NoError(t, err)

	encodedD, err := b32d.Encode(data)
	assert.NoError(t, err)

	rPrefix, rVersion, rData, err = b32d.Decode(encodedD)
	assert.NoError(t, err)
	assert.Equal(t, data, rData)
	assert.Equal(t, 0x2bc830a3, rVersion)
	assert.Equal(t, "am", rPrefix)

	b32e, err := b32d.Clone(hr32.WithExcludePrefix(true))
	assert.NoError(t, err)
	assert.NotNil(t, b32e)

	encodedDE, err := b32e.Encode(data)
	assert.NoError(t, err)

	rPrefix, rVersion, rData, err = b32e.Decode("MY_BUSINESS_" + encodedDE)
	assert.NoError(t, err)
	assert.Equal(t, data, rData)
	assert.Equal(t, 0x2bc830a3, rVersion)
	assert.Equal(t, "MY_BUSINESS_am", rPrefix)

	fmt.Println(encodedD)
	fmt.Println(encoded)
	fmt.Println(encodedM)
}
func TestConvertBase58(t *testing.T) {
	var ver byte = 1

	encoded := base58.CheckEncode(data, ver)
	resData, resVer, err := base58.CheckDecode(encoded)

	assert.NoError(t, err)
	assert.Equal(t, resVer, ver)
	assert.Equal(t, resData, data)

	fmt.Println(encoded)
}
func TestConvertBU32E(t *testing.T) {
	conv, err := btcutil.ConvertBits(data, 8, 5, true)
	assert.NoError(t, err)
	encoded, err := btcutil.Encode(hrp, conv)
	assert.NoError(t, err)

	resHrp, conv, err := btcutil.Decode(encoded)
	assert.NoError(t, err)
	resData, err := btcutil.ConvertBits(conv, 5, 8, false)
	assert.NoError(t, err)

	assert.Equal(t, resHrp, hrp)
	assert.Equal(t, resData, data)

	fmt.Println(encoded)
}
