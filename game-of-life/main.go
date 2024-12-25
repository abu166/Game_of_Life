package main

import (
	"crunch03/elements"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	helpFlag        bool
	verboseFlag     bool
	delayMs         int = 0
	fileFlag        string
	edgesPortal     bool
	randomFlag      string
	fullscreenFlag  bool
	footprintsFlag  bool
	coloredFlag     bool
	clearScreenFlag bool
)

// Parsing flags and handling errors
func parseArgs(args []string) error {
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			flag := arg[2:]
			if strings.Contains(flag, "=") {
				parts := strings.SplitN(flag, "=", 2)
				key := parts[0]
				value := parts[1]
				switch key {
				case "delay-ms":
					var err error
					if delayMs == 0 {
						delayMs, err = strconv.Atoi(value)
					}
					if delayMs < 100 {
						err = errors.New("Numbers lesser than 100 are not supported")
					}
					if err != nil {
						fmt.Printf("Error: Invalid value for --delay-ms: %s\n", value)
						return err
					}
				case "file":
					if fileFlag != "" {
						break
					}
					if randomFlag == "" {
						fileFlag = value
					}
				case "random":
					if randomFlag != "" {
						break
					}
					if fileFlag == "" {
						randomFlag = value
					}
				default:
					fmt.Printf("Error: Unsupported flag --%s\n", key)
					return errors.New("unsupported flag")
				}
			} else {
				switch flag {
				case "help":
					helpFlag = true
				case "verbose":
					verboseFlag = true
				case "clear-screen":
					clearScreenFlag = true
				case "edges-portal":
					edgesPortal = true
				case "fullscreen":
					fullscreenFlag = true
				case "footprints":
					footprintsFlag = true
				case "colored":
					coloredFlag = true
				default:
					fmt.Printf("Error: Unsupported flag --%s\n", flag)
					return errors.New("unsupported flag")
				}
			}
		} else {
			fmt.Printf("Error: Invalid argument format '%s'\n", arg)
			return errors.New("invalid argument")
		}
	}
	if delayMs == 0 {
		delayMs = 2500
	}
	return nil
}

// main map of the game
var game elements.GameOfLife

func main() {
	err1 := parseArgs(os.Args[1:])

	if err1 != nil {
		fmt.Println("Use --help to see available options.")
		os.Exit(1)
	}
	// initizaling required inputs
	game := &elements.GameOfLife{
		Footprints:   footprintsFlag,
		Colored:      coloredFlag,
		LiveCellChar: "×",
		DeadCellChar: "·",
		VisitedChar:  "∘",
	}

	if helpFlag {
		elements.PrintHelp()
		return
	}
	var err error

	// Three different ways of filling map depending on mode

	if fileFlag != "" {
		err = game.ReadFromFile(fileFlag)
	} else if randomFlag != "" {
		err = game.GenerateRandomGrid(randomFlag)
	} else {
		err = game.ReadFromStdin()
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// adding color in case of flag
	if coloredFlag {
		game.LiveCellColor = "\033[32m"
		game.VisitedColor = "\033[33m"
	}

	// initializing visited slice for the flag

	if fullscreenFlag == true {
		game.Fullscreen()
	}

	game.Visited = make([][]bool, game.Rows)

	visited := game.Visited
	for i := range visited {
		
		visited[i] = make([]bool, game.Cols)
	}

	// main cycle for the game
	for {
		if clearScreenFlag {
			game.ClearScreen()
		}
		if verboseFlag {
			fmt.Printf("Tick: %d\n", game.Tick+1)                  // Prints Tick count
			fmt.Printf("Grid Size: %dx%d\n", game.Rows, game.Cols) // Prints grid Size height amd witdth
			fmt.Printf("Live Cells: %d\n", game.CountLiveCells())  //
			fmt.Printf("DelayMs: %dms\n\n", delayMs)
			if edgesPortal {
				game.CountLiveNeighbors(edgesPortal, game.Rows, game.Cols)
				fmt.Printf("Edges Portal: %v\n\n", edgesPortal)
			}

		}

		game.PrintGrid()
		if game.IsEmpty() {
			fmt.Println("The game is Over! All living cells are died, HE-HE-HE))")
			break
		}
		game.NextTick(edgesPortal)
		time.Sleep(time.Duration(delayMs) * time.Millisecond)
	}
}
