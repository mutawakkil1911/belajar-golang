package main

import "fmt"

func dataDiri(nama string, umur int, alamat string) {
	fmt.Println("Nama saya", nama, "Umur", umur, "Tahun", "Alamat", alamat)
}

func main() {
	nama := "Mutawakkil"
	umur := 18
	alamat := "Sumenep"

	dataDiri(nama, umur, alamat)
}
