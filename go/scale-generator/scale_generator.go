package scale

import (
	"strings"
)

func Scale(tonic, interval string) []string {
	intervalSplit := strings.Split(interval, "")
	scale := []string{}
	major := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#"}
	minor := []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A","Bb", "B", "C", "Db"}

		if tonic == "C" || tonic == "G" || tonic == "D" || tonic == "A" || tonic == "E" || tonic == "B" || tonic == "F#" || tonic == "e" || tonic == "b" || tonic == "f#" || tonic == "c#" || tonic == "g#" || tonic == "d#" {
			// major 	
			if len(intervalSplit) == 0 {
				scale = append(scale, tonic)
				for i:=0; i<14; i++ {
					if major[i] == scale[len(scale)-1] {
						if i < len(major){
							scale = append(scale, major[i+1])
						}
					}
				}
				return scale 
			}
			for _, interval := range intervalSplit {
					if interval == "M" {
						for i, note := range major {
							if scale[len(scale)-1] == note {
								scale = append(scale, scale[i+2])
							}
						}
					} else {
						for i, note := range major {
							if scale[len(scale)-1] == note {
								scale = append(scale, scale[i+1])
							}
						}
					}
			}
		} else {
			// minor 
			if len(intervalSplit) == 0 {
				scale = append(scale, tonic)
				for j:=0; j<len(minor)-1; j++ {
					if minor[j] == scale[len(scale)-1] && len(scale) < 12 {
						if j < len(minor){
							scale = append(scale, minor[j+1])
						}
					}
				}
				return scale 
			}
			for _, interval := range intervalSplit {
					if interval == "M" {
						for i, note := range minor {
							if scale[len(scale)-1] == note {
								scale = append(scale, scale[i+2])
							}
						}
					} else {
						for i, note := range minor {
							if scale[len(scale)-1] == note {
								scale = append(scale, scale[i+1])
							}
						}
					}
			}
		}
	return scale
}