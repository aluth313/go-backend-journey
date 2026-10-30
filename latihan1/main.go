package main

import (
	"errors"
	"fmt"
)

func hitungTotal(harga []int, diskonPersen int, ongkir int) int {
	total := 0
	for _, h := range harga {
		total += h
	}
	if diskonPersen > 0 {
		total -= total * diskonPersen / 100
	}
	if ongkir > 0 {
		total += ongkir
	}
	return total
}

func cekDiskon(diskonPersen int) error {
	if diskonPersen > 100 {
		return errors.New("diskon lebih dari 100")
	}
	return nil
}

func main() {
	keranjang := []int{25000, 40000, 15000}
	diskonPersen := 10
	ongkir := 10000
	if err := cekDiskon(diskonPersen); err != nil {
		fmt.Println("Validasi gagal: ", err)
		return
	}
	bayar := hitungTotal(keranjang, diskonPersen, ongkir)
	fmt.Printf("Total bayar: Rp%d\n", bayar)
}