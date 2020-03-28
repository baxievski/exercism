import math


def classify(number):
    if number <= 0:
        raise ValueError("Only positive numbers can be classified")
    factors = sum(
        [
            [i, number // i]
            for i in range(1, math.ceil(number ** 0.5))
            if number % i == 0
        ],
        [],
    )
    if sum(factors) == 2 * number:
        return "perfect"
    if sum(factors) > 2 * number:
        return "abundant"
    return "deficient"


if __name__ == "__main__":
    print(f"{classify(6)=}")
    print(f"{classify(8)=}")
    print(f"{classify(9)=}")
    print(f"{classify(10)=}")
    print(f"{classify(11)=}")
    print(f"{classify(12)=}")
