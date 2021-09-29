package unitTest

func Add(x, y int) int {
	return x + y
}

func FindName(findName string, nameList []string) bool {
	for _, name := range nameList {
		if name == findName {
			return true
		}
	}
	return false
}
