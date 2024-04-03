package property

func ConvertToRoman(arabic int) string {

	result := ""
	for i := 0; i < arabic; i++ {
		result += "I"
	}

	return result

}
