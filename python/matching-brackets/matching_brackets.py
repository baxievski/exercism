import re


def is_paired(input_string):
    cleaned_string = re.sub(r"[^\[\]\(\)\{\}]", "", input_string)
    stack = []
    opening_braces = "({["
    closing_braces = ")}]"
    for current_brace in cleaned_string:
        if current_brace in opening_braces:
            stack.append(current_brace)
            continue
        if len(stack) == 0:
            return False
        last_brace = stack.pop()
        opening_brace = opening_braces[closing_braces.index(current_brace)]
        if last_brace != opening_brace:
            return False
    return len(stack) == 0


if __name__ == "__main__":
    print(f"{is_paired('([])[]')=}")
