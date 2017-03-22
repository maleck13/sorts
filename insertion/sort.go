package insertion

import "fmt"

// Sort is an implementation of the insertion sort algorthem
func Sort(arr []int) {
	sorts := 0
	// assume first element sorted so start at index 1 rather than 0
	for i := 1; i < len(arr); i++ {
		// start looking from the previous element
		j := i - 1
		// set a temp value for the element value we are sorting
		value := arr[i]
		// if our value is less than the previous value
		for j >= 0 && value < arr[j] {
			// move the previous value up one duplicating it across two spaces [j] and [j+1] which gives us a space to put the value in once we find the right place
			arr[j+1] = arr[j]
			sorts++
			// decrement our counter and go again
			j--
		}
		//we have now found an element that is >= value so add this value to space we created which is our current pos +1
		arr[j+1] = value
	}
	fmt.Print(sorts)
}
