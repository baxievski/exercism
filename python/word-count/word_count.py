import re
from collections import defaultdict


def count_words(phrase):
    words = re.sub(r"[^a-z0-9']+", " ", phrase.lower()).split()
    words = [word.strip("'") for word in words]
    result = defaultdict(int)
    for word in words:
        result[word] += 1
    return dict(result)


if __name__ == "__main__":
    phrase = "''hey''"
    print(f"{phrase=}\n{count_words(phrase)=}")