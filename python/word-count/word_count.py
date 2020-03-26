import re
from collections import defaultdict


def count_words(phrase):
    words_regex = re.compile(r"([a-z0-9][a-z0-9']*[a-z0-9]|[a-z0-9]+)")
    result = defaultdict(int)
    for word in words_regex.findall(phrase.lower()):
        result[word] += 1
    return dict(result)


if __name__ == "__main__":
    print(f"{count_words('one two three, four, one, one')}")
