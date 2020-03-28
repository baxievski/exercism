import math


def classify(number):
    if number <= 0:
        raise ValueError('Only positive numbers can be classified')
    factors = sum([[i, number // i] for i in range(1, math.ceil(number ** 0.5)) if number % i == 0], [])
    if sum(factors) == 2 * number:
        result = 'perfect'
    elif sum(factors) > 2 * number:
        result = 'abundant'
    elif sum(factors) < 2 * number:
        result = 'deficient'
    return result


if __name__ == '__main__':
    pass