from grafo import Grafo

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
    
    print("=== ✅ TODOS LOS TESTS PASARON EXITOSAMENTE ===")

# Ejecutar tests
if __name__ == "__main__":
    test_grafo()