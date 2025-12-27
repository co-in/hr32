package hr32

func convert(data []byte, from, to uint8, pad bool) ([]byte, error) {
	result := make([]byte, 0, len(data)*int(from)/int(to)+1)
	next := byte(0)
	filled := uint8(0)

	for _, b := range data {
		b <<= 8 - from
		rFrom := from

		for rFrom > 0 {
			rTo := to - filled
			e := rFrom

			if rTo < e {
				e = rTo
			}

			next = (next << e) | (b >> (8 - e))
			b <<= e
			rFrom -= e
			filled += e

			if filled == to {
				result = append(result, next)
				filled = 0
				next = 0
			}
		}
	}

	if pad && filled > 0 {
		next <<= to - filled
		result = append(result, next)

		filled = 0
		next = 0
	}

	if filled > 0 && (filled > 4 || next != 0) {
		return nil, ErrGroupIncompleted{}
	}

	return result, nil
}

func poly(prefix string, values, checksum []byte, generators []int) int {
	chk := 1
	lenPrefix := len(prefix)

	for i := 0; i < lenPrefix; i++ {
		b := chk >> 25
		hi := int(prefix[i]) >> 5
		chk = (chk&0x1ffffff)<<5 ^ hi
		for j := uint(0); j < 5; j++ {
			if (b>>j)&1 == 1 {
				chk ^= generators[j]
			}
		}
	}

	b := chk >> 25
	chk = (chk & 0x1ffffff) << 5
	for i := 0; i < 5; i++ {
		if (b>>uint(i))&1 == 1 {
			chk ^= generators[i]
		}
	}

	for i := 0; i < lenPrefix; i++ {
		b = chk >> 25
		lo := int(prefix[i]) & 31
		chk = (chk&0x1ffffff)<<5 ^ lo
		for j := uint(0); j < 5; j++ {
			if (b>>j)&1 == 1 {
				chk ^= generators[j]
			}
		}
	}

	for _, v := range values {
		b = chk >> 25
		chk = (chk&0x1ffffff)<<5 ^ int(v)
		for i := 0; i < 5; i++ {
			if (b>>uint(i))&1 == 1 {
				chk ^= generators[i]
			}
		}
	}

	if checksum == nil {
		for v := 0; v < 6; v++ {
			b = chk >> 25
			chk = (chk & 0x1ffffff) << 5
			for i := 0; i < 5; i++ {
				if (b>>uint(i))&1 == 1 {
					chk ^= generators[i]
				}
			}
		}
	} else {
		for _, v := range checksum {
			b = chk >> 25
			chk = (chk&0x1ffffff)<<5 ^ int(v)
			for i := 0; i < 5; i++ {
				if (b>>uint(i))&1 == 1 {
					chk ^= generators[i]
				}
			}
		}
	}

	return chk
}

func calcChecksum(result *[6]byte, pm int, alphabet string) {
	for i := 0; i < 6; i++ {
		result[i] = alphabet[byte((pm>>uint(5*(5-i)))&31)]
	}
}
