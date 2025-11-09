package main

import (
	"fmt"
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

// func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {

// 	if balance < newItem.GetCost(){
// 		return nil, 0.0, errors.New("insufficient funds")
// 	}

// 	newUserBalance := balance - newItem.GetCost();
// 	newList := append(oldItems, newItem)

// 	return newList, newUserBalance, nil
// }

// // don't edit below this line
// type lineItem interface {
// 	GetCost() float64
// 	GetName() string
// }

// type subscription struct {
// 	userEmail string
// 	startDate time.Time
// 	interval  string
// }

// func (s subscription) GetName() string {
// 	return fmt.Sprintf("%s subscription", s.interval)
// }

// func (s subscription) GetCost() float64 {
// 	if s.interval == "monthly" {
// 		return 25.00
// 	}
// 	if s.interval == "yearly" {
// 		return 250.00
// 	}
// 	return 0.0
// }

// type oneTimeUsagePlan struct {
// 	userEmail        string
// 	numEmailsAllowed int
// }

// func (otup oneTimeUsagePlan) GetName() string {
// 	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
// }

// func (otup oneTimeUsagePlan) GetCost() float64 {
// 	const costPerEmail = 0.03
// 	return float64(otup.numEmailsAllowed) * costPerEmail
// }

// ------------------------------------------------ lesson 7
// type biller[C customer] interface {
// 	Charge(C) bill
// 	Name() string
// }

// // don't edit below this line

// type userBiller struct {
// 	Plan string
// }

// func (ub userBiller) Charge(u user) bill {
// 	amount := 50.0
// 	if ub.Plan == "pro" {
// 		amount = 100.0
// 	}
// 	return bill{
// 		Customer: u,
// 		Amount:   amount,
// 	}
// }

// func (sb userBiller) Name() string {
// 	return fmt.Sprintf("%s user biller", sb.Plan)
// }

// type orgBiller struct {
// 	Plan string
// }

// func (ob orgBiller) Name() string {
// 	return fmt.Sprintf("%s org biller", ob.Plan)
// }

// func (ob orgBiller) Charge(o org) bill {
// 	amount := 2000.0
// 	if ob.Plan == "pro" {
// 		amount = 3000.0
// 	}
// 	return bill{
// 		Customer: o,
// 		Amount:   amount,
// 	}
// }

// type customer interface {
// 	GetBillingEmail() string
// }

// type bill struct {
// 	Customer customer
// 	Amount   float64
// }

// type user struct {
// 	UserEmail string
// }

// func (u user) GetBillingEmail() string {
// 	return u.UserEmail
// }

// type org struct {
// 	Admin user
// 	Name  string
// }

// func (o org) GetBillingEmail() string {
// 	return o.Admin.GetBillingEmail()
// }

// ================================== chapter 15: enums (or lacked there of)

// func (a *analytics) handleEmailBounce(em email) error {
// 	if err := em.recipient.updateStatus(em.status); err != nil{
// 		return fmt.Errorf("error updating user status: %w", err)
// 	}

// 	if err := a.track(em.status); err != nil {
// 		return fmt.Errorf("error tracking user bounce: %w", err)
// 	}

// 	em.recipient.updateStatus(em.status)
// 	a.track(em.status)
// 	return nil
// }

// +++++++++++++++++++++++++++++++++++++ apperently the onle way to have "enums" in go

type emailStatus int;

const (
	emailBounced emailStatus = iota
	emailInvalid
	emailDelivered
	emailOpened
)


func main() {
	fmt.Println("app started")
}
