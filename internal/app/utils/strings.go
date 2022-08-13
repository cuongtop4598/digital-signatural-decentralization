package utils

func RemoveDuplicateString(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range strSlice {
			if _, value := allKeys[item]; !value {
					allKeys[item] = true
					list = append(list, item)
			}
	}
	return list
}

func StringContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}