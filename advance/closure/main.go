package main

import "fmt"

func main() {
	useGiftCard1 := activateGiftCard()
	useGiftCard2 := activateGiftCard()

	fmt.Println(useGiftCard1(34))
	fmt.Println(useGiftCard2(66))
	fmt.Println(useGiftCard2(2))
}

// here, since the returned function debitFunc depends on
// variable amount the userGiftCard1 also contains its own version
// of amount intially 100 along with the debitFunc, this is called as closure
func activateGiftCard() func(int) int {
	amount := 100

	debitFunc := func(debitAmount int) int {
		amount -= debitAmount
		return amount
	}
	return debitFunc
}
