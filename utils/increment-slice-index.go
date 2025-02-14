package utils

func IncrementSliceIndex(index *int, slice []int) {
	if len(slice) == 0 {
		*index = 0

		return
	}

	lastIndex := len(slice) - 1

	if *index == lastIndex {
		*index = 0

		return
	}

	*index++
}
