package main

import "fmt"

type Alamat struct {
	Kecamatan, Kota, Provinsi string
}

func main() {
	// Pass by Value
	fmt.Println("pass by value")
	fmt.Println("---")

	alamat1 := Alamat{"Gapura", "Sumenep", "Jawa Timur"}
	alamat2 := alamat1

	alamat2.Kecamatan = "Batang-Batang"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

	// pass by reference
	fmt.Println("pass by reference")
	fmt.Println("---")

	alamat3 := &alamat1

	alamat3.Kecamatan = "Bluto"

	fmt.Println(alamat1)
	fmt.Println(alamat3)

	// operator pointer *

	// fmt.Println("operator pointer *")
	// fmt.Println("---")

	// alamat4 := &alamat1

	// *alamat4 = Alamat{"Nangkaan", "Bondowoso", "Jawa Timur"}

	// fmt.Println(alamat1)
	// fmt.Println(alamat2)
	// fmt.Println(alamat3)
	// fmt.Println(alamat4)

	// function new

	fmt.Println("function new")
	fmt.Println("---")

	alamat5 := new(Alamat)

	alamat5.Kota = "Bondowoso"

	fmt.Println(alamat5)

}
