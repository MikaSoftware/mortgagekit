package mortgagekit

import (
    "fmt"
    "github.com/rhymond/go-money"
)

type MortgageCalculator struct {
    totalAmount *money.Money
    downPaymentAmount *money.Money
    amortizationYear int64
    annualInterestRate float64
    paymentFrequency int64
    compoundingPeriod int64
    firstPaymentDate string
    currency string
}


func New(
    totalAmount *money.Money,
    downPaymentAmount *money.Money,
    amortizationYear int64,
    annualInterestRate float64,
    paymentFrequency int64,
    compoundingPeriod int64,
    firstPaymentDate string,
    currency string) MortgageCalculator {
    mc := MortgageCalculator {
        totalAmount,
        downPaymentAmount,
        amortizationYear,
        annualInterestRate,
        paymentFrequency,
        compoundingPeriod,
        firstPaymentDate,
        currency,
    }
    return mc
}


func ExampleHello() bool {
    fmt.Println("hello")
    // Output: hello
    return true
}
