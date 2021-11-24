package strand

import "strings"

func ToRNA(dna string) string {
	rna := []string{}
	dnaArr := strings.Split(dna, "")
	for _, v := range dnaArr {
		switch v {
		case "G": 
			rna = append(rna, "C")
		case "C": 
			rna = append(rna, "G")
		case "T": 
			rna = append(rna, "A")
		case "A": 
			rna = append(rna, "U")
		}
	}
	return strings.Join(rna, "")
}
