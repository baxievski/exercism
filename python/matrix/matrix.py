class Matrix(object):
    def __init__(self, matrix):
        self.matrix = [[int(c) for c in r.split()] for r in matrix.split('\n')]

    def row(self, i):
        return self.matrix[i-1]
    
    def column(self, i):
        return [r[i-1] for r in self.matrix]


if __name__ == '__main__':
    test_matrix = Matrix("9 8 7\n5 3 2\n 6 6 7")
    print(test_matrix.column(1))