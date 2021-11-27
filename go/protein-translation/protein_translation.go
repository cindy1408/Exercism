package protein

import (
	"errors"
	"strings"
)

var ErrStop = errors.New("stop")
var ErrInvalidBase = errors.New("invalid base")

func FromRNA(rna string) ([]string, error) {
	coden := []string{}
	rnaString := strings.Split(rna, "")
	for i:=0; i<len(rnaString)-2; i+=3 {
		eachCoden := strings.Join(rnaString[i:i+3], "")
		coden = append(coden, eachCoden)
	}
	codenArr := []string{}
	for _, value := range coden {
		if value == "AUG" {
			codenArr = append(codenArr, "Methionine")
		} else if value == "UUU" || value == "UUC" {
			codenArr = append(codenArr, "Phenylalanine")
		} else if value == "UUA" || value == "UUG" {
			codenArr = append(codenArr, "Leucine")
		} else if value == "UCU" || value == "UCC" || value == "UCA" || value == "UCG" {
			codenArr = append(codenArr, "Serine")
		} else if value == "UAU" || value == "UAC" {
			codenArr = append(codenArr, "Tyrosine")
		} else if value == "UGU" || value == "UGC" {
			codenArr = append(codenArr, "Cysteine")
		} else if value == "UGG" {
			codenArr = append(codenArr, "Tryptophan")
		} else if value == "UAA" || value == "UAG" || value == "UGA" {
			break
		} else {
			keyArr := make(map[string]bool)
			result := []string{}
			for _, repeatCoden := range codenArr {
			if _, isRepeated := keyArr[repeatCoden]; !isRepeated {
			keyArr[repeatCoden] = true 
			result = append(result, repeatCoden)
			}
			}
			return result, ErrInvalidBase
		}
	}
	keyArr := make(map[string]bool)
	result := []string{}
	for _, repeatCoden := range codenArr {
		if _, isRepeated := keyArr[repeatCoden]; !isRepeated {
			keyArr[repeatCoden] = true 
			result = append(result, repeatCoden)
		}
	}
	return result, nil
}

func FromCodon(codon string) (string, error) {
	if codon == "AUG" {
		return "Methionine", nil
	} else if codon == "UUU" || codon == "UUC" {
		return "Phenylalanine", nil
	} else if codon == "UUA" || codon == "UUG" {
		return "Leucine", nil 
	} else if codon == "UCU" || codon == "UCC" || codon == "UCA" || codon == "UCG" {
		return "Serine", nil 
	} else if codon == "UAU" || codon == "UAC" {
		return "Tyrosine", nil
	} else if codon == "UGU" || codon == "UGC" {
		return "Cysteine", nil
	} else if codon == "UGG" {
		return "Tryptophan", nil
	} else if codon == "UAA" || codon == "UAG" || codon == "UGA" {
		return "", ErrStop
	} else {
		return "", ErrInvalidBase
	}
}