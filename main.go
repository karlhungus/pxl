package main

import (
	"fmt"
	"github.com/karlhungus/lib_pxl"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <filename>...\n\n", os.Args[0])
		fmt.Println("Close the image with <ESC> or by pressing 'q'.")
		os.Exit(1)
	}

	lib_pxl.Init()
	defer lib_pxl.Close()
	for i := 1; i < len(os.Args); i++ {
		lib_pxl.DisplayFile(os.Args[i])
	}
}
