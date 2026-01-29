#!/usr/bin/python3
# -*- coding: utf-8 -*-
from grafo import Grafo
from biblioteca import (
    camino_mas,
    camino_escalas,
    centralidad,
    nueva_aerolinea,
    itinerario,
    cargar_grafo_desde_csv,
    imprimir_camino,
    exportar_kml,
)

import sys

BARATO = "barato"
RAPIDO = "rapido"
CAMINO_MAS = "camino_mas"
CAMINO_ESCALAS = "camino_escalas"
CENTRALIDAD = "centralidad"
NUEVA_AEROLINEA = "nueva_aerolinea"
ITINERARIO = "itinerario"
SALIR = "salir"
EXPORTAR_KML = "exportar_kml"

def chequearArgumentos(comando, argumento, cantidad):
    if len(argumento) != cantidad:
        print("Error en la cantidad de elementos")
        return False

    if comando == CAMINO_MAS:
        if argumento[0] not in ("rapido", "barato"):
            print("Error")
            return False
    elif comando == CENTRALIDAD:
        try:
            if int(argumento[0]) < 1:
                print("Error: El grado de centralidad debe ser un número entero positivo.")
                return False
        except ValueError:
            print("Error: El grado de centralidad debe ser un número entero.")
            return False

    return True

def leer_comando(grafo, linea, info_aeropuertos, ciudad_a_codigos, ultima_ruta):
    linea = linea.strip()
    if not linea:
        return ultima_ruta # Devuelve ultima_ruta para que pueda ser actualizada en main

    partes = linea.split(' ', 1)
    comando = partes[0]
    raw_argumentos_str = "" 

    if len(partes) > 1:
        raw_argumentos_str = partes[1].strip()

    argumentos_procesados = [] 
    if raw_argumentos_str:
        argumentos_procesados = [arg.strip() for arg in raw_argumentos_str.split(',')]

    if comando == CAMINO_MAS:
        if chequearArgumentos(comando, argumentos_procesados, 3):
            condicion, origen, destino = argumentos_procesados
            camino = camino_mas(grafo, condicion, origen, destino, ciudad_a_codigos)
            imprimir_camino(camino, True)
            if camino: # Solo actualiza si se encontró un camino
                ultima_ruta = camino
    elif comando == CAMINO_ESCALAS:
        if chequearArgumentos(comando, argumentos_procesados, 2):
            origen, destino = argumentos_procesados
            camino = camino_escalas(grafo, origen, destino, ciudad_a_codigos)
            imprimir_camino(camino, True)
            if camino: # Solo actualiza si se encontró un camino
                ultima_ruta = camino
    elif comando == CENTRALIDAD:
        if chequearArgumentos(comando, argumentos_procesados, 1):
            camino = centralidad(grafo, int(argumentos_procesados[0]))
            imprimir_camino(camino, False)
            # No se guarda en ultima_ruta ya que no es una ruta aérea
    elif comando == NUEVA_AEROLINEA:
        if chequearArgumentos(comando, argumentos_procesados, 1):
            nueva_aerolinea(grafo, argumentos_procesados[0])
    elif comando == ITINERARIO:
        if chequearArgumentos(comando, argumentos_procesados, 1):
            itinerario(grafo, argumentos_procesados[0], ciudad_a_codigos)
    elif comando == EXPORTAR_KML:
        if chequearArgumentos(comando, argumentos_procesados, 1) and ultima_ruta:
            exportar_kml(argumentos_procesados[0], ultima_ruta, info_aeropuertos)
        else:
            print("Error: No hay una ruta previa para exportar o argumentos incorrectos.")
    elif comando == SALIR:
        sys.exit(0)
    else:
        print(f"Comando '{comando}' no reconocido.")
    
    return ultima_ruta


def main():

    if len(sys.argv) != 3:
        print("Uso: ./flycombi <archivo_aeropuertos.csv> <archivo_vuelos.csv>")
        sys.exit(1)

    archivo_aeropuertos = sys.argv[1]
    archivo_vuelos = sys.argv[2]
    
    grafo, info_aeropuertos, ciudad_a_codigos = cargar_grafo_desde_csv(archivo_aeropuertos, archivo_vuelos, False)
    ultima_ruta = None
    for linea in sys.stdin:
        ultima_ruta = leer_comando(grafo, linea, info_aeropuertos, ciudad_a_codigos, ultima_ruta) # Actualiza ultima_ruta

if __name__ == "__main__":
    main()