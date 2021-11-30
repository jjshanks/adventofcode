package challenge

func convert(in string) string {
	input := []byte(in)
	output := []byte{}
	last, count := input[0], byte(0)
	var c byte
	for _, c = range input {
		if last == c {
			count++
		} else {
			output = append(output, count+byte(48))
			output = append(output, last)
			last = c
			count = 1
		}
	}
	output = append(output, count+byte(48))
	output = append(output, c)
	return string(output)
}
