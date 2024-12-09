package utils

func IsValueInList(value string, validValues []string) bool {
	for _, validValue := range validValues {
		if value == validValue {
			return true
		}
	}

	return false
}

func IsValueInMap(value string, validValues map[string]string) bool {
	_, exists := validValues[value]
	return exists
}
