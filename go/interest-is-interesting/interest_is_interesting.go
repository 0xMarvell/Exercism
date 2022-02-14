package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	//panic("Please implement the InterestRate function")
	switch {
	case balance >= 0 && balance < 1000:
		return float32(0.5)
	case balance >= 1000 && balance < 5000:
		return float32(1.621)
	case balance >= 5000:
		return float32(2.475)
	default:
		return float32(3.213)
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	//panic("Please implement the Interest function")
	rate := float64(InterestRate(balance) / 100)
	return rate * balance
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	//panic("Please implement the AnnualBalanceUpdate function")
	return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	//panic("Please implement the YearsBeforeDesiredBalance function")
	var target int
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		target += 1
	}
	return target
}
