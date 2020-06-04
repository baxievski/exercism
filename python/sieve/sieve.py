def primes(limit):
    result = []
    primes = [True if x >= 2 else False for x in range(limit + 1)]

    for i, prime in enumerate(primes):
        if not prime:
            continue

        result.append(i)

        for j in range(i, limit + 1, i):
            primes[j] = False

    return result


if __name__ == "__main__":
    print(f"{primes(100)=}")
