package roman

func RomanToArabic(roman string) (int, bool) {
	romanNums := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4,
		"V": 5, "VI": 6, "VII": 7, "VIII": 8,
		"IX": 9,
	}

	if val, ok := romanNums[roman]; ok {
		return val, true
	}

	return 0, false
}
