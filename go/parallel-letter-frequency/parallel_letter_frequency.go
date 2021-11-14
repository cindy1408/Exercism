package letter

import "fmt"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func CalculateFrequency(channel chan FreqMap, eachString string) {
	m := FreqMap{}
	for _, r := range eachString {
		fmt.Println("this is m: ", m)
		fmt.Println("this is r: ", r)
		m[r]++ 
	}
	fmt.Println("this is what channel is receiving: ", m)
	channel <- m 
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	concurrent := FreqMap{}
	channel := make(chan FreqMap)
	for _, eachString := range(l) {
		go CalculateFrequency(channel, eachString)
		fmt.Println("THIS IS CONCURRENT: ", concurrent)
		if len(concurrent) > 0 {
			newFrequencyMap := <- channel
			for v, k := range(newFrequencyMap){
				concurrent[v] += k
			}
		} else {
			concurrent = <- channel
		}
	}
	return concurrent
}



func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}