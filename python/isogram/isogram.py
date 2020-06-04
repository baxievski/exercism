import re
from collections import defaultdict


def is_isogram(string_to_check):
    character_counts = defaultdict(int)

    string_to_check = re.sub(r"\s+|\-", "", string_to_check, flags=re.UNICODE)
    string_to_check = string_to_check.lower()
    for character in string_to_check:
        character_counts[character] += 1
        if character_counts[character] <= 1:
            continue

        return False

    return True


if __name__ == "__main__":
    print(f'{is_isogram("bojan")=}')
    print(f'{is_isogram("axievski")=}')
