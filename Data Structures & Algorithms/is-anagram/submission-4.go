func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hs := make(map[string]int)
	ht := make(map[string]int)
	
	for i := 0; i < len(s); i++ {
		if _, exists := hs[string(s[i])]; exists {
			hs[string(s[i])] = hs[string(s[i])] + 1
		} else {
			hs[string(s[i])] = 1
		}

		if _, exists := ht[string(t[i])]; exists {
			ht[string(t[i])] = ht[string(t[i])] + 1
		} else {
			ht[string(t[i])] = 1
		}
	}

	fmt.Println(hs, ht)

	for ks, vs := range hs {
		vt, exists := ht[ks]

		if !exists {
			return false
		}

		if vt != vs {
			return false
		}
	}

	return true
}
