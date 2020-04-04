def factors(number):
    factors = []
    counter = 2
    while counter * counter <= number:
        if number % counter:
            counter += 1
            continue
        number = number // counter
        factors.append(counter)
    if number > 1:
        factors.append(number)
    return factors


if __name__ == "__main__":
    print(f"{factors(93819012551)=}")
