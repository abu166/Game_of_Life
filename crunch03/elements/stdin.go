package elements

import (
	"bufio"
	"fmt"
	"os"
)

func (g *GameOfLife) ReadFromStdin() error {
	fmt.Println("Enter the height and width for the map:")
	scanner := bufio.NewScanner(os.Stdin)
	return g.ReadInput(scanner)
}
