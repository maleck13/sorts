package insertion

func Sort(vals []int) {
	// assume first is sorted as nothing to compare to thats why k=1
	for k := 1; k < len(vals); k++ {
		sorting := vals[k] // hold the val we are sorting
		i := k             // we are going to count back so assign a new val
		for i >= 0 {       // count back if we are at the start or the previous val is < than what we are sorting we are done
			if i == 0 || vals[i-1] < sorting {
				vals[i] = sorting
				break
			}
			// assign the current val with the previous val (move it)
			if vals[i-1] > sorting {
				vals[i] = vals[i-1]
				i--
			}

		}
	}
}
