

class UnionFind:
    def __init__(self, elems: list):
        self.grupos = dict()
        for e in elems:
            self.grupos[e] = e        
    
    def find(self, v):
        """
        UnionFInd.find(v) devuelve el grupo donde pertenece V
        """
        if self.grupos[v] == v:
            return v
        # Actualizar el valor y devolver valor del grupo
        self.grupos[v] = self.find(self.grupos[v])
        return self.grupos[v]
    
    def union(self, u, v):
        """
        UnionFind.union(u, v) une a U al grupo de V
        """
        nuevo_grupo = self.find(u)
        otro = self.find(v)
        self.grupos[otro] = nuevo_grupo
    