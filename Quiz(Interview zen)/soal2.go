package main

import "fmt"

func main() {
	var arr = [5][2]string{
		[2]string{"Row3", "Row5"},
		[2]string{"Row1", "Row2"},
		[2]string{"Row4", ""},
		[2]string{"", "Row6"},
		[2]string{"Row7", "Row8"},
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < 2; j++ {
			if j == 0 && arr[i][j] != "" {
				fmt.Printf("%s = S , ", arr[i][j])
			} else if j == 1 && arr[i][j] != "" {
				fmt.Printf("%s = A", arr[i][j])
			}
		}
		fmt.Println()
	}

}
