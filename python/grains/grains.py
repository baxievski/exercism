def square(number):
    if number < 1:
        raise ValueError("Number of squares must be 1 or greater.")
    if number > 64:
        raise ValueError("Number of squares must be less than 64.")
    return 2 ** (number - 1)


def total():
    return 2 ** 64 - 1


if __name__ == '__main__':
    print(f"{square(3)=}")
    print(f"{square(4)=}")
    print(f"{square(5)=}")
    print(f"{square(6)=}")
