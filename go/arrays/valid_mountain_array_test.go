package arrays

import "testing"

func TestValidMountainArray1(t *testing.T) {
	arr := []int{2, 1}
	actual := validMountainArray(arr)
	if actual != false {
		t.FailNow()
	}
}

func TestValidMountainArray2(t *testing.T) {
	arr := []int{3, 5, 5}
	actual := validMountainArray(arr)
	if actual != false {
		t.FailNow()
	}
}

func TestValidMountainArray3(t *testing.T) {
	arr := []int{0, 3, 2, 1}
	actual := validMountainArray(arr)
	if actual != true {
		t.FailNow()
	}
}

func TestValidMountainArray4(t *testing.T) {
	arr := []int{3, 2, 1}
	actual := validMountainArray(arr)
	if actual != false {
		t.FailNow()
	}
}

func TestValidMountainArray5(t *testing.T) {
	arr := []int{1, 2, 2, 1}
	actual := validMountainArray(arr)
	if actual != false {
		t.FailNow()
	}
}

func TestValidMountainArray6(t *testing.T) {
	arr := []int{1, 2, 3, 9, 2, 1}
	actual := validMountainArray(arr)
	if actual != true {
		t.FailNow()
	}
}

func TestValidMountainArray7(t *testing.T) {
	arr := []int{1, 2, 3, 5, 6}
	actual := validMountainArray(arr)
	if actual != false {
		t.FailNow()
	}
}
