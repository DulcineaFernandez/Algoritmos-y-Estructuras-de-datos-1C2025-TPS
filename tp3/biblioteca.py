from grafo import Grafo
import heapq
from collections import deque
import random
import csv
from collections import defaultdict
from xml.sax.saxutils import escape

#CONSTANTES
BARATO = "barato"
RAPIDO = "rapido"
FRECUENCIA_INVERSA = 'frecuencia_inversa'
INVERTIR_FRECUENCIA = 1.0

def cargar_grafo_desde_csv(aeropuertos_path, vuelos_path, es_dirigido):
    grafo = Grafo(dirigido=es_dirigido)
    info_aeropuertos = {}
    ciudad_a_codigos = defaultdict(list)

    try:
        with open(aeropuertos_path, newline='', encoding='utf-8') as f:
            reader = csv.reader(f)
            for i, fila in enumerate(reader):
                linea_num = i + 1
                if len(fila) != 4:
                    raise ValueError(f"Error en aeropuertos.csv (línea {linea_num}): Se esperaban 4 columnas, se encontraron {len(fila)}. Fila: {fila}")

                ciudades_raw = fila[0].strip()
                codigo = fila[1].strip()
            
                if not ciudades_raw:
                    raise ValueError(f"Error en aeropuertos.csv (línea {linea_num}): El campo 'ciudad' no puede estar vacío. Fila: {fila}")
                if not codigo:
                    raise ValueError(f"Error en aeropuertos.csv (línea {linea_num}): El campo 'código de aeropuerto' no puede estar vacío. Fila: {fila}")

                try:
                    lat = float(fila[2].strip())
                    lon = float(fila[3].strip())
                except ValueError as e:
                    raise ValueError(f"Error en aeropuertos.csv (línea {linea_num}): Los valores de latitud o longitud no son números válidos. Fila: {fila}. Detalle: {e}")
                
                info_aeropuertos[codigo] = {
                    "ciudad": ciudades_raw,
                    "latitud": lat,
                    "longitud": lon
                }

                grafo.agregar_vertice(codigo)

                ciudad_a_codigos[ciudades_raw].append(codigo) 
    except FileNotFoundError:
        raise FileNotFoundError(f"El archivo de aeropuertos no se encontró en la ruta especificada: '{aeropuertos_path}'")
    except Exception as e:
        raise RuntimeError(f"Ocurrió un error inesperado al procesar aeropuertos.csv: {e}")
    
    try:
        with open(vuelos_path, newline='', encoding='utf-8') as f:
            reader = csv.reader(f)
            for i, fila in enumerate(reader):
                linea_num = i + 1

                if len(fila) != 5:
                    raise ValueError(f"Error en vuelos.csv (línea {linea_num}): Se esperaban 5 columnas, se encontraron {len(fila)}. Fila: {fila}")

                origen_codigo = fila[0].strip()
                destino_codigo = fila[1].strip()

                if not origen_codigo:
                    raise ValueError(f"Error en vuelos.csv (línea {linea_num}): El campo 'aeropuerto_i' (origen) no puede estar vacío. Fila: {fila}")
                if not destino_codigo:
                    raise ValueError(f"Error en vuelos.csv (línea {linea_num}): El campo 'aeropuerto_j' (destino) no puede estar vacío. Fila: {fila}")

                try:
                    tiempo = int(fila[2].strip())
                    precio = int(fila[3].strip())
                    cant_vuelos = int(fila[4].strip())
                except ValueError as e:
                    raise ValueError(f"Error en vuelos.csv (línea {linea_num}): Los valores de tiempo, precio o cant_vuelos no son enteros válidos. Fila: {fila}. Detalle: {e}")

                if origen_codigo not in info_aeropuertos:
                    raise ValueError(f"Error en vuelos.csv (línea {linea_num}): El aeropuerto de origen '{origen_codigo}' no se encontró en aeropuertos.csv. Fila: {fila}")
                if destino_codigo not in info_aeropuertos:
                    raise ValueError(f"Error en vuelos.csv (línea {linea_num}): El aeropuerto de destino '{destino_codigo}' no se encontró en aeropuertos.csv. Fila: {fila}")

                grafo.agregar_arista(origen_codigo, destino_codigo, (tiempo, precio, cant_vuelos))

    except FileNotFoundError:
        raise FileNotFoundError(f"El archivo de vuelos no se encontró en la ruta especificada: '{vuelos_path}'")
    except Exception as e:
        raise RuntimeError(f"Ocurrió un error inesperado al procesar vuelos.csv: {e}")

    return grafo, info_aeropuertos, ciudad_a_codigos

def camino_mas(grafo, condicion, ciudad_origen, ciudad_destino, ciudad_a_codigos): #O(F log A) + O(L), peor de los casos, L=F --> O(F log A + F) = O(F log A) ya que F log A crece mas rapido que F.si considero que la cantidad de ciudades que tiene un aeropuerto es pequenia como para considerarla una constante
    aeropuertos_origen = ciudad_a_codigos.get(ciudad_origen, [])
    aeropuertos_destino = set(ciudad_a_codigos.get(ciudad_destino, []))

    if not aeropuertos_origen:
        raise ValueError(f"No se encontraron aeropuertos para la ciudad origen '{ciudad_origen}'.")
    if not aeropuertos_destino:
        raise ValueError(f"No se encontraron aeropuertos para la ciudad destino '{ciudad_destino}'.")


    mejor_camino = None
    mejor_dist = float('inf')

    for origen in aeropuertos_origen:
        padres, dist = camino_minimo_dijkstra(grafo, origen, tipo=condicion)

        for destino in aeropuertos_destino:
            if dist.get(destino, float('inf')) < mejor_dist:
                mejor_dist = dist[destino]
                mejor_camino = reconstruir_camino(padres, destino)

    if mejor_camino is None:
        print(f"No hay camino entre {ciudad_origen} y {ciudad_destino}")
        return []

    return mejor_camino

def reconstruir_camino(padres, destino): #O(L), L: cant vertices, peor de los casos podria ser O(F)
    recorrido = []
    while destino is not None:
        recorrido.append(destino)
        destino = padres.get(destino)
    return recorrido[::-1]

def imprimir_camino(camino, criterio):
    if not camino:
        print("No hay camino.")
    elif criterio:
        print(" -> ".join(camino))
    else:
        print(", ".join(camino))

def bfs(grafo, origen): #O(F + A)
    if not grafo.obtener_vertices():
        raise ValueError("El grafo está vacío.")
    if origen not in grafo.obtener_vertices():
        raise ValueError(f"El vértice de origen '{origen}' no existe en el grafo.")
    visitados = set()
    padres = {}
    orden = {}

    for v in grafo:
        orden[v] = float('inf')
        padres[v] = None

    padres[origen] = None
    orden[origen] = 0
    visitados.add(origen)

    q = deque()
    q.append(origen)

    while q:
        v = q.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                orden[w] = orden[v] + 1
                visitados.add(w)
                q.append(w)
    return padres, orden

def camino_escalas(grafo, ciudad_origen, ciudad_destino, ciudad_a_codigos): #O(F + A), si considero que la cantidad de ciudades que tiene un aeropuerto es pequenia como para considerarla una constante
    codigos_origen = ciudad_a_codigos.get(ciudad_origen, [])
    codigos_destino = ciudad_a_codigos.get(ciudad_destino, [])

    if not codigos_origen:
        raise ValueError(f"No se encontraron aeropuertos para la ciudad origen '{ciudad_origen}'.")
    if not codigos_destino:
        raise ValueError(f"No se encontraron aeropuertos para la ciudad destino '{ciudad_destino}'.")
    
    mejor_camino = None

    for origen in codigos_origen:
        padres, _ = bfs(grafo, origen)
        for destino in codigos_destino:
            if destino in padres:
                camino = reconstruir_camino(padres, destino)
                if mejor_camino is None or len(camino) < len(mejor_camino):
                    mejor_camino = camino

    return mejor_camino

def mst_prim(grafo, vertice_inicial, visitados_totales):
    visitados_actuales = set()
    min_heap = []
    mst_aristas = []

    vertice = vertice_inicial
    visitados_actuales.add(vertice)
    visitados_totales.add(vertice)

    for vecino in grafo.adyacentes(vertice):
        if vecino not in visitados_actuales:
            dato_arista = grafo.peso_arista(vertice, vecino)
            if dato_arista:
                heapq.heappush(min_heap, (dato_arista[1], vertice, vecino, dato_arista))

    while min_heap:
        price, u, v, full_dato_arista = heapq.heappop(min_heap)

        if v in visitados_actuales:
            continue

        visitados_actuales.add(v)
        visitados_totales.add(v)

        if u < v:
            mst_aristas.append((u, v, full_dato_arista[0], full_dato_arista[1], full_dato_arista[2]))
        else:
            mst_aristas.append((v, u, full_dato_arista[0], full_dato_arista[1], full_dato_arista[2]))

        for vecino in grafo.adyacentes(v):
            if vecino not in visitados_actuales:
                dato_arista_vecino = grafo.peso_arista(v, vecino)
                if dato_arista_vecino:
                    heapq.heappush(min_heap, (dato_arista_vecino[1], v, vecino, dato_arista_vecino))
    
    return mst_aristas

#Ahora busca en todas las componentes, porque si algunas no quedan conectadas, el mst no las agarra, asique ahora si agarra todos los vertices
def nueva_aerolinea(grafo, salida_mst_path): # O(F log A)
    if not grafo.obtener_vertices():
        raise ValueError("El grafo está vacío.")

    vertices = []
    visitados = set()

    for vertice in grafo.obtener_vertices():
        if vertice not in visitados:
            sub_mst = mst_prim(grafo, vertice, visitados)
            vertices.extend(sub_mst)
    try:
        with open(salida_mst_path, "w") as f:
            aristas = sorted(vertices, key=lambda x: (x[0], x[1]))
            for u, v, tiempo, precio, cant_vuelos in aristas:
                f.write(f"{u},{v},{tiempo},{precio},{cant_vuelos}\n")
            print("OK")
    except IOError as e:
        raise IOError(f"Error al escribir el archivo de salida: {e}")

def grados_entrada(grafo):
    g_ent = {v: 0 for v in grafo}
    for v in grafo:
        for w in grafo.adyacentes(v):
            g_ent[w] += 1
    return g_ent

def topologico_grados(grafo):
    g_ent = grados_entrada(grafo)
    q = deque()
    resultado = []

    for v in grafo.obtener_vertices():
        if g_ent[v] == 0:
            q.append(v)

    while q:
        v = q.popleft()
        resultado.append(v)
        for w in grafo.adyacentes(v):
            g_ent[w] -= 1
            if g_ent[w] == 0:
                q.append(w) 

    if len(resultado) != len(grafo):
        raise ValueError("El grafo tiene un ciclo")
    
    return resultado

def cargar_itinerario(ruta_csv):
    try:
            with open(ruta_csv, newline='', encoding='utf-8') as archivo:
                lector = csv.reader(archivo)
                try:
                    ciudades_itinerario = next(lector)
                    if not ciudades_itinerario:
                        raise ValueError(f"El archivo de itinerario '{ruta_csv}' no contiene ciudades en la primera línea.")
                    
                    restricciones = []
                    for i, linea in enumerate(lector):
                        if len(linea) != 2:
                            raise ValueError(f"Error en itinerario.csv (línea {i+2}): Las restricciones deben tener exactamente 2 ciudades. Línea: {linea}")
                        restricciones.append(tuple(linea))
                    
                except StopIteration:
                    raise ValueError(f"El archivo de itinerario '{ruta_csv}' está vacío o no tiene el formato esperado.")
                except Exception as e:
                    raise ValueError(f"Error de formato en el archivo de itinerario '{ruta_csv}': {e}")
    except FileNotFoundError:
        raise FileNotFoundError(f"El archivo de itinerario no se encontró en la ruta especificada: '{ruta_csv}'")
    except Exception as e:
        raise RuntimeError(f"Ocurrió un error inesperado al cargar el itinerario desde '{ruta_csv}': {e}")
        
    return ciudades_itinerario, restricciones

def construir_grafo_dependencias(ciudades, restricciones):
    grafo = Grafo(dirigido=True)
    for ciudad in ciudades:
        grafo.agregar_vertice(ciudad)
    for desde, hasta in restricciones:
        grafo.agregar_vertice(desde)
        grafo.agregar_vertice(hasta)
        grafo.agregar_arista(desde, hasta)
    return grafo

def itinerario(grafo, ruta_itinerario_csv, ciudad_a_codigos):
    ciudades_itinerario, restricciones = cargar_itinerario(ruta_itinerario_csv)
    grafo_dep = construir_grafo_dependencias(ciudades_itinerario, restricciones)
    orden_ciudades = topologico_grados(grafo_dep)
    print(", ".join(orden_ciudades))

    for i in range(len(orden_ciudades) - 1):
        ciudad_origen = orden_ciudades[i]
        ciudad_destino = orden_ciudades[i + 1]

        camino_aeropuertos = camino_escalas(grafo, ciudad_origen, ciudad_destino, ciudad_a_codigos)
        
        if not camino_aeropuertos:
            raise ValueError(f"No se pudo encontrar un camino entre '{ciudad_origen}' y '{ciudad_destino}' para el itinerario.")
        else:
            imprimir_camino(camino_aeropuertos, True)

def camino_minimo_dijkstra(grafo, origen, destino=None, tipo=RAPIDO):
    if tipo not in {RAPIDO, BARATO, FRECUENCIA_INVERSA}:
        raise ValueError("El tipo debe ser 'rapido', 'barato'")

    dist = {v: float('inf') for v in grafo.obtener_vertices()}
    padre = {v: None for v in grafo.obtener_vertices()}
    dist[origen] = 0
    
    heap = []
    heapq.heappush(heap, (0, origen))

    while heap:
        actual_dist, v = heapq.heappop(heap)

        if actual_dist > dist[v]:
            continue

        for w in grafo.adyacentes(v):
            dato_arista = grafo.peso_arista(v,w)
            if tipo == FRECUENCIA_INVERSA:
                cant_vuelos = dato_arista[2] 
                if cant_vuelos <= 0:
                    peso = float('inf') 
                else:
                    peso = INVERTIR_FRECUENCIA / cant_vuelos
            elif tipo == RAPIDO:
                peso = dato_arista[0]
            elif tipo == BARATO:
                peso = dato_arista[1]

            if peso == float('inf'):
                continue
            
            nueva_dist = actual_dist + peso

            if nueva_dist < dist[w]:
                dist[w] = nueva_dist
                padre[w] = v
                heapq.heappush(heap, (nueva_dist, w))

    return padre, dist

#Como uso Dijkstra para buscar el camino min, necesito invertir la cant de vuelos, pa que me de el "mayor", pero me da el menor por ser de camino min, aprovecho el bug y al invertirlo me da el mayor
def centralidad(grafo, k):
    if not grafo.obtener_vertices():
        raise ValueError("El grafo está vacío.")
    
    cent = {v: 0 for v in grafo.obtener_vertices()} # Inicializa la centralidad de todos los vértices a 0

    for v in grafo.obtener_vertices():
        padres, distancias = camino_minimo_dijkstra(grafo, v, tipo=FRECUENCIA_INVERSA) 
        dist_vertices = []
        for w_node in grafo.obtener_vertices():
            if distancias[w_node] != float('inf'):
                dist_vertices.append((distancias[w_node], w_node))
        
        dist_vertices.sort(key=lambda item: item[0], reverse=True)
        vertices_ordenados = [item[1] for item in dist_vertices]

        cent_aux = {x: 0 for x in grafo.obtener_vertices()}

        for w_node in vertices_ordenados:
            if padres.get(w_node) is not None and padres[w_node] != v:
                cent_aux[padres[w_node]] += (1 + cent_aux[w_node])

        for w_node in grafo.obtener_vertices():
            if w_node != v:
                cent[w_node] += cent_aux[w_node]

    return heapq.nlargest(k, cent.keys(), key=lambda node: cent[node])


def exportar_kml(nombreKml, camino, info_aeropuertos):
    with open(nombreKml, "w") as archivoKml:
        archivoKml.write('<?xml version="1.0" encoding="UTF-8"?>\n')
        archivoKml.write('<kml xmlns="http://earth.google.com/kml/2.1">\n')
        archivoKml.write('<Document>\n')
        archivoKml.write('<name>ULTIMA RUTA</name>\n')
        archivoKml.write('<description>Muestra de la ultima ruta guardada</description>\n')

        for aeropuerto in camino:
            longitud = info_aeropuertos[aeropuerto]["longitud"]
            latitud = info_aeropuertos[aeropuerto]["latitud"]
            archivoKml.write('<Placemark>\n')
            archivoKml.write(f'<name>{aeropuerto}</name>\n')
            archivoKml.write('<Point>\n')
            archivoKml.write(f'<coordinates>{longitud},{latitud}</coordinates>\n')
            archivoKml.write('</Point>\n')
            archivoKml.write('</Placemark>\n')

        for i in range(len(camino) - 1):
            aeropuerto1 = camino[i]
            aeropuerto2 = camino[i + 1]
            longitud1 = info_aeropuertos[aeropuerto1]["longitud"]
            latitud1 = info_aeropuertos[aeropuerto1]["latitud"]
            latitud2 = info_aeropuertos[aeropuerto2]["latitud"]
            longitud2 = info_aeropuertos[aeropuerto2]["longitud"]
            info2 = info_aeropuertos[aeropuerto2]
            archivoKml.write('    <Placemark>\n')
            archivoKml.write('      <LineString>\n')
            archivoKml.write('        <coordinates>')
            archivoKml.write(f'{longitud1},{latitud1} {longitud2},{latitud2}')
            archivoKml.write('</coordinates>\n')
            archivoKml.write('      </LineString>\n')
            archivoKml.write('    </Placemark>\n')
        archivoKml.write('</Document>\n')
        archivoKml.write('</kml>\n')

    print("Ok")