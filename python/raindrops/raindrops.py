def convert(number):
    sound = ""
    if number % 3 == 0:
        sound += "Pling"
    if number % 5 == 0:
        sound += "Plang"
    if number % 7 == 0:
        sound += "Plong"
    if sound:
        return sound
    return str(number)


if __name__ == "__main__":
    print(f"{convert(1)=}")
    print(f"{convert(2)=}")
    print(f"{convert(3)=}")
    print(f"{convert(4)=}")
    print(f"{convert(5)=}")
    print(f"{convert(6)=}")
    print(f"{convert(15)=}")
    print(f"{convert(105)=}")
