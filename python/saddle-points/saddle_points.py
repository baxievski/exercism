def saddle_points(matrix):
    if any(len(matrix[0]) != len(matrix[i]) for i, _ in enumerate(matrix)):
        raise ValueError('The matrix has to be regular')
    saddles = []
    for i, row in enumerate(matrix):
        for j, element in enumerate(row):
            if element != max(row):
                continue
            if element == min([row[j] for row in matrix]):
                saddles.append({"row": i + 1, "column": j + 1})
    return saddles


if __name__ == "__main__":
    matrix = [[4, 5, 4], [3, 5, 5], [1, 5, 4]]
    print(f"{matrix=}")
    print(f"{saddle_points(matrix)=}")
