def is_valid(isbn):
    isbn_no_dash = isbn.replace("-", "")

    if len(isbn_no_dash) != 10:
        return False

    if isbn_no_dash[-1] not in "0123456789X":
        return False

    if not isbn_no_dash[:-1].isdigit():
        return False

    result = 0
    for position, character in enumerate(isbn_no_dash):
        if position == 9 and character == "X":
            character = "10"
        result += int(character) * (10 - int(position))

    return result % 11 == 0


if __name__ == "__main__":
    pass
