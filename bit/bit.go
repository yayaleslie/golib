package bit

// 位运算

func SwitchBit(val, oldTag, newTag int) int {
	// 位清零
	val &= ^oldTag
	// 位赋值
	val |= newTag
	return val
}

func SetBit(val, tag int) int {
	val |= tag
	return val
}
