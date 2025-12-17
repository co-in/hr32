package tests_test

import (
	"testing"

	btcutil "github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/btcsuite/btcutil/base58"
	"github.com/co-in/hr32"
)

func Benchmark_Encode_HR32(b *testing.B) {
	b32, _ := hr32.BIP173.Clone(hr32.WithPrefix(hrp))

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = b32.Encode(data)
	}
}

func Benchmark_Decode_HR32(b *testing.B) {
	b32, _ := hr32.BIP173.Clone(hr32.WithPrefix(hrp))
	encoded, _ := b32.Encode(data)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, _, _ = hr32.BIP173.Decode(encoded)
	}
}

func Benchmark_Encode_BTCU(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		conv, _ := btcutil.ConvertBits(data, 8, 5, true)
		_, _ = btcutil.Encode(hrp, conv)
	}
}
func Benchmark_Decode_BTCU(b *testing.B) {
	conv, _ := btcutil.ConvertBits(data, 8, 5, true)
	encoded, _ := btcutil.Encode(hrp, conv)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, conv, _ = btcutil.Decode(encoded)
		_, _ = btcutil.ConvertBits(conv, 5, 8, false)
	}
}

func Benchmark_Encode_Base58(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = base58.CheckEncode(data, 1)
	}
}

func Benchmark_Decode_Base58(b *testing.B) {
	s := base58.CheckEncode(data, 1)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, _ = base58.CheckDecode(s)
	}
}
