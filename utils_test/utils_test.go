package utils_test

import (
	"testing"

	"github.com/rafa-souza-dev/dog-balancer/utils"
)

func TestIncrementSliceIndexWhenRecieveLastIndex(t *testing.T) {
	slice := []int{1, 2, 3}
	lastIndex := len(slice) - 1
	index := lastIndex

	utils.IncrementSliceIndex(&index, slice)

	if index != 0 {
		t.Fatalf("the index value is not zero %v", index)
	}
}

func TestIncrementSliceIndexWhenRecieveMediumIndex(t *testing.T) {
	slice := []int{1, 2, 3}
	index := 1

	utils.IncrementSliceIndex(&index, slice)

	if index != 2 {
		t.Fatalf("the index was not correctly incremented %v", index)
	}
}
