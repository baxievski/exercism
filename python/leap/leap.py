def leap_year(year):
    if year % 4 != 0:
        return False

    if year % 100 == 0:
        if year % 400 == 0:
            return True
        return False

    return True

if __name__ == "__main__":
    print(f"1997: {leap_year(1997)}")
    print(f"1996: {leap_year(1996)}")
    print(f"1900: {leap_year(1900)}")
    print(f"2000: {leap_year(2000)}")