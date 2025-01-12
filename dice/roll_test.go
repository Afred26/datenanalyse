package dice

import "fmt"

func ExampleRollSingleDieOnce() {
	fmt.Println(RollSingleDieOnce())

	//Output:
}

func ExampleRollMultipleDiceOnce() {
	fmt.Println(RollMultipleDiceOnce(1))

	//Output:
}
func ExampleRollMany() {
	fmt.Println(RollMany(1, 2))
	fmt.Println(RollMany(2, 5))
	fmt.Println(RollMany(1, 5))
	fmt.Println(RollMany(1, 5))
	fmt.Println(RollMany(1, 5))

	//Output:
}
