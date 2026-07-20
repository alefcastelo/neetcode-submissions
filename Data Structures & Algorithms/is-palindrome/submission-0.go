func isPalindrome(s string) bool {

	re := regexp.MustCompile(`[^A-Za-z0-9]`)
	clean := strings.ToLower(re.ReplaceAllString(s, ""))
	fmt.Println(clean)
	r := []rune(clean)
	i, j := 0, len(r) - 1
	for j >= i {

		if r[i] != r[j] {
			return false
		}

		i++
		j--
	}

	return true
}
