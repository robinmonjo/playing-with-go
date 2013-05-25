package sorting
	
//interface for sortable object
type Sortable interface {
	length() int 						//return the size of the slice
	less(i, j int) bool 		//return self[i] < self[j]
	swap(i, j int) 					//swap self[i] sefl[j]
}

//sortable Numbers
type SortableNumbers []float64

func (numbers SortableNumbers) length() int {
	return len(numbers)
}

func (numbers SortableNumbers) less(i, j int) bool {
	return numbers[j] <  numbers[i]
}

func (numbers SortableNumbers) swap(i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}

//sorting algo	
func Sort(sortable Sortable) {
	n := sortable.length();
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if sortable.less(i - 1, i) {
				sortable.swap(i - 1, i)
				swapped = true
			}
		}
		n--
	}
} 
