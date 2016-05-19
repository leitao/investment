package main

import "fmt"
import "os"
import "strings"
import "encoding/json"

type setOfInv struct {
	Investments []string
}

func ShowHelp() {
	fmt.Println("Help mode")
	os.Exit(0)
}

func listInvInit() {
	fmt.Println("Listing Investments")
}

func listInv(set setOfInv) {
	array := set.Investments

	for i := range array {
		fmt.Println(array[i])
	}
}

func encode(set setOfInv) {
	b, _ := json.Marshal(set)
	fmt.Println(string(b))
}

func addInv(investNew string) {
	fmt.Println("Adding a new investment item name ", investNew)
	set := setOfInv{}
	array := append(set.Investments, investNew)
	set.Investments = array
	listInv(set)
	encode(set)
}

func Parseinput() {
	Args := os.Args
	for item := range Args {
		if strings.Compare(Args[item], "-h") == 0 {
			ShowHelp()
		}
		if strings.Compare(Args[item], "-l") == 0 {
			listInvInit()
		}
		if strings.Compare(Args[item], "-a") == 0 {
			if item+1 <= len(Args) {
				addInv(Args[item+1])
			} else {
				ShowHelp()
			}
		}
	}

}

func main() {
	fmt.Println("Breno's investment control system")

	Parseinput()

}
