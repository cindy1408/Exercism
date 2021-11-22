package scale

import (
	"fmt"
	"strings"
)


func Scale(tonic, interval string) []string {
	sharps := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	flats := []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	steps := strings.Split(interval, "")
	
	format := strings.Split(tonic, "")
	upper := strings.ToUpper(format[0])
	var newFormat string 
	if len(format) == 1 {
		newFormat = strings.ToUpper(format[0])
	} else {
		newFormat = upper + format[1]
	}


	result := []string{newFormat}
 
	if tonic == "C" || tonic == "G" || tonic == "D" || tonic == "A" || tonic == "E" || tonic == "B" || tonic == "F#" ||  tonic == "a" || tonic == "e" || tonic == "b" || tonic == "f#" || tonic == "c#" || tonic == "g#" || tonic == "d#" {
		//interval is 0 
		if len(steps) == 0 {
			for i:=0; i< len(sharps)-2; i++ {
				if sharps[i] == result[len(result)-1] && len(result) < 12 {
					result = append(result, sharps[i+1])
				}
				if i == len(sharps)-3 && len(result) < 12 {
					i = 0 
				}
			}
		} else {
			// interval is given 
			for _, step := range steps {
				fmt.Println("this is the step for: ", tonic, step)
				if step == "M" {
					// major
					for i:=0; i< len(sharps)-3; i++ {
						if result[len(result)-1] == sharps[i] && len(result) < len(steps) {
							result = append(result, sharps[i+2])
							fmt.Println("this is major")
							break
						}
						if i == len(sharps)-2 && len(result) < len(steps) {
							i = 0 
						}
					} 
				} else if step == "m" {
					// minor 
					for i:=0; i< len(sharps)-3; i++ {
						if result[len(result)-1] == sharps[i] && len(result) < len(steps) {
							result = append(result, sharps[i+1])
							fmt.Println("this is minor")
							break
						}
						if i == len(sharps)-2 && len(result) < len(steps) {
							i = 0 
						}
					} 
				} else {
					// augment
					for i:=0; i< len(sharps)-3; i++ {
						if result[len(result)-1] == sharps[i] && len(result) < len(steps) {
							result = append(result, sharps[i+3])
							fmt.Println("this is augmented")
							break
						}
						if i == len(sharps)-2 && len(result) < len(steps) {
							i = 0 
						}
					} 
				}
			}
		}
	} else {
		if len(steps) == 0 {
			for i:=0; i< len(flats)-2; i++ {
				if flats[i] == result[len(result)-1] && len(result) < 12 {
					result = append(result, flats[i+1])
				}
				if i == len(flats)-3 && len(result) < 12 {
					i = 0 
				}
			}
		} else {
			for _, step := range steps {
				fmt.Println(step)
				if step == "M" {
					// major
					for i:=0; i< len(flats)-3; i++ {
						if result[len(result)-1] == flats[i] && len(result) < len(steps) {
							result = append(result, flats[i+2])
							fmt.Println("this is major")
							break
						}
					} 
				} else if step == "m" {
					// minor 
					for i:=0; i< len(flats)-3; i++ {
						if result[len(result)-1] == flats[i] && len(result) < len(steps) {
							result = append(result, flats[i+1])
							fmt.Println("this is minor")
							break
						}
					} 
				} else {
					for i:=0; i< len(flats)-3; i++ {
						if result[len(result)-1] == flats[i] && len(result) < len(steps) {
							result = append(result, flats[i+3])
							fmt.Println("this is augmented")
							break
						}
					} 
				}
			}
		}
	}
	return result

}