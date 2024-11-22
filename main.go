package main

import (
	"chatthing/server"
	"embed"
	"fmt"
)

const PORT string = ":5000"

//go:embed static
var staticFiles embed.FS

func main() {
	fmt.Println(logo)
	fmt.Println("Listening on the port", PORT)
	var err = server.Start(PORT, staticFiles)
	if err != nil {
		fmt.Println(err)
	}
}

const logo string = `
  ______ _                    _     _             
 / _____) |          _   _   | |   (_)            
| /     | | _   ____| |_| |_ | | _  _ ____   ____ 
| |     | || \ / _  |  _)  _)| || \| |  _ \ / _  |
| \_____| | | ( ( | | |_| |__| | | | | | | ( ( | |
 \______)_| |_|\_||_|\___)___)_| |_|_|_| |_|\_|| |
                                           (_____|
`
