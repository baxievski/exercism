from datetime import timedelta


def add(birth_date, seconds_to_add=10**9):
    return birth_date + timedelta(seconds=seconds_to_add)