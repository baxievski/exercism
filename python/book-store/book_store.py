from itertools import combinations


def combination_price(combination):
    if len(combination) in [0, 1]:
        return len(combination) * 800
    if len(combination) == 2:
        return len(combination) * 800 * (1 - 0.05)
    if len(combination) == 3:
        return len(combination) * 800 * (1 - 0.10)
    if len(combination) == 4:
        return len(combination) * 800 * (1 - 0.20)
    if len(combination) == 5:
        return len(combination) * 800 * (1 - 0.25)
    return None


def total(books):
    if len(set(books)) < 2:
        return len(books) * 800
    possible_prices = []
    for n in [2, 3, 4, 5]:
        for book_combination in combinations(set(books), n):
            books_not_in_combination = books[:]
            for book in book_combination:
                books_not_in_combination.remove(book)
            price = combination_price(book_combination) + total(books_not_in_combination)
            possible_prices.append(price)
    return min(possible_prices)


if __name__ == '__main__':
    print(f"{total([1])=}")
    print(f"{total([1, 2])=}")
    print(f"{total([1, 2, 3])=}")
    print(f"{total([1, 2, 3, 4])=}")
    print(f"{total([1, 2, 3, 4, 5])=}")
    print(f"{total([1, 1, 3, 3, 1])=}")
    print(f"{total([1, 2, 3, 4, 5])=}")
