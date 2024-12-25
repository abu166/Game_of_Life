package elements

import (
	"bufio"
	"fmt"
	"os"
)

func (g *GameOfLife) ReadFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	return g.ReadInput(scanner)
}
