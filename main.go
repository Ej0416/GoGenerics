package main

import (
	"errors"
	"fmt"
	"time"
)

// ----------------------------------------- lesson 1

// func getLast[T any](s []T) T {
// 	var zeroVal T;
// 	if len(s) == 0 {
// 		return zeroVal
// 	}

// 	return s[len(s)-1]
// }

//  ------------------------------------------ lesson 2

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	if balance < newItem.GetCost(){
		return nil, 0.0, errors.New("insufficient funds")
	}

	newList := append(oldItems, newItem)
	newUserBalance := balance - newItem.GetCost();

	return newList, newUserBalance, nil
}

// don't edit below this line
type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}

func main(){
	fmt.Println("app started")
}