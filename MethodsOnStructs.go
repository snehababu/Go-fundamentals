package main

import (
	"fmt"
	"time"
)

//Budget us the budget for campaign
type Budget struct {
	CampaignId string
	Balance float64 //USD
	Expires time.Time
}

/*
b is th reciever of type Budget
Timeleft is the name of the method
We take the expiration time and substract the current time and return it
*/
func (b Budget) Timeleft() time.Duration {
	return b.Expires.Sub(time.Now().UTC())
}	

/*
update method:
this method is going to change the balance inside the budget
inside the budget we are going to change the balance so we are going to pass a pointer of budget
this is known as pointer reciever
Everytime you change the struct, you are going to pass a pointer reciever
*/
func (b *Budget) Update(sum float64){
	b.Balance += sum
}	

func main() {
	//we create a budget that expires in 7 days and print out how much time is left
	b := Budget{"Kittens", 22.3, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b.Timeleft())

	b.Update(10.5)
	fmt.Println(b.Balance)
}