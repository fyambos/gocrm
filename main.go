package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrez votre nom : ")
	nom, _ := reader.ReadString('\n')
	fmt.Println("Bonjour", nom)
}
