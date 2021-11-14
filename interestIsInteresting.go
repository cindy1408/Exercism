package main

import "fmt"

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	interest := 0.0
	if balance < 0 {
		interest = 3.213
	} else if balance < 1000 && balance >= 0 {
		interest = 0.5
	} else if balance >= 1000 && balance < 5000{
		interest = 1.621
	} else if balance >= 5000 {
		interest = 2.475
	}
	fmt.Println(float32(interest))
	return float32(interest)
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	interestRate := float64(InterestRate(balance))
	interest := (interestRate/100)*balance
	fmt.Println(interest)
	return interest
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	interestToAdd := Interest(balance)
	totalBalance := balance + interestToAdd
	fmt.Println(totalBalance)
	return totalBalance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	yearsNeeded := 0
	totalBalance := balance

	for i:=0; i< int(targetBalance); i++ {
		if totalBalance < targetBalance {
			updatedBalance := totalBalance
			totalBalance = AnnualBalanceUpdate(updatedBalance)
			fmt.Println(totalBalance)
			// nextYear := AnnualBalanceUpdate(currentYear)
			yearsNeeded = yearsNeeded +1 
		} 
	}
	fmt.Println(yearsNeeded)
	return yearsNeeded
}


func main() {
	// InterestRate(200.75)
	// Interest(200.75)
	Interest(34600.800000)
	AnnualBalanceUpdate(200.75)
	AnnualBalanceUpdate(898124017.826243)
	Interest(34600.800000)
	// balance := 200.75
	// targetBalance := 214.88
	// YearsBeforeDesiredBalance(balance, targetBalance)
}