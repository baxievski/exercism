def sum_of_multiples(limit, multiples):
    result = []
    for m in multiples:
        [
            result.append(n)
            for n in range(limit)
            if m != 0 and n % m == 0
        ]
    return sum(set(result))


if __name__ == "__main__":
    print(f"{sum_of_multiples(10000, [43, 47])=}")
