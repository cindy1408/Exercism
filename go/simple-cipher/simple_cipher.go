package cipher

// Define the shift and vigenere types here.
type shift string
type vigenere string
// Both types should satisfy the Cipher interface.

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
}

func (c shift) Encode(input string) string {
	panic("Please implement the Encode function")
}

func (c shift) Decode(input string) string {
	panic("Please implement the Decode function")
}

func NewVigenere(key string) Cipher {
	panic("Please implement the NewVigenere function")
}

func (v vigenere) Encode(input string) string {
	panic("Please implement the Encode function")
}

func (v vigenere) Decode(input string) string {
	panic("Please implement the Decode function")
}
