package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func giveMarble(giver *int, receiver *int, marble int) {
	if *giver >= marble {
		*giver = *giver - marble
		*receiver = *receiver - *giver
	} else {
		fmt.Println("Kelereng tidak cukup untuk dibagikan")
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	giver, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	receiver, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	marble, _ := strconv.Atoi(scanner.Text())
	giveMarble(&giver, &receiver, marble)
	fmt.Printf("%d \n", giver)
	fmt.Printf("%d", receiver)
}
