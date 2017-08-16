package main

import (
	"fmt"
)

func main() {
	result, err := Transform("https://github.com/gographics/imagick")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

}
