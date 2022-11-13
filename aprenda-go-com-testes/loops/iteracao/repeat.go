package iteracao

func Repeat(character string, quantity int64) string {
	var repetitions string
	for i := int64(0); i < quantity; i++ {
		repetitions += character
	}

	return repetitions
}
