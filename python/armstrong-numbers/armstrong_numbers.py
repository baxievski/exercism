import math


def is_armstrong_number(number):
    if number == 0:
        return True

    result = 0
    digits = math.ceil(math.log10(abs(number)))
    for digit in range(digits + 1):
        result += ((number // (10 ** digit)) % 10) ** digits

    return result == number


if __name__ == "__main__":
    print(f"{is_armstrong_number(153)=}")
