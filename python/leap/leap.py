def leap_year(year):
    if year % 4 != 0:
        return False
    if year % 100 != 0:
        return True
    if year % 400 == 0:
        return True
    return False

if __name__ == "__main__":
    print(f"{leap_year(1997)=}")
    print(f"{leap_year(1996)=}")
    print(f"{leap_year(1900)=}")
    print(f"{leap_year(2000)=}")