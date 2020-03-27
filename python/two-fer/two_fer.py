def two_fer(name=None):
    if name is None:
        name = "you"
    return f"One for {name}, one for me."


if __name__ == "__main__":
    print(f"{two_fer()=}")
    print(f'{two_fer("Nevenka")=}')
