# Game of Life - Go Implementation

## Overview

This project is an implementation of Conway's Game of Life written in Go. The Game of Life is a cellular automaton devised by mathematician John Conway. It's a zero-player game, meaning the evolution of the game depends solely on the initial state, requiring no further input once the simulation begins. The game follows a set of simple rules to evolve a grid of live and dead cells across time ticks.

In this project, you can control various aspects of the game through command-line flags, such as animation speed, grid size, grid source, and visualization options. You can read more about Conway's Game of Life [here](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) or play it [here](https://playgameoflife.com).

## Features

### Live and Dead Cells Representation:
- Use `×` for live cells.
- Use `·` for dead cells.

### Flags:
- `--help`: Displays help information about the program's capabilities.
- `--verbose`: Shows detailed information about the simulation (grid size, number of ticks, number of live cells, etc.).
- `--delay-ms=X`: Sets the animation speed in milliseconds (default is 2500ms).
- `--file=X`: Reads the initial grid from a specified file.
- `--random=WxH`: Generates a random grid with width `W` and height `H`.
- `--edges-portal`: Enables portal edges where cells that exit one side of the grid appear on the opposite side.
- `--fullscreen`: Adjusts the grid to fit the terminal size with empty cells.
- `--footprints`: Adds visual traces of visited cells, displayed as `∘`.
- `--colored`: Adds color to live cells and traces if footprints are enabled.

## Game Rules
1. **Underpopulation**: Any live cell with fewer than two live neighbors dies.
2. **Survival**: Any live cell with two or three live neighbors survives to the next generation.
3. **Overpopulation**: Any live cell with more than three live neighbors dies.
4. **Reproduction**: Any dead cell with exactly three live neighbors becomes a live cell.

The game will continue running until there are no live cells left.

## How to Run

To compile and run the program, use the following command:

```bash
go run main.go [options]
```

Basic Run:
```bash
go run main.go
```

You will be prompted to input the grid directly via the terminal or through flags.



Using a File:
```bash
go run main.go --file=path/to/grid.txt
```

Loads the initial grid from a file.


Verbose Mode:
```bash
go run main.go --verbose
```

Displays detailed information about the simulation at each tick.



Random Grid:
```bash
go run main.go --random=5x5
```

Generates a random 5x5 grid and runs the simulation.



Portal Edges:
```bash
go run main.go --edges-portal
```

Enables the feature where cells that exit the grid appear on the opposite side.



Full-Screen Mode:
```bash
go run main.go --fullscreen
```

Adjusts the grid to fit the terminal size.



Clear-screen Mode:
```bash
go run main.go --clear-screen
```

Clears the previous grid.



Input Format

The input grid consists of two symbols:

'#' represents a live cell.
'.' represents a dead cell.

## Example Input
```bash
4 4
....
.#..
.##.
....
```

## Example Output
```bash
· · · ·
· × · ·
· × × ·
· · · ·
```

## Error Handling

- If the input is invalid or the grid is not well-formed, the program will display an error message.
- If no valid grid is provided through the terminal, file, or random grid flag, the program will prompt the user for correct input.
- If conflicting flags are used, the first one will be accepted, and others will be ignored.
- Any unrecognized flags or arguments will result in an error.


### Example of Invalid Input
```bash
Error: invalid input.
```


## Bonus Features

- Portal Edges: The --edges-portal flag allows cells to "wrap around" the grid, appearing on the opposite side when they exit one side.
- Footprints: The --footprints flag adds traces of visited cells for better visualization.
- Color: The --colored flag enhances the display by adding color to live cells and footprints.


# Assessment Criteria

## Functionality
- Correctness: The program reads input correctly and applies the Game of Life rules to evolve the grid.
- Edge Cases: Handles minimum-size grids (3x3) and larger grids. Terminates correctly when no live cells are left or when invalid input is detected.
- Flag Handling: The program handles all flags correctly and displays appropriate error messages for invalid inputs.


## Code Quality
- Readability: The code is well-organized and easy to understand, with clear functions and comments.
- Efficiency: The program uses efficient algorithms and data structures to ensure smooth performance.
- Usability
- User Interaction: Provides clear instructions and error messages, ensuring a smooth experience.
- Game Mechanics: Correctly displays the grid evolution according to the Game of Life rules.
