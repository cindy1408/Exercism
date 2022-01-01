package railfence

import (
	"fmt"
	"strings"
)

func Encode(message string, rails int) string {
	fmt.Println("rails: ", rails)
	messageArr := strings.Split(message, "")
	encoded := []string{}
	for i:= 0; i<rails; i++ {
		if rails <= 2 {
			for j:=i; j<len(messageArr); j+=rails {
				encoded = append(encoded, messageArr[j])
			}
		} else {
			increment := (rails-1)*2
			if i == 0 ||  i == rails-1 {
				for j:=i; j<len(messageArr); j+= increment {
					encoded = append(encoded, messageArr[j])
				}
			} else if i % 2 != 0 {
				increment := (rails-(i+1))*2
				for j:=i; j<len(messageArr); j+= increment {
					encoded = append(encoded, messageArr[j])
				}
			} else if i % 2 == 0 {
				for j:=i; j<len(messageArr); j+=i {
					encoded = append(encoded, messageArr[j])
				}
			}
		}
	}
	fmt.Println("message: ", message)
	fmt.Println(encoded)
	return strings.Join(encoded, "")
}

func Decode(message string, rails int) string {
	panic("Please implement the Decode function")
}
