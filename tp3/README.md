# TP3 - Flight System Implementation (FlyCombi)

This project consists of a flight query system using graphs. The program allows you to find minimum paths between airports (by price, time, or stops), calculate airport centrality, generate a flight network for a new airline (MST), plan itineraries, and export routes to KML.

## Project Files

*   **`flycombi.py`**: Main program that processes standard input and executes commands.
*   **`biblioteca.py`**: Contains the implementation of graph algorithms (Dijkstra, BFS, Prim, Topological Sort, etc.) and helper functions.
*   **`grafo.py`**: Implementation of the Graph data structure (Graph ADT) using adjacency lists (dictionaries).

## Execution

The program runs from the command line, receiving two CSV files as arguments: one with airport information and another with flight information.

python3 flycombi.py <airports_file.csv> <flights_file.csv>


Example:

python3 flycombi.py aeropuertos.csv vuelos.csv

## Available Commands

The program reads commands from standard input (stdin). The available commands are:

### 1. Best Path (`camino_mas`)

Finds the minimum path between two cities based on a criterion.

`camino_mas <rapido|barato> <origin_city>,<destination_city>`

**Example**: `camino_mas rapido Buenos Aires,Madrid`

### 2. Path with Fewer Stops (`camino_escalas`)

Finds the path with the fewest number of stops (flights) between two cities.

`camino_escalas <origin_city>,<destination_city>`

**Example**: `camino_escalas Buenos Aires,Tokio`

### 3. Centrality (`centralidad`)

Shows the `n` most important airports (highest centrality) in the network.

`centralidad <n>`

**Example**: `centralidad 10`

### 4. New Airline (`nueva_aerolinea`)

Generates a file with the minimum flight network needed to connect all airports (Minimum Spanning Tree).

`nueva_aerolinea <output_filename>`

**Example**: `nueva_aerolinea red_minima.csv`

### 5. Itinerary (`itinerario`)

Calculates a visiting order for cities respecting order restrictions given in a file.

`itinerario <itinerary_file.csv>`

**Example**: `itinerario itinerario_turista.csv`

### 6. Export KML (`exportar_kml`)

Exports the last calculated route to a KML file to be viewed in Google Earth.

`exportar_kml <filename.kml>`

**Example**: `exportar_kml ruta.kml`

### 7. Exit (`salir`)

Ends the program execution.

## Considerations

*   For `camino_mas` and `camino_escalas`, if there are multiple airports in a city, the program searches for the best path considering all possible combinations.
*   `centralidad` uses the inverse frequency of flights to determine importance.
