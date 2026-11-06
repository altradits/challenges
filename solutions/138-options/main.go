package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}
	var opts uint32
	for _, arg := range os.Args[1:] {
		if len(arg) >= 2 && arg[0] == '-' {
			if arg[1] == 'h' {
				fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
				return
			}
			for i := 1; i < len(arg); i++ {
				if arg[i] < 'a' || arg[i] > 'z' {
					fmt.Println("Invalid Option")
					return
				}
			}
			for i := 1; i < len(arg); i++ {
				opts |= 1 << (arg[i] - 'a')
			}
		} else if arg == "-" {
			fmt.Println("Invalid Option")
			return
		}
		// Arguments that are not options are ignored.
	}
	fmt.Printf("%08b %08b %08b %08b\n",
		byte(opts>>24), byte(opts>>16), byte(opts>>8), byte(opts))
}
