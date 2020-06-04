import string


def is_pangram(sentence):
    alphabet = set(string.ascii_lowercase)
    return alphabet.issubset(set(sentence.lower()))


if __name__ == "__main__":
    print(f"{is_pangram('abcdef')=}")
    print(f"{is_pangram('abcdefghijklmnopqrstuvwxyz')=}")
