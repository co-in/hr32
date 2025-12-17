package hr32

var Default, _ = NewConfig(
	WithSeparator('0'),
	WithAlphabet("ABC2DEF3GHJ4KLM5NPQ6RST7UVW8XYZ9"), //without 01IO
	WithMaxLen(60),
	WithGenerators([]int{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}),
	WithVersion(0x2bc830a3),
)

var BIP173, _ = NewConfig(
	WithSeparator('1'),
	WithMinLen(8),
	WithMaxLen(90),
	WithAlphabet("qpzry9x8gf2tvdw0s3jn54khce6mua7l"),
	WithGenerators([]int{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}),
	WithVersion(0x2bc830a3),
)
