package mortgagekit

import (
    "errors"
    // "fmt"
    "github.com/rhymond/go-money"
)

type MortgageCalculator struct {
    totalLoan *money.Money
    downPayment *money.Money
    amortizationYear int64
    annualInterestRate float64
    paymentFrequency int64
    compoundingPeriod int64
    firstPaymentDate string
    currency string
}

// Initializer function which will create `MortgageCalculator` object.
func New(
    totalLoan *money.Money,
    downPayment *money.Money,
    amortizationYear int64,
    annualInterestRate float64,
    paymentFrequency int64,
    compoundingPeriod int64,
    firstPaymentDate string,
    currency string) MortgageCalculator {
    mc := MortgageCalculator {
        totalLoan,
        downPayment,
        amortizationYear,
        annualInterestRate,
        paymentFrequency,
        compoundingPeriod,
        firstPaymentDate,
        currency,
    }
    return mc
}

func (mc *MortgageCalculator) GetPercentOfLoanFinanced() (float64, error) {
    // Defesnive code.
    if mc.totalLoan.Amount() == 0 {
        return 0, errors.New("Total amount cannot be zero.")
    }

    // Calculate our loan princinple.
    remainingLoan, err := mc.totalLoan.Subtract(mc.downPayment)

    // Defensive Code.
    if err != nil {
        return 0, err
    }

    // Calculate the remaining rate.
    remainingLoanFloat := float64(remainingLoan.Amount())
    totalLoanFloat := float64(mc.totalLoan.Amount())
    amountFinancedPercent := remainingLoanFloat / totalLoanFloat

    // Return the remaining percent.
    return amountFinancedPercent * 100, nil
}
