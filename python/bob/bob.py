def response(phrase):
    phrase = phrase.strip()
    if not phrase:
        return "Fine. Be that way!"

    if phrase.isupper() and phrase.endswith("?"):
        return "Calm down, I know what I'm doing!"

    if phrase.isupper():
        return "Whoa, chill out!"

    if phrase.endswith("?"):
        return "Sure."

    return "Whatever."


if __name__ == "__main__":
    print(f"{response()=}")
