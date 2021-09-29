// $ go test -v -cover=true add_test.go add.go
// -v		detail
// -cover	test percent, if some returns not be test, it would not be 100%
package unitTest

import "testing"

func TestAdd(t *testing.T) {
	ans := Add(1, 2)
	if ans == 3 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func TestFindName(t *testing.T) {
	ans := FindName("chen", []string{"chen", "ting"})
	if ans {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
func TestFindName2(t *testing.T) {
	ans := FindName("scott", []string{"chen", "ting"})
	if !ans {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
