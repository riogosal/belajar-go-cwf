package main

import (
	"fmt"
	"strconv"
)

type Mahasiswa struct {
	Nama    string
	NIM     string
	Jurusan string
	Gender  string
	Umur    string
	Phone   int32
}

func AddPhoneNumberMhs(mhs Mahasiswa, phoneStr string) (Mahasiswa, error) {
	phoneInt, err := strconv.Atoi(phoneStr)
	if err != nil {
		return Mahasiswa{}, err // jika parsing gagal maka nilai yang ada pada Mahasiswa akan kosong
	}
	mhs.Phone = int32(phoneInt)
	return mhs, nil
}

func main() {
	mhs1 := Mahasiswa{
		Nama:    "Eliser Randa",
		NIM:     "4150101123",
		Jurusan: "Teknik Sipil",
		Gender:  "Laki-laki",
		Umur:    "22",
	}

	mhs2 := Mahasiswa{
		Nama:    "Igor",
		NIM:     "654321",
		Jurusan: "Teknik Elektro",
		Gender:  "Laki-laki",
		Umur:    "24",
	}

	daftarMahasiswa := []Mahasiswa{mhs1, mhs2}

	// Menambahkan nomor telepon ke mhs1
	mhsAddPhone, err := AddPhoneNumberMhs(mhs1, "1234567890")
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	for data, mhs := range daftarMahasiswa {
		if mhs.NIM == mhs1.NIM {
			daftarMahasiswa[data] = mhsAddPhone
		}
	}

	fmt.Println("\nDaftar Mahasiswa setelah penambahan nomor telepon:")
	for _, mhs := range daftarMahasiswa {
		fmt.Println("Nama     : ", mhs.Nama)
		fmt.Println("Nim      : ", mhs.NIM)
		fmt.Println("Jurusan  : ", mhs.Jurusan)
		fmt.Println("Gender   : ", mhs.Gender)
		fmt.Println("Umur     : ", mhs.Umur)
		fmt.Println("Telepon  : ", mhs.Phone)
	}
}
