# TP2 - Flight Management System

- File handling
- Data structures
- Command processing

## Description

This project implements a system to manage flight information.
The program loads flight data from CSV files and allows different queries on the flight board.

The program (`algueiza.go`) reads commands from standard input (stdin).
It uses data structures to store the information and answer the commands efficiently.

## How to Run

1. Go to the project directory (`tps/tp2`).
2. Run the main program:

`go run tp2/algueiza.go`

3. Enter commands using standard input.
You can also use an input file:

go run tp2/algueiza.go < input.txt

> Available Commands

`agregar_archivo <filename>`
Loads flights from the given CSV file.
If a flight already exists, its information is updated.
Example: agregar_archivo flights.csv

`info_vuelo <flight_code>`
Shows all the information of a flight with the given code.
Example: info_vuelo 1234

`prioridad_vuelos <k>`
Shows the k flights with the highest priority.
Example: prioridad_vuelos 10

`ver_tablero <k> <mode> <from> <to>`
Shows k flights ordered by date.

<mode>: asc (ascending) or desc (descending).

<from>: Start date (YYYY-MM-DD or YYYY-MM-DDTHH:MM:SS).

<to>: End date.
Example: ver_tablero 5 asc 2023-10-01 2023-10-31

`siguiente_vuelo <origin> <destination> <date>`
Shows the next available flight from origin to destination after the given date.
Example: siguiente_vuelo EZE MAD 2023-10-15

`borrar <from> <to>`
Removes all flights with dates between <from> and <to>.
The deleted flights are printed to standard output.
Example: borrar 2020-01-01 2020-12-31

## Project Structure

`tp2/`: Main program (`algueiza.go`)

`tp2/comandos/`: Command logic

`tp2/tablero/`: Board ADT

`tp2/vuelo/`: Flight ADT

`tdas/`: Data structures (Hash, Heap, Dictionary, etc.)