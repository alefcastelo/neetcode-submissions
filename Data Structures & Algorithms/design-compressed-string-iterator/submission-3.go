
type Element struct {
	char byte
	amount int	
}

type StringIterator struct {
	position int
	compressedRunes []rune
	elements map[int]Element
}

func Constructor(compressedString string) StringIterator {
	si := StringIterator{
		position: 0,
		compressedRunes: []rune(compressedString),
		elements: make(map[int]Element),
	}

	si.LoadElements()

	return si
}

func (this *StringIterator) LoadElements() {
	charIndex := 0
	index := 0
	for index < len(this.compressedRunes) {
		el := Element{
			char: byte(this.compressedRunes[index]),
			amount: 0,
		}

		index++

		amountString := ""
		for index < len(this.compressedRunes) {
			if unicode.IsDigit(this.compressedRunes[index]) {
				amountString += string(this.compressedRunes[index])
				index++
				continue
			}

			break
		}

		amount, _ := strconv.Atoi(amountString)
		el.amount = amount
		this.elements[charIndex] = el
		charIndex++
	}

}

func (this *StringIterator) Next() byte {
	if this.elements[this.position].amount > 1 {
		el := this.elements[this.position]
		el.amount--
		this.elements[this.position] = el
		return el.char
	}

	el := this.elements[this.position]
	el.amount = el.amount - 1
	this.elements[this.position] = el
	this.position++

	return el.char
}

func (this *StringIterator) HasNext() bool {
	_, exists := this.elements[this.position]
	return exists
}

/**
 * Your StringIterator object will be instantiated and called as such:
 * obj := Constructor(compressedString);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
