package essentials

import "fmt"

func Arrays() {
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	// You can also initialize an array with values
	arr2 := [5]int{6, 7, 8, 9, 10}

	fmt.Println("Array 1:", arr)
	fmt.Println("Array 2:", arr2)

	//println("Array:", arr)
	println("Length:", len(arr))
	println("Capacity:", cap(arr))
}

// Slices are similar to arrays, but are more powerful and flexible.
// Like arrays, slices are also used to store multiple values of the same type in a single variable.
// However, unlike arrays, the length of a slice can grow and shrink as you see fit.

// In Go, there are several ways to create a slice:

// Using the []datatype{values} format
// Create a slice from an array
// Using the make() function

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

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	println(len(myslice2))
	println(cap(myslice2))
	println(myslice2)

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))
}
