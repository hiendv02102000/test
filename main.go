package main

import (
	"fmt"
)

func stringRepeat(listName1, listName2 []string) (result []string) {
	m := make(map[string]int)
	for i, v := range listName1 {
		m[v] = i
	}
	for _, v := range listName2 {
		_, ok := m[v]
		if ok {
			result = append(result, v)
		}
	}
	return
}

func main() {

	listName1 := []string{"anna", "devoe", "flash", "iris", "kante", "kovakic"}
	listName2 := []string{"kante", "namor", "elsa", "stark", "devoe"}

	fmt.Println(stringRepeat(listName1, listName2))

}
