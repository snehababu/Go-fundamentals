package main

import (
	"fmt"
	"time"
)

//Budget is a budget for campaign
type Budget struct {
	CampaignID string
	Balance float64 //USD
	Expires time.Time
}

func main(){
	b1 := Budget{"Kittens", 22.3, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b1)
	/*
	o/p: 
	{Kittens 22.3 2023-07-31 16:44:04.4745835 +0530 IST m=+604800.007044201}
	*/

	//If you want to add more details on the budget
	fmt.Printf("Struct1 = %#v\n", b1)

	/*
	o/p:
	{Kittens 22.3 2023-07-31 16:45:25.4993816 +0530 IST m=+604800.011294701}
main.Budget{Campaign:"Kittens", Balance:22.3, Expires:time.Date(2023, time.July, 31, 16, 45, 25, 499381600, time.Local)}
	*/

	//If you want to access single field from the Budget
	fmt.Println(b1.CampaignID)

	//You can also create structs by specifying fields by name and not by positions
	b2 := Budget {
		Balance: 19.3,
		CampaignID: "Puppies",
	}
	fmt.Printf("Struct2 = %#v\n", b2)

	var b3 Budget
	fmt.Printf("%#v\n", b3)
}   

