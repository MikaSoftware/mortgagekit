package mortgagekit_test

import (
    "testing"
    "github.com/rhymond/go-money"

    "github.com/mikasoftware/mortgagekit"
)


func TestNew(t *testing.T) {
    totalAmount := money.New(250000, "CAD")
    downPaymentAmount := money.New(50000, "CAD")

    calculator := mortgagekit.New(
        totalAmount,
        downPaymentAmount,
        25,
        0.04,
        mortgagekit.Month,
        mortgagekit.SemiAnnual,
        "2008-01-01",
        "CAD",
    )

    if &calculator == nil {
        t.Error("Failed initializing calculator using `New` function.")
    }
}
