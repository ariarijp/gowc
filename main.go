package main

import "os"

func main() {
	app := GetApp()

	app.Run(os.Args)
}
