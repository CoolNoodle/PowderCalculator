package main

import "fmt"

func main(){
		//# of powders used to max Legendary Equips
	const LLw = 405.6 // Leg wep powder
	const LUw = 746.5 // Unique wep powder
	const LLa = 312.0 // Leg armor powder
	const LUa = 574.23 // Unique armor powder
	const powMulti = 99.0

		//a = amount of armor powder
	Ua := 6.0
	La := 3.0
//	Uw := 10.0
//	Lw := 0.0
		// = 99x powder 
//	LegWepPowder := Lw*powMulti
//	UniWepPowder := Uw*powMulti
	UniArmPowder := Ua*powMulti
	LegArmPowder := La*powMulti

		//Maxing with Unique armor powder
	Mua := UniArmPowder/LUa 
	Mla := LegArmPowder/LLa
	Sum := Mua + Mla
	fmt.Println("You can max out", Sum , "legendary armors.")
	fmt.Println("\t(Unique:", Mua, ")\n\t(Legendary:", Mla, ")")
}

