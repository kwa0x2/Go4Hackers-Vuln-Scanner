package main

import (
	"github.com/kwa0x2/go4hackers-vuln-scanner/pkg/actions"
	"github.com/mbndr/figlet4go"
	"fmt"
)

func main() {
	ascii := figlet4go.NewAsciiRender()
	renderStr, _ := ascii.Render("Go4Hackers")

	fmt.Println(renderStr)

	actions.Commands()
}