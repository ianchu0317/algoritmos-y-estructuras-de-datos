from grafo import Grafo
from biblioteca import comunidades, obtener_vertices_entrada

def test_grafo():
    print("=== Iniciando tests del Grafo ===\n")
    
    # Test 1: Crear grafo no dirigido
    print("Test 1: Crear grafo no dirigido")
    grafo_nd = Grafo(es_dirigido=False)
    print("✅ Grafo no dirigido creado\n")
    
    # Test 2: Agregar vértices
    print("Test 2: Agregar vértices")
    grafo_nd.agregar_vertice("A")
    grafo_nd.agregar_vertice("B")
    grafo_nd.agregar_vertice("C")
    assert grafo_nd.hay_vertice("A")
    assert grafo_nd.hay_vertice("B")
    assert grafo_nd.hay_vertice("C")
    print("✅ Vértices A, B, C agregados\n")
    
    # Test 3: Intentar agregar vértice duplicado
    print("Test 3: Intentar agregar vértice duplicado")
    try:
        grafo_nd.agregar_vertice("A")
        print("❌ ERROR: Debería lanzar excepción")
    except ValueError as e:
        print(f"✅ Excepción correcta: {e}\n")
    
    # Test 4: Agregar aristas
    print("Test 4: Agregar aristas")
    grafo_nd.agregar_arista("A", "B", 5)
    grafo_nd.agregar_arista("B", "C", 3)
    assert grafo_nd.hay_arista("A", "B")
    assert grafo_nd.hay_arista("B", "A")  # Grafo no dirigido
    assert grafo_nd.hay_arista("B", "C")
    print("✅ Aristas A-B (peso 5) y B-C (peso 3) agregadas\n")
    
    # Test 5: Verificar pesos
    print("Test 5: Verificar pesos de aristas")
    assert grafo_nd.peso_arista("A", "B") == 5
    assert grafo_nd.peso_arista("B", "A") == 5  # Simétrico
    assert grafo_nd.peso_arista("B", "C") == 3
    print("✅ Pesos correctos\n")
    
    # Test 6: Intentar agregar arista duplicada
    print("Test 6: Intentar agregar arista duplicada")
    try:
        grafo_nd.agregar_arista("A", "B", 10)
        print("❌ ERROR: Debería lanzar excepción")
    except ValueError as e:
        print(f"✅ Excepción correcta: {e}\n")
    
    # Test 7: Obtener adyacentes
    print("Test 7: Obtener adyacentes")
    adyacentes_b = grafo_nd.adyacentes("B")
    assert "A" in adyacentes_b
    assert "C" in adyacentes_b
    print(f"✅ Adyacentes de B: {adyacentes_b}\n")
    
    # Test 8: Obtener todos los vértices
    print("Test 8: Obtener todos los vértices")
    vertices = grafo_nd.obtener_vertices()
    assert set(vertices) == {"A", "B", "C"}
    print(f"✅ Vértices: {vertices}\n")
    
    # Test 9: Borrar arista
    print("Test 9: Borrar arista")
    grafo_nd.borrar_arista("A", "B")
    assert not grafo_nd.hay_arista("A", "B")
    assert not grafo_nd.hay_arista("B", "A")  # También se borra en grafo no dirigido
    print("✅ Arista A-B borrada\n")
    
    # Test 10: Borrar vértice
    print("Test 10: Borrar vértice")
    grafo_nd.borrar_vertice("C")
    assert not grafo_nd.hay_vertice("C")
    assert "C" not in grafo_nd.adyacentes("B")
    print("✅ Vértice C borrado\n")
    
    # Test 11: Grafo dirigido
    print("Test 11: Crear grafo dirigido")
    grafo_d = Grafo(es_dirigido=True)
    grafo_d.agregar_vertice("X")
    grafo_d.agregar_vertice("Y")
    grafo_d.agregar_arista("X", "Y", 7)
    assert grafo_d.hay_arista("X", "Y")
    assert not grafo_d.hay_arista("Y", "X")  # Grafo dirigido
    print("✅ Grafo dirigido: X→Y (peso 7)\n")
    
    # Test 12: Intentar acceder a vértice inexistente
    print("Test 12: Intentar acceder a vértice inexistente")
    try:
        grafo_d.adyacentes("Z")
        print("❌ ERROR: Debería lanzar excepción")
    except ValueError as e:
        print(f"✅ Excepción correcta: {e}\n")
    
    # Test 13: Intentar agregar arista con vértice inexistente
    print("Test 13: Intentar agregar arista con vértice inexistente")
    try:
        grafo_d.agregar_arista("X", "Z", 5)
        print("❌ ERROR: Debería lanzar excepción")
    except ValueError as e:
        print(f"✅ Excepción correcta: {e}\n")
    
    # Test 14: Peso de arista inexistente
    print("Test 14: Peso de arista inexistente")
    try:
        grafo_d.peso_arista("Y", "X")
        print("❌ ERROR: Debería lanzar excepción")
    except ValueError as e:
        print(f"✅ Excepción correcta: {e}\n")
    
    # Test 15: Crear grafo con vértices iniciales
    print("Test 15: Crear grafo con vértices iniciales")
    grafo_inicial = Grafo(es_dirigido=False, vertices=["P", "Q", "R"])
    assert grafo_inicial.hay_vertice("P")
    assert grafo_inicial.hay_vertice("Q")
    assert grafo_inicial.hay_vertice("R")
    print("✅ Grafo creado con vértices P, Q, R\n")
    
    print("=== ✅ TODOS LOS TESTS BÁSICOS PASARON EXITOSAMENTE ===\n")


def test_vertices_entrada():
    """
    Tests para obtener_vertices_entrada()
    """
    print("=== Tests de vertices_entrada ===\n")
    
    # Test 1: Grafo dirigido simple
    print("Test 1: Grafo dirigido simple")
    grafo = Grafo(es_dirigido=True)
    grafo.agregar_vertice("A")
    grafo.agregar_vertice("B")
    grafo.agregar_vertice("C")
    grafo.agregar_arista("A", "B")
    grafo.agregar_arista("A", "C")
    grafo.agregar_arista("B", "C")
    
    # A --> B
    # A --> C
    # B --> C
    
    vertices_entrada = obtener_vertices_entrada(grafo)
    
    assert vertices_entrada["A"] == []  # A no recibe aristas
    assert set(vertices_entrada["B"]) == {"A"}  # B recibe de A
    assert set(vertices_entrada["C"]) == {"A", "B"}  # C recibe de A y B
    print(f"✅ Vértices de entrada correctos: {vertices_entrada}\n")
    
    # Test 2: Grafo con ciclo
    print("Test 2: Grafo con ciclo")
    grafo2 = Grafo(es_dirigido=True)
    grafo2.agregar_vertice("X")
    grafo2.agregar_vertice("Y")
    grafo2.agregar_vertice("Z")
    grafo2.agregar_arista("X", "Y")
    grafo2.agregar_arista("Y", "Z")
    grafo2.agregar_arista("Z", "X")
    
    # X --> Y --> Z --> X (ciclo)
    
    vertices_entrada2 = obtener_vertices_entrada(grafo2)
    
    assert set(vertices_entrada2["X"]) == {"Z"}  # X recibe de Z
    assert set(vertices_entrada2["Y"]) == {"X"}  # Y recibe de X
    assert set(vertices_entrada2["Z"]) == {"Y"}  # Z recibe de Y
    print(f"✅ Ciclo detectado correctamente: {vertices_entrada2}\n")
    
    print("=== ✅ TESTS DE vertices_entrada PASARON ===\n")


def test_comunidades():
    """
    Tests para detección de comunidades con Label Propagation
    """
    print("=== Tests de Label Propagation (Comunidades) ===\n")
    
    # Test 1: Grafo con 2 comunidades claramente separadas
    print("Test 1: Dos comunidades separadas")
    grafo = Grafo(es_dirigido=True)
    
    # Comunidad 1: A, B, C (muy conectados entre sí)
    grafo.agregar_vertice("A")
    grafo.agregar_vertice("B")
    grafo.agregar_vertice("C")
    grafo.agregar_arista("A", "B")
    grafo.agregar_arista("A", "C")
    grafo.agregar_arista("B", "C")
    grafo.agregar_arista("C", "A")
    grafo.agregar_arista("C", "B")
    
    # Comunidad 2: D, E (conectados entre sí)
    grafo.agregar_vertice("D")
    grafo.agregar_vertice("E")
    grafo.agregar_arista("D", "E")
    grafo.agregar_arista("E", "D")
    
    """
    Grafo:
    A <--> B
    ↕      ↓
    C <----┘
    
    D <--> E
    """
    
    comun = comunidades(grafo)
    
    # Debe haber 2 comunidades
    assert len(comun) == 2, f"Esperado 2 comunidades, obtuvo {len(comun)}"
    
    # Verificar que A, B, C están juntos
    comunidad_abc = None
    for c in comun:
        if "A" in c:
            comunidad_abc = c
            break
    
    assert comunidad_abc is not None, "Comunidad de A no encontrada"
    assert {"A", "B", "C"}.issubset(comunidad_abc), f"A, B, C deberían estar juntos. Comunidad: {comunidad_abc}"
    
    # Verificar que D, E están juntos (en comunidad diferente)
    comunidad_de = None
    for c in comun:
        if "D" in c:
            comunidad_de = c
            break
    
    assert comunidad_de is not None, "Comunidad de D no encontrada"
    assert {"D", "E"}.issubset(comunidad_de), f"D, E deberían estar juntos. Comunidad: {comunidad_de}"
    
    print(f"✅ Comunidades detectadas correctamente:")
    for i, c in enumerate(comun, 1):
        print(f"   Comunidad {i}: {sorted(c)}")
    print()
    
    
    # Test 2: Grafo lineal (cadena)
    print("Test 2: Cadena lineal (podría formar 1 comunidad)")
    grafo2 = Grafo(es_dirigido=True)
    grafo2.agregar_vertice("1")
    grafo2.agregar_vertice("2")
    grafo2.agregar_vertice("3")
    grafo2.agregar_vertice("4")
    grafo2.agregar_arista("1", "2")
    grafo2.agregar_arista("2", "3")
    grafo2.agregar_arista("3", "4")
    
    # 1 --> 2 --> 3 --> 4
    
    comun2 = comunidades(grafo2)
    print(f"   Cadena lineal formó {len(comun2)} comunidad(es):")
    for i, c in enumerate(comun2, 1):
        print(f"   Comunidad {i}: {sorted(c)}")
    
    # En cadena, puede variar (no determinista)
    # Solo verificamos que devuelve algo válido
    assert len(comun2) >= 1, "Debe haber al menos 1 comunidad"
    total_vertices = sum(len(c) for c in comun2)
    assert total_vertices == 4, f"Debe tener 4 vértices en total, tiene {total_vertices}"
    print("✅ Cadena procesada correctamente\n")
    
    
    # Test 3: Grafo con vértice aislado
    print("Test 3: Vértice aislado")
    grafo3 = Grafo(es_dirigido=True)
    grafo3.agregar_vertice("Solo")
    grafo3.agregar_vertice("X")
    grafo3.agregar_vertice("Y")
    grafo3.agregar_arista("X", "Y")
    
    comun3 = comunidades(grafo3)
    
    # "Solo" debería estar en su propia comunidad
    comunidad_solo = None
    for c in comun3:
        if "Solo" in c:
            comunidad_solo = c
            break
    
    assert comunidad_solo is not None, "Vértice aislado no encontrado"
    assert len(comunidad_solo) == 1, f"Vértice aislado debería estar solo, comunidad: {comunidad_solo}"
    
    print(f"✅ Vértice aislado detectado:")
    for i, c in enumerate(comun3, 1):
        print(f"   Comunidad {i}: {sorted(c)}")
    print()
    
    
    # Test 4: Grafo completo (todos conectados)
    print("Test 4: Grafo completo (clique)")
    grafo4 = Grafo(es_dirigido=True)
    vertices_clique = ["P", "Q", "R", "S"]
    for v in vertices_clique:
        grafo4.agregar_vertice(v)
    
    # Conectar todos con todos
    for v1 in vertices_clique:
        for v2 in vertices_clique:
            if v1 != v2:
                grafo4.agregar_arista(v1, v2)
    
    comun4 = comunidades(grafo4)
    
    # Todos deberían estar en la misma comunidad
    assert len(comun4) == 1, f"Clique debería formar 1 comunidad, formó {len(comun4)}"
    assert set(comun4[0]) == set(vertices_clique), f"Todos los vértices deberían estar juntos"
    
    print(f"✅ Clique detectado como 1 comunidad: {sorted(comun4[0])}\n")
    
    
    # Test 5: Red social realista
    print("Test 5: Red social realista")
    grafo5 = Grafo(es_dirigido=True)
    
    # Grupo 1: Amigos de la facu
    grafo5.agregar_vertice("Alice")
    grafo5.agregar_vertice("Bob")
    grafo5.agregar_vertice("Charlie")
    grafo5.agregar_arista("Alice", "Bob")
    grafo5.agregar_arista("Bob", "Alice")
    grafo5.agregar_arista("Bob", "Charlie")
    grafo5.agregar_arista("Charlie", "Bob")
    grafo5.agregar_arista("Alice", "Charlie")
    
    # Grupo 2: Amigos del trabajo
    grafo5.agregar_vertice("David")
    grafo5.agregar_vertice("Eve")
    grafo5.agregar_vertice("Frank")
    grafo5.agregar_arista("David", "Eve")
    grafo5.agregar_arista("Eve", "David")
    grafo5.agregar_arista("Eve", "Frank")
    grafo5.agregar_arista("Frank", "Eve")
    
    # Conexión débil entre grupos (Alice conoce a David)
    grafo5.agregar_arista("Alice", "David")
    
    comun5 = comunidades(grafo5)
    
    print(f"   Red social formó {len(comun5)} comunidad(es):")
    for i, c in enumerate(comun5, 1):
        print(f"   Comunidad {i}: {sorted(c)}")
    
    # Verificar que hay al menos 1 comunidad
    assert len(comun5) >= 1, "Debe haber al menos 1 comunidad"
    
    # La conexión débil podría unir o no los grupos (no determinista)
    # Solo verificamos estructura válida
    total_vertices5 = sum(len(c) for c in comun5)
    assert total_vertices5 == 6, f"Debe haber 6 vértices en total, tiene {total_vertices5}"
    
    print("✅ Red social procesada correctamente\n")
    
    
    # Test 6: Grafo vacío
    print("Test 6: Grafo vacío")
    grafo6 = Grafo(es_dirigido=True)
    comun6 = comunidades(grafo6)
    assert comun6 == [], "Grafo vacío debería devolver lista vacía"
    print("✅ Grafo vacío manejado correctamente\n")
    
    
    print("=== ✅ TODOS LOS TESTS DE COMUNIDADES PASARON ===\n")


# Ejecutar todos los tests
if __name__ == "__main__":
    test_grafo()
    test_vertices_entrada()
    test_comunidades()
    
    print("\n" + "="*60)
    print("🎉 ¡TODOS LOS TESTS PASARON EXITOSAMENTE! 🎉")
    print("="*60)