package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    var stars = []string{}
    for i := 0; i < numStarsPerLine; i++ {
        stars = append(stars, "*")
    }
	message := strings.Join(stars, "")
	return message + "\n" + welcomeMsg + "\n" + message
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	messageSlice := strings.Split(oldMsg, "")
    var cleanMessage = []string{}
    for i := 0; i < len(messageSlice); i++ {
        if messageSlice[i] != "*" {
            cleanMessage = append(cleanMessage, messageSlice[i])
        }
    }
	message := strings.Join(cleanMessage, "")
    message = strings.TrimSpace(message)
	return message
}
