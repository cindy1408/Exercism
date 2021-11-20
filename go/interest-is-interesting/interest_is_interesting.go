package interest

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
	return float32(interest)
}


func Interest(balance float64) float64 {
	interestRate := float64(InterestRate(balance))
	interest := (interestRate/100)*balance
	return interest
}

func AnnualBalanceUpdate(balance float64) float64 {
	interestToAdd := Interest(balance)
	totalBalance := balance + interestToAdd
	return totalBalance
}

func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	yearsNeeded := 0
	totalBalance := balance

	for i:=0; i< int(targetBalance); i++ {
		if totalBalance < targetBalance {
			updatedBalance := totalBalance
			totalBalance = AnnualBalanceUpdate(updatedBalance)
			yearsNeeded = yearsNeeded +1 
		} 
	}
	return yearsNeeded
}
