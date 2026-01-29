# TP1 - Algorithms and Data Structures

- Infix notation to postfix notation.

## Description

The program (`infix.go`) reads mathematical expressions from standard input (stdin) and converts them to postfix notation using the Shunting-yard algorithm.

The program supports:
- Integer numbers.
- Basic arithmetic operators: Addition (`+`), Subtraction (`-`), Multiplication (`*`), Division (`/`), and Power (`^`).
- Parentheses `(` and `)` to group operations and define precedence.

## How to Run

1. Go to the project directory (`tps/tp1`).
2. Run the program: 

For example, you can run it interactively:

`go run tp1/infix.go`


And then type the expressions (one per line).

Or you can redirect an input file:

`go run tp1/infix.go < input_file.txt`

## Project Structure

- `tdas/`: Package containing the Abstract Data Type implementations (Stack and Queue) used in the algorithm.
