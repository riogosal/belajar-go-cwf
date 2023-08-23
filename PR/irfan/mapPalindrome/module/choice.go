package module

import (
	"app-palindrome/model"
	"fmt"
)

func ChoiceMenu() model.Palindrome {

	for {
		fmt.Println("Masukkan kata Palindrome/kata serupa :")
		var choice string
		fmt.Scanln(&choice)

		var tampungData string = ""

		countWords := len(choice) - 1

		for i := 0; i <= countWords; countWords-- {
			convKata := choice[countWords]
			tampungData += string(convKata)
		}

		if choice == tampungData {
			fmt.Println("kata anda adalah Palindrome ^__^")

			return model.Palindrome{
				Name: tampungData,
			}

		} else {
			fmt.Println("bukan kata palindrome -___-")
		}
	}
}

func CreateData(dataP map[model.Palindrome]bool) {
	returnData := ChoiceMenu()
	dataP[returnData] = true

}
