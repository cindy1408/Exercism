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
				increments := (rails-(i+1))*2
				firstLoop := true 
				if firstLoop {
					for j:=i; j<len(messageArr); j+= increments {
						encoded = append(encoded, messageArr[j])
						break 
					}
					firstLoop = false 
				} else {
					for j:=i; j<len(messageArr); j+= increments/2 {
						encoded = append(encoded, messageArr[j])
						break 
					}
					firstLoop = true
				}
			} else if i % 2 == 0 {
				firstLoop := true
				if firstLoop {
					for j:=i; j<len(messageArr); j+=i {
						encoded = append(encoded, messageArr[j])
						break
					} 
					firstLoop = false 
				} else {
					for j:=i; j<len(messageArr); j+=i*2 {
						encoded = append(encoded, messageArr[j])
						break
					}
					firstLoop = true 
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
