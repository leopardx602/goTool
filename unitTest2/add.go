// $ go get -u github.com/cweill/gotests/...
// $ export PATH="$HOME/go/bin:$PATH"
// $ gotests -all -w add.go add_test.go
// $ go test -v -cover=true add_test.go add.go

package unitTest2

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
