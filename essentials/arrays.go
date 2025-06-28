package essentials

func Arrays() {
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	// You can also initialize an array with values
	arr2 := [5]int{6, 7, 8, 9, 10}

	println("Array 1:", arr)
	println("Array 2:", arr2)

	//println("Array:", arr)
	println("Length:", len(arr))
	println("Capacity:", cap(arr))
}
func Slices() {
	slice := []int{1, 2, 3, 4, 5}
	println("Slice:", slice)
	println("Length:", len(slice))
	println("Capacity:", cap(slice))

	// Append to slice
	slice = append(slice, 6, 7, 8)
	println("After append:", slice)

	// Slice a portion of the slice
	subSlice := slice[2:5]
	println("Sub-slice:", subSlice)

	// Copying slices
	copiedSlice := make([]int, len(slice))
	copy(copiedSlice, slice)
	println("Copied Slice:", copiedSlice)
}
