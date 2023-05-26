package main

import "fmt"

type Address struct {
	city, province, country string
}

func ChangeCountryToIndonesia(alamat *Address) {
	alamat.country = "indonesia"
}

func main() {
	var address1 Address = Address{"surabaya", "jawa tengah", "indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.city = "bandung"

	*address2 = Address{"jakarta", "DKI JAKARTA", "indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := Address{
		city:     "jakarta",
		province: "Dki jakarta",
		country:  "",
	}
	ChangeCountryToIndonesia(&address4)
	fmt.Println(address4)
}
