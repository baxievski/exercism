def square_of_sum(n):
    return sum(x for x in range(1, n+1)) ** 2


def sum_of_squares(n):
    return sum(x**2 for x in range(1, n+1))


def difference_of_squares(n):
    return square_of_sum(n) - sum_of_squares(n)


if __name__ == "__main__":
    print(f"{sum_of_squares(1)=}")
    print(f"{sum_of_squares(2)=}")
    print(f"{square_of_sum(1)=}")
    print(f"{square_of_sum(2)=}")
