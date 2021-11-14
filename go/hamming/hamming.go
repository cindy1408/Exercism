package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	aArr := strings.Split(a, "")
	bArr := strings.Split(b, "")
	length := bArr
	if len(aArr) != len(bArr){
		return 0, errors.New("error")
	}
	count := 0
	for index := range length {
		if aArr[index] != bArr[index] {
			count ++
		}
	}
	return count, nil 
}
