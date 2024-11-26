package math

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func MaxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func B2I(v bool) int {
	if v {
		return 1
	}

	return 0
}

func RoundInt(x float64) int {
	return int(math.Floor(x + 0.5))
}

func CeilMode(n, m int) int {
	v := n / m

	if v*m < n {
		return v + 1
	}

	return v
}

func IIfInt(b bool, n, m int) int {
	if b {
		return n
	}

	return m
}

func IIfInt64(b bool, n, m int64) int64 {
	if b {
		return n
	}

	return m
}

func IIfFloat(b bool, n, m float64) float64 {
	if b {
		return n
	}

	return m
}

func Range(start int, end int) []int {
	nums := make([]int, end-start+1)

	for n := start; n <= end; n++ {
		nums[n-start] = n
	}

	return nums
}

// 精度问题
func FloatToFloat(f float64) float64 {
	realNum, _ := strconv.ParseFloat(fmt.Sprintf("%.8f", f), 64)
	return realNum
}

// 保留小数[四舍五入]
func FloatRound(f float64, count int) float64 {
	s := fmt.Sprintf("%s.%df", "%", count)
	realNum, _ := strconv.ParseFloat(fmt.Sprintf(s, f), 64)
	return realNum
}

// 保留小数 [舍去]
func FloatFloor(f float64, count int) float64 {
	s := strconv.FormatFloat(float64(int(f*math.Pow(10, float64(count))))/math.Pow(10, float64(count)), 'f', count, 64)
	realNum, _ := strconv.ParseFloat(s, 64)
	return realNum
}

// 随机数 范围【start, end】
func RandomInt(start int, end int) int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(end - start)
	random = start + random
	return random
}

func InIntArray(val int, list []int) (exists bool, index int) {
	exists, index = false, -1

	for i, v := range list {
		if val == v {
			exists, index = true, i
			break
		}
	}

	return
}

func InInt64Array(val int64, list []int64) (exists bool, index int) {
	exists, index = false, -1

	for i, v := range list {
		if val == v {
			exists, index = true, i
			break
		}
	}

	return
}

func InFloatArray(val float64, list []float64) (exists bool, index int) {
	exists, index = false, -1

	for i, v := range list {
		if val == v {
			exists, index = true, i
			break
		}
	}

	return
}

func InIntSection(n, min, max int, contain ...bool) bool {
	minContain, maxContain := true, true
	if len(contain) >= 1 {
		minContain = contain[0]
	}
	if len(contain) >= 2 {
		maxContain = contain[1]
	}

	if minContain && maxContain {
		if n >= min && n <= max {
			return true
		}
	} else if minContain && !maxContain {
		if n >= min && n < max {
			return true
		}
	} else if !minContain && maxContain {
		if n > min && n <= max {
			return true
		}
	} else if !minContain && !maxContain {
		if n > min && n < max {
			return true
		}
	}

	return false
}

func InInt64Section(n, min, max int64, contain ...bool) bool {
	minContain, maxContain := true, true
	if len(contain) >= 1 {
		minContain = contain[0]
	}
	if len(contain) >= 2 {
		maxContain = contain[1]
	}

	if minContain && maxContain {
		if n >= min && n <= max {
			return true
		}
	} else if minContain && !maxContain {
		if n >= min && n < max {
			return true
		}
	} else if !minContain && maxContain {
		if n > min && n <= max {
			return true
		}
	} else if !minContain && !maxContain {
		if n > min && n < max {
			return true
		}
	}

	return false
}

func InFloatSection(n, min, max float64, contain ...bool) bool {
	minContain, maxContain := true, true
	if len(contain) >= 1 {
		minContain = contain[0]
	}
	if len(contain) >= 2 {
		maxContain = contain[1]
	}

	if minContain && maxContain {
		if n >= min && n <= max {
			return true
		}
	} else if minContain && !maxContain {
		if n >= min && n < max {
			return true
		}
	} else if !minContain && maxContain {
		if n > min && n <= max {
			return true
		}
	} else if !minContain && !maxContain {
		if n > min && n < max {
			return true
		}
	}

	return false
}
