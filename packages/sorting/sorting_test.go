package sorting

import(
 "testing"
 "math/rand"
)

func randomSlice() (res []float64) {
	for i := 0; i < 7; i++ {
		res = append(res, rand.Float64())
	}
	return
}

func TestSort(t *testing.T) {
	randomSlice := randomSlice()
	Sort(SortableNumbers(randomSlice))
	for i := 1; i < len(randomSlice); i++ {
		if (randomSlice[i - 1] > randomSlice[i]) {
			t.Log("Array ", randomSlice, "has value", randomSlice[i - 1], "before", randomSlice[i])
			t.FailNow()
		}
	}
}