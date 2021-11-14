package dna

import (
	"errors"
)

type Histogram map[rune]int 

type DNA []rune

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{rune('A'):0, rune('C'):0, rune('G'):0, rune('T'):0}
	if len(d) == 0 {
		return h, nil 
	}
	for _, dna := range d {
		switch dna {
		case 'A': 
		count := h['A']
			h['A']= count +1 
			count += 1 
		case 'C':
			count := h['C']
			h['C'] = count +1
			count += 1 
		case 'G':
			count := h['G']
			h['G'] = count+1
			count += 1 
		case 'T':
			count := h['T']
			h['T'] = count +1
			count += 1 
		default: 
			return nil, errors.New("invalid dna")
		}
	}
	return h, nil
}
