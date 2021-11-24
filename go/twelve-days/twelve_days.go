package twelve

import (
	"fmt"
	"strings"
)

func Verse(i int) string {
	days := []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
	gifts := []string{"a Partridge in a Pear Tree", "two Turtle Doves", "three French Hens", "four Calling Birds", "five Gold Rings", "six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming"}
	
	firstPart := fmt.Sprintf("On the %v day of Christmas my true love gave to me: ", days[i-1])	
	secondPartArr := []string{}

	for j:=i-1 ; j >=0 ; j-- {
		if j == 0 && i != 1 {
			endGift := "and " + gifts[0]
			secondPartArr = append(secondPartArr, endGift)
		} else {
			secondPartArr = append(secondPartArr, gifts[j])
		}
	}
	secondPart := strings.Join(secondPartArr, ", ")
	verse := firstPart + secondPart + "."
	return verse
}

func Song() string {
	numDays := []int{1,2,3,4,5,6,7,8,9,10,11,12}
	songArr := []string{}
	for _, num := range numDays {
		songArr = append(songArr, Verse(num))
	}
	song := strings.Join(songArr, "\n")
	return song
}
