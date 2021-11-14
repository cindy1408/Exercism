package letter

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
		m[r]++ 
	}
	// passing m to the channel (sending)
	channel <- m 
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	concurrent := FreqMap{}
	channel := make(chan FreqMap)
	for _, eachString := range(l) {
		go CalculateFrequency(channel, eachString)
		// if concurrent was not empty... eg l was not an array of strings, then we'll pass the channel to a new frequency Map 
		if len(concurrent) > 0 {
			newFrequencyMap := <- channel
			// we'll iterate through the new map with assigning concurrent's index to the corrosponding value
			for v, k := range(newFrequencyMap){
				concurrent[v] += k
			}
		} else {
			// sending the channel into concurrent (recieving end)
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