package calculator

//Add means public and add - is private function that can't accessible outside
func Add(x ...float64) float64 {
	var sum float64
	sum = 0
	for _, v := range x {
		sum += v
	}
	return sum
}
