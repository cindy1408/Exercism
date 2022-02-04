package railfence

func Encode(message string, rails int) string {
	codes := make([][]rune, rails)
	rail := 0
	down := true
	for _, r := range message {
		codes[rail] = append(codes[rail], r)
		if down {
			rail++
			if rail == rails {
				down = false
				rail -= 2
			}
		} else {
			rail--
			if rail < 0 {
				down = true
				rail += 2
			}
		}
	}
	encoded := make([]rune, 0, len(message))
	for _, line := range codes {
		encoded = append(encoded, line...)
	}
	return string(encoded)
}

func Decode(message string, rails int) string {
	length := len(message)
	result := make([]rune, length)
	rail := 0
	initDist := rails*2 - 2
	dist := initDist
	idx := -initDist
	for _, r := range message {
		idx += dist
		if idx >= length {
			rail++
			idx = rail
			dist = initDist - 2*rail
			if dist == 0 {
				dist = initDist
			}
			result[idx] = r
		} else {
			result[idx] = r
			dist = initDist - dist
			if dist == 0 {
				dist = initDist
			}
		}
	}
	return string(result)
}
