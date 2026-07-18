type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	es := "["

	for key, w := range strs {
		es += "\"" + w + "\""

		if key < len(strs) - 1 {
			es += ","
		}
	}

	es += "]"

	// fmt.Println(es)

	return es
}

func (s *Solution) Decode(encoded string) []string {
	r := []rune(encoded)

	list := []string{}
	w := ""
	for i := 2; i < len(r); i++ {
		if string(r[i]) == "\"" && string(r[i+1]) == "," {
			list = append(list, w)
			w = ""
			i += 2
			continue
		}

		if string(r[i]) == "\"" && string(r[i+1]) == "]" {
			list = append(list, w)
			break
		}

		w += string(r[i])
	}

	fmt.Println(list)

	return list
}
