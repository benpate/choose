package choose

// String returns the first string from the set of arguments that is NOT an empty string.
func String(values ...string) string {

	for index := range values {
		if values[index] != "" {
			return values[index]
		}
	}

	return ""
}
