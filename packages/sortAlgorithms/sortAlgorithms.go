package sortAlgorithms
	
import (
	"fmt"
)
	
func swap(slice []float64, indexA int, indexB int) {
	c := slice[indexA]
	slice[indexA] = slice[indexB]
	slice[indexB] = c
}
	
func BubleSort(slice []float64) {
	fmt.Println("Starting bublle sort")
	n := len(slice)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if slice[i-1] > slice[i] {
				swap(slice, i-1, i)
				swapped = true
			}
		}
		n--
	}
} 
