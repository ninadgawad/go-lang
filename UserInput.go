package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Fetch Input as String
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Name:")
	myname, _ := reader.ReadString('\n')
	fmt.Println(myname)

	// Fetch Input as Float
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Age:")
	myAge, _ := reader2.ReadString('\n')
	myFloatAge, _ := strconv.ParseFloat(strings.TrimSpace(myAge), 64)
	fmt.Println(myFloatAge)

}
