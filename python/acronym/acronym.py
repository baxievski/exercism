import re


def abbreviate(words):
    words_regex = re.compile(r"([a-z0-9][a-z0-9']*[a-z0-9]|[a-z0-9]+)")
    result = (word[0].upper() for word in words_regex.findall(words.lower()))
    return "".join(result)


if __name__ == "__main__":
    print(f"{abbreviate('Central intelligence agency')=}")
    print(f"{abbreviate('Crazy labs')=}")
