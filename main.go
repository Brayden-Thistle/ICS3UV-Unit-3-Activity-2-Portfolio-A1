/*
* @author Brayden Thistle
* @version 1.0.0
* @date 2025-11-19
* @fileoverfiew this program sings the user happy birthday
*/

package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)
func main () {
	var PersonsName string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is your name? ")
	PersonsName, _ = reader.ReadString('\n')
	PersonsName = strings.TrimSpace(PersonsName)

	fmt.Println("Happy birthday to you. Happy birthday to you. Happy birthday, dear " + PersonsName + ". happy birthday to you.")

	fmt.Println("\nDone")
}