import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {

	fm := make(map[string][]string)

	for _, str := range strs {
		hash := wHash(str)

		_, exists := fm[hash]

		if !exists {
			fm[hash] = []string{}
		}

		fm[hash] = append(fm[hash], str)
	}

	list := [][]string{}

	for _, val := range fm {
		list = append(list, val)
	}

	return list
}

func wHash(w string) string {
	r := []rune(w)
	slices.Sort(r)
	return string(r)
}
