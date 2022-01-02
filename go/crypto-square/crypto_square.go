package cryptosquare

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	m := regexp.MustCompile(`[a-zA-Z0-9]`)
	filtered := m.FindAllString(pt, -1)
	lowercase := []string{}
	for _, v := range filtered {
		lowercase = append(lowercase, strings.ToLower(v))
	}
	if len(lowercase) == 0{
		return ""
	}
	charLength := len(lowercase)
	r := int(math.Sqrt(float64(charLength)))
	c := r 
	if c * r < len(lowercase){
		c++
	}
	resultArr := [][]string{}
	part := []string{}
	for i:=0; i<len(lowercase); i++ {
		if len(part) < c-1 && i + 1 == len(lowercase){
			part = append(part, lowercase[i])
			for i:=0; len(part)< c; i++{
				part = append(part, " ")
			}
			
			resultArr = append(resultArr, part)
			part = nil 
		} else if len(part) < c-1 {
			part = append(part, lowercase[i])
		} else if len(part) == c -1{
			part = append(part, lowercase[i])
			fmt.Println(part)
			resultArr = append(resultArr, part)
			part = nil 
		} 
	}
	for index,value := range resultArr{
		if len(value) < c {
			resultArr[index] = append(resultArr[index], " ")
		}
	}
	if len(part) != 0 {
		resultArr = append(resultArr, part)
	}
	encode := make([]string,c)
	j := 0
	for j := range resultArr[j] {
		for _, v := range resultArr {
			encode = append(encode, v[j])
		}
		if j != len(resultArr[0])-1 {
			encode = append(encode, " ")
		}
	}
	return strings.Join(encode, "")
}
