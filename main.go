package main

import (
	"fmt"
	"latihan/repositories"
	"os"
)

func main() {
	args := os.Args
	id := args[1]

	biodata := repositories.GetBiodataById(id)
	fmt.Println("Id", biodata.Id)
	fmt.Println("Name", biodata.Name)
	fmt.Println("Address", biodata.Address)
	fmt.Println("Pekerjaan", biodata.Pekerjaan)
	fmt.Println("Alasan", biodata.Alasan)
}
