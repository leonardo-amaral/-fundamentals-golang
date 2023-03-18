package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	buyTv50 := trab1 && trab2
	buyTv32 := trab1 != trab2
	buyIceDream := trab1 || trab2

	return buyTv50, buyTv32, buyIceDream
}

func main() {
	tv50, tv32, iceDream := compras(true, true)
	fmt.Printf("TV50: %t, TV32: %t, ICEDREAM: %t",
		tv50, tv32, !iceDream)
}
