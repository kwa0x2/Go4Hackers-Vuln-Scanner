package main

import (
	"fmt"
	"github.com/kwa0x2/go4hackers-vuln-scanner/pkg/actions"
	"github.com/mbndr/figlet4go"
)

func main() {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		// Colors can be given by default ansi color codes...
		figlet4go.ColorRed,
		figlet4go.ColorMagenta,
		figlet4go.ColorCyan,
	
	}
	renderStr, _ := ascii.RenderOpts("Go4Hackers", options)

	fmt.Println(renderStr)

	actions.Commands()
}
