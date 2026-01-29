class Grafo:
    """
    Representación de un grafo utilizando un diccionario de diccionarios.
    Permite crear grafos dirigidos o no dirigidos.
    """
    def __init__(self, dirigido=False):
        self.vertices = {}
        self.dirigido = dirigido

    # Agrega un vértice al grafo. Si el vértice ya existe, no hace nada.
    def agregar_vertice(self, vertice):
        if vertice not in self.vertices:
            self.vertices[vertice] = {}

    # Elimina un vértice del grafo y todas las aristas asociadas a él.
    # Lanza un error si el vértice no existe en el grafo.
    # Lanza un error si el grafo esta vacio.
    def sacar_vertice(self, vertice):
        if self.es_vacio():
            raise ValueError("El grafo esta vacio")
        if not self.existe_vertice(vertice):
            raise ValueError(f"El vértice '{vertice}' no existe en el grafo.")
        del self.vertices[vertice]
        for adyacentes in self.vertices.values():
            adyacentes.pop(vertice, None)

    # Agrega una arista entre dos vértices. Si el grafo es no dirigido,
    #  también agrega la arista en la dirección opuesta.
    # Lanza error si origen o destino no existe.
    # Lanza error si el grafo no existe.
    def agregar_arista(self, origen, destino, dato_arista=1):
        if self.es_vacio():
            raise ValueError("El grafo esta vacio")
        if not self.existe_vertice(origen) or not self.existe_vertice(destino):
            raise ValueError("Uno de los vértices no existe")
        self.vertices[origen][destino] = dato_arista
        if not self.dirigido:
            self.vertices[destino][origen] = dato_arista


    # Elimina una arista entre dos vértices. Si el grafo es no dirigido,
    #  también elimina la arista en la dirección opuesta.
    # Lanza error si origen o destino no existe.
    # Lanza error si la arista a sacar no existe.
    # Lanza error si el grafo esta vacio.
    def sacar_arista(self, origen, destino):
        if self.es_vacio():
            raise ValueError("El grafo esta vacio")
        if not self.existe_vertice(origen) or not self.existe_vertice(destino):
            raise ValueError("Uno de los vértices no existe")
        if destino not in self.vertices[origen]:
            raise ValueError(f"La arista de '{origen}' a '{destino}' no existe.")
        self.vertices[origen].pop(destino, None)
        if not self.dirigido:
            self.vertices[destino].pop(origen, None)

    # Devuelve un conjunto de todos los vértices adyacentes a un vértice dado.
    # Lanza un error si el grafo esta vacio.
    def adyacentes(self, vertice):
        if not self.existe_vertice(vertice):
            raise ValueError("El vértice no existe")
        return set(self.vertices[vertice].keys())


    # Devuelve un iterador de tuplas (vecino, peso) para los vértices adyacentes.
    # Lanza un error si el grafo esta vacio.
    def adyacentes_con_pesos(self, vertice):
        if not self.existe_vertice(vertice):
            raise ValueError("El vértice no existe")
        return self.vertices[vertice].items()

    # Verifica si existe una arista entre dos vértices.
    # Lanza error si el grafo esta vacio, o lanza un error si alguno de los vertices no existe.
    def estan_unidos(self, v1, v2):
        if self.es_vacio():
            raise ValueError("El grafo esta vacio")
        if not self.existe_vertice(v1) or not self.existe_vertice(v2):
            raise ValueError("Uno de los vertices no existe")
        return v2 in self.vertices.get(v1, {})

    # Devuelve true si el vertice pertenece al grafo.
    def existe_vertice(self, vertice):
        return vertice in self.vertices

    # Devuelve una lista de todos los vértices en el grafo.
    def obtener_vertices(self):
        return list(self.vertices.keys())

    # Permite iterar directamente sobre los vértices del grafo (ej. for v in grafo:).
    def __iter__(self):
        return self.iterar_vertices()

    # Generador que produce cada vértice del grafo.
    def iterar_vertices(self):
        for v in self.vertices:
            yield v
    
    # Devuelve el peso de la arista entre dos vértices.
    # Lanza error si la arista no existe o si el grafo esta vacio.
    def peso_arista(self, origen, destino):
        if not self.estan_unidos(origen, destino):
            raise ValueError("La arista no existe")
        return self.vertices[origen][destino]
    
    # Devuelve la cantidad de vertices en el grafo.
    def __len__(self):
        return len(self.vertices)
    
    # Devuelve true si no tiene vertices, false en caso contrario.
    def es_vacio(self):
      return len(self.vertices) == 0