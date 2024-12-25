package elements

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func (g *GameOfLife) ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default: // Unix-like systems
		fmt.Print("\033[H\033[2J")
	}
}
