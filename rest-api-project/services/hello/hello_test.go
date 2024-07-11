package hello_test

import "testing"

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func TestMain(t *testing.T) {
	t.Run("number testing", func(t *testing.T) {
		if isEven(2) != true {
			t.Errorf("test failed num")
		}
	})

}

func TestSecond(t *testing.T) {
	t.Run("number testing 2", func(t *testing.T) {
		if isEven(3) != true {
			t.Errorf("test 2 failed num")
		}
	})
}
