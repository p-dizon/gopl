package reverseptr

func reversePtrOddLength(arr *[5]int) {
	for i := 0; i < len(*arr)/2; i++ {
		lastIndex := len(*arr) - 1
		(*arr)[i], (*arr)[lastIndex-i] = (*arr)[lastIndex-i], (*arr)[i]
	}
}

func reversePtrEvenLength(arr *[6]int) {
	for i := 0; i < len(*arr)/2; i++ {
		lastIndex := len(*arr) - 1
		(*arr)[i], (*arr)[lastIndex-i] = (*arr)[lastIndex-i], (*arr)[i]
	}
}
