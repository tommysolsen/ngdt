package helpers

// InArray lets you check if a needle exists in a haystack
func InArray(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
