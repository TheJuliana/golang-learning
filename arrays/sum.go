package arrays

func Sum(mas [5]int) interface{} {
	s := 0
	for i := 0; i < len(mas); i++ {
		s += mas[i]
	}
	return s
}
