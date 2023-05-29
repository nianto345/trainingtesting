package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// membuat class mobil
type mobil struct {
	//mendefinisikan beberapa attribute
	Ban   int32
	Pintu int32
}

func (m *mobil) inBan() {
	fmt.Println("Ban", m.Ban)
	fmt.Println("Pintu", m.Pintu)

}

func main() {
	//melalukann instansi
	var kendaraan mobil = mobil{Ban: 4, Pintu: 2}

	//memanggil method
	kendaraan.inBan()
	// Jenis ban
	var JenisBan []string = []string{"karet", "kayu", "besi"}
	//Method Ban
	for _, JBan := range JenisBan {
		fmt.Println(JBan)
	}

	var pintu = "knock"
	switch pintu {
	case "open":
		fmt.Println("Beep")
	case "knock":
		fmt.Println("Tutut")
	}

}

func TestSomething(t *testing.T) {

	assert.True(t, true, "True is true!")

}
