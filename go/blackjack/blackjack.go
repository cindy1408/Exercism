package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch(card){
        case "ace":
    		return 11
        case "two": 
    		return 2
        case "three":
    		return 3
        case "four":
    		return 4
        case "five": 
			return 5
        case "six":
    		return 6
        case "seven":
    		return 7
        case "eight":
    		return 8
        case "nine":
    		return 9 
        case "ten":
    		return 10
        case "jack":
    		return 10
        case "queen":
    		return 10
        case "king": 
    		return 10
        default:
    		return 0
    }
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	cardOne := ParseCard(card1)
    cardTwo := ParseCard(card2)
    if cardOne + cardTwo == 21 {
        return true
    }
	return false
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if isBlackjack {
		if dealerScore < 10 {
			return "W"
		}
		return "S"
	}
	return "P"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
    result := ""
	switch {
        case handScore < 12:
    		result = "H"
        case handScore >11 && handScore < 17 && dealerScore < 7:
    		result = "S"
        case handScore > 16: 
    		result = "S"
        case handScore >11 && handScore < 17 && dealerScore > 6:
    		result = "H"
        default: 
    		result = "S"
    }
	return result
}
