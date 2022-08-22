package pkg

// RemoveDuplicatedItems function used for remove all duplicated items in a string array
// pass a string array as parameter
// return a string array
func RemoveDuplicatedItems(items []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range items {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// RemoveItemInArray function used for remove a specific string item in a string array
// pass a string array and a string item as parameter
// return a string array
func RemoveItemInArray(l []string, item string) []string {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}
