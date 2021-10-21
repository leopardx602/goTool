// There are several ways to create a slice with a certain length and value.
// Please write benchmarks for these functions,
// and explain why their performance is different.

package unitTest_benchmark

const length = 1000

// genByAppend creates a slice of integers whose values are 0 to 999.
func GenByAppend() []int {
	var s []int
	for i := 0; i < length; i++ {
		s = append(s, i)
	}
	return s
}

// genByAppendCap creates a slice of integers whose values are 0 to 999.
func GenByAppendCap() []int {
	s := make([]int, 0, length)
	for i := 0; i < length; i++ {
		s = append(s, i)
	}
	return s
}

// genByAssign creates a slice of integers whose values are 0 to 999.
func GenByAssign() []int {
	s := make([]int, length)
	for i := 0; i < length; i++ {
		s[i] = i
	}
	return s
}
