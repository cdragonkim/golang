package main

import (
	"fmt"
)

func main() {
	i := 9
	switch {
	case i < 0:
		fmt.Println(i, "는 음수입니다.")
	case i == 0:
		fmt.Println(i, "는 O입니다.")
	case i > 0:
		fmt.Println(i, "는 양수입니다.")

	}
}
