def all_pallindromes(beginning, end):
    result = [
        (i * j, (i, j))
        for i in range(beginning, end + 1)
        for j in range(i, end + 1)
        if str(i * j) == str(i * j)[::-1]
    ]
    return result


def largest(max_factor, min_factor=1):
    return max(all_pallindromes(min_factor, max_factor))


def smallest(max_factor, min_factor=1):
    return min(all_pallindromes(min_factor, max_factor))


if __name__ == "__main__":
    pass