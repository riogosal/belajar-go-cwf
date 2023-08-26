package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type PilihanMenu int

const (
	TampilkanFilm PilihanMenu = 1
	BuatFilm      PilihanMenu = 2
	UpdateFilm    PilihanMenu = 3
	HapusFilm     PilihanMenu = 4
	Exit          PilihanMenu = 5
)

func TampilkanMenu() PilihanMenu {
	user_input := bufio.NewScanner(os.Stdin)
	fmt.Println("1. Tampilkan Film")
	fmt.Println("2. Tambah Data Film")
	fmt.Println("3. Update Film")
	fmt.Println("4. Hapus Film")
	fmt.Println("5. Save and Exit")

	fmt.Print("Pilih Menu : ")
	user_input.Scan()

	pilihan_menu, _ := strconv.Atoi(user_input.Text())
	return PilihanMenu(pilihan_menu)
}
