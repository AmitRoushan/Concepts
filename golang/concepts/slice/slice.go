package main

import "fmt"

/*
   Slice is composite data type which act as reference to an array.
   Zero-Value slice means a slice with nil value
*/

func main() {
	// zero value slice
	var slice []int
	// check slice value
	fmt.Println(fmt.Sprintf("Print the value of zero slice %v", slice))
	// test slice with nil
	fmt.Println(fmt.Sprintf("test zero value slice with nil %v", slice == nil))

	// create an array and make zero value slice reference it
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice = array[1:4]
	// now print slice
	fmt.Println(fmt.Sprintf("Print the value of slice %v", slice))

	// to test it refer array
	slice[0] = 10
	fmt.Println(fmt.Sprintf("Print the value of array[1] = %v and slice[0] = %v", array[1], slice[0]))

	// -> length and capacity of slice
	// length represent how may elements slice is current refering
	// capacity represent how much more can be accomodated
	fmt.Println(fmt.Sprintf("Print length = %v and capacity = %v", len(slice), cap(slice)))

	/*
	   Actually slice can be represented as
	    type slice struct {
	    zerothElement *type  // pointer point first element of an array which slice refer
	    len int
	    cap int
	    }
	*/
	// check slice represent same array
	fmt.Println(fmt.Sprintf("Print &array[1] = %v and &slice[0] = %v", &array[1], &slice[0]))

	// -> append functionality
	// append increases len of slice with the number of element added
	// append uses same array and updates value in array if appended values are mot more than array
	slice = append(slice, 10)
	fmt.Println(fmt.Sprintf("Print the values of slice %v", slice))
	fmt.Println(fmt.Sprintf("Print the values of array %v", array))
	fmt.Println(fmt.Sprintf("Print the slice length = %v and capacity %v", len(slice), cap(slice)))

	// But if adds value more than array then slice create new array and double the sized of earlier capacity and append values in it.
	oldSlice := slice
	slice = append(slice, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
	fmt.Println(fmt.Sprintf("Print the values of slice %v", slice))
	fmt.Println(fmt.Sprintf("Print the values of array %v", array))
	fmt.Println(fmt.Sprintf("Print the values of old slice %v", oldSlice))
	fmt.Println(fmt.Sprintf("Print the slice length = %v and capacity %v", len(slice), cap(slice)))

	// anonymous array slice
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(fmt.Sprintf("Print the values of slice %v", slice))
	fmt.Println(fmt.Sprintf("Print the slice length = %v and capacity %v", len(slice), cap(slice)))
	// add just one more element
	slice = append(slice, 10)
	fmt.Println(fmt.Sprintf("Print the values of slice %v", slice))
	fmt.Println(fmt.Sprintf("Print the slice length = %v and capacity %v", len(slice), cap(slice)))

	// -> copy function
	// copy only elements which is minimum of src/dest slice
	// it will not copy any element if slice is zero value

	// append one slice to other by unpack operator
	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, 0, 5)
	slice2 = append(slice2, slice1...)
	fmt.Println(fmt.Sprintf("Print the values of slice %v", slice2))
	fmt.Println(fmt.Sprintf("Print the slice length = %v and capacity %v", len(slice2), cap(slice2)))


	b := make([]int, len(slice1))
	b = append(slice[0:0], slice...)
	fmt.Println(fmt.Sprintf("Print the values of slice %v", b))
	fmt.Println(fmt.Sprintf("Print the values of slice %v", slice))

}
