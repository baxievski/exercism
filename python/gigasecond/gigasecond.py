from datetime import timedelta, datetime


def add(birth_date, seconds_to_add=10**9):
    return birth_date + timedelta(seconds=seconds_to_add)


if __name__ == "__main__":
    print(f"{add(datetime(1977, 11, 8))=}")
