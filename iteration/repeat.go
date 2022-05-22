package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func ExampleRepeat(char string, n int) string {
	var rep string
	for i := 0; i < n; i++ {
		rep += char
	}
	return rep
}
