package rajaongkir

import "fmt"

func SeparateCourierList(list []string) (word string) {
	if len(list) == 1 {
		word = fmt.Sprintf("%s", list[0])
	} else {
		for i, _ := range list {
			if i == 0 {
				word = fmt.Sprintf("%s:", list[i])
			} else if i == (len(list) - 1) {
				word = fmt.Sprintf("%s%s", word, list[i])
			} else {
				word = fmt.Sprintf("%s%s:", word, list[i])
			}
		}
	}
	return
}
