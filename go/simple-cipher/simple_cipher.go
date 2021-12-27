package cipher

import (
	"regexp"
	"strings"
)

// Define the shift and vigenere types here.
type shift struct {
	distance int
}
type vigenere struct {
	key string
}

// Both types should satisfy the Cipher interface.

func NewCaesar() Cipher {
	return NewShift(3)
}

func MapBasedCipher(distance int) map[string]string {
	cipher := make(map[string]string)
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	if distance > 0 {
		adjustedAlphabet := append(alphabet, alphabet[:distance]...)
		for i := 0; i < len(alphabet); i++ {
			cipher[alphabet[i]] = adjustedAlphabet[i+distance]
		}
	} else {
		adjustedAlphabet := append(alphabet[len(alphabet)+distance:], alphabet...)
		for i := 0; i < len(alphabet); i++ {
			cipher[alphabet[i]] = adjustedAlphabet[i]
		}
	}
	return cipher
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance < -25 || distance > 25 {
		return nil
	}
	return &shift{distance: distance}
}

func (c shift) Encode(input string) string {
	r := regexp.MustCompile(`[a-zA-Z]`)
	filtered := r.FindAllString(input, -1)
	result := []string{}
	cipher := MapBasedCipher(c.distance)
	for _, char := range filtered {
		result = append(result, strings.ToLower(char))
	}
	encoded := []string{}
	for _, char := range result {
		for k, v := range cipher {
			if char == k {
				encoded = append(encoded, v)
			}
		}
	}
	return strings.Join(encoded, "")
}

func (c shift) Decode(input string) string {
	stringArr := strings.Split(input, "")
	cipher := MapBasedCipher(c.distance)
	decoded := []string{}
	for _, char := range stringArr {
		for k, v := range cipher {
			if char == v {
				decoded = append(decoded, k)
			}
		}
	}
	return strings.Join(decoded, "")
}

func MapVigenere() map[string]int {
	cipherMap := make(map[string]int)
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for index, char := range alphabet {
		cipherMap[char] = index
	}
	return cipherMap
}

func NewVigenere(key string) Cipher {
	for _, v := range key {
		if v < 'a' || v > 'z' {
			return nil
		}
	}
	count := 0
	for _, char := range key {
		if char == 'a' {
			count++
		}
	}
	if count == len(key) {
		return nil
	}
	if key == "" {
		return nil
	}
	return vigenere{key: key}
}

func (v vigenere) Encode(input string) string {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	// inputArr := strings.Split(input, "")
	numSequence := []int{}
	r := regexp.MustCompile(`[a-zA-Z]`)
	filtered := r.FindAllString(input, -1)
	for _, char := range strings.Split(v.key, "") {
		for alpha, num := range MapVigenere() {
			if strings.ToLower(char) == alpha {
				numSequence = append(numSequence, num)
			}
		}
	}
	finalResult := []string{}
	numSequence = append(numSequence, numSequence...)
	numSequence = append(numSequence, numSequence...)
	for i, v := range filtered {
		for j, a := range alphabet {
			if a == strings.ToLower(v) {
				alphabets := append(alphabet, alphabet[:numSequence[i]]...)
				finalResult = append(finalResult, alphabets[j+numSequence[i]])
				alphabets = nil
			}
		}
	}
	return strings.Join(finalResult, "")
}

func (v vigenere) Decode(input string) string {
	// increments key
	numSequence := []int{}
	for _, char := range strings.Split(v.key, "") {
		for alpha, num := range MapVigenere() {
			if strings.ToLower(char) == alpha {
				numSequence = append(numSequence, num)
			}
		}
	}
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	numSequence = append(numSequence, numSequence...)
	numSequence = append(numSequence, numSequence...)
	inputArr := strings.Split(input, "")
	finalResult := []string{}
	for i, v := range inputArr {
		for j, alp := range alphabet {
			j += numSequence[i]
			if strings.ToLower(v) == alp {
				alphabets := append(alphabet[len(alphabet)-numSequence[i]:], alphabet...)
				finalResult = append(finalResult, alphabets[j-numSequence[i]])
				alphabets = nil
			}
		}
	}
	return strings.Join(finalResult, "")
}
