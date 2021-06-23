package main

import (
	"fmt"
	"strings"
)

func main() {

	// mengecek keberadaan sebuah string
	fmt.Println(strings.Contains("Alfi Nurindiana", "Alfi"))  //True
	fmt.Println(strings.Contains("Alfi Nurindiana", "Pichu")) //false

	// memisah string dan menjadikan slace
	fmt.Println(strings.Split("Alfi Nurindiana", " "))

	// menjadikan lower case
	fmt.Println(strings.ToLower("Alfi Nurindiana"))

	// menjadikan upper case
	fmt.Println(strings.ToUpper("Alfi Nurindiana"))

	// menghapus kanan dan kiri string
	fmt.Println(strings.Trim("   Alfi Nurindiana   ", " "))

	// mengubah string
	fmt.Println(strings.ReplaceAll("Alfi Pichu Alfi Pichu Alfi Pichu", "Pichu", "Nurindiana"))
}
