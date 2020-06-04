def steps(number, number_of_steps=0):
    if number <= 0:
        raise ValueError(f"Defined only for positive numbers, not '{number}'.")

    if number == 1:
        return number_of_steps

    if number % 2 == 0:
        return steps(number / 2, number_of_steps + 1)

    return steps(3 * number + 1, number_of_steps + 1)


if __name__ == "__main__":
    print(f"{steps(120)=}")
    print(f"{steps(1000000)=}")
