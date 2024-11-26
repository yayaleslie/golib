package convert

func BytesToBits(data []byte) []uint8 {
	dst := make([]uint8, 0)
	for _, v := range data {
		for i := 0; i < 8; i++ {
			move := uint8(7 - i)
			dst = append(dst, (v>>move)&1)
		}
	}
	return dst
}

func ByteToBits(data byte) []uint8 {
	dst := make([]uint8, 8)
	for i := 0; i < 8; i++ {
		move := uint8(7 - i)
		dst[i] = (data >> move) & 1
	}
	return dst
}

func BitToInt(data []byte) int {
	var dst int
	for i := 0; i < len(data); i++ {
		dst += int(data[i]) << i
	}
	return dst
}
