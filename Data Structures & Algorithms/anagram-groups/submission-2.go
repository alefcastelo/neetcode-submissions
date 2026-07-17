func groupAnagrams(strs []string) [][]string {
	fm := make(map[string][]string)

	for _, w := range strs {
		hash := hashWord(w)

		fm[hash] = append(fm[hash], w)
	}

	var list [][]string

	for _, item := range fm {
		list = append(list, item)
	}

	return list
}

func hashWord(w string) string {
	r := []rune(w)
	fm := make(map[string]int)

	for i := 0; i < len(w); i++ {
		_, exists := fm[string(r[i])]

		if exists {
			fm[string(r[i])]++
			continue
		}

		fm[string(r[i])] = 1
	}

	data, _ := json.Marshal(fm)

	return string(data)
}
