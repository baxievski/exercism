def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError("strand_a, strand_b have different lengths")
    result = 0
    for i, c_str_a in enumerate(strand_a):
        if c_str_a != strand_b[i]:
            result += 1
    return result


if __name__ == "__main__":
    print(f"{distance('AABB', 'AABC')=}")
    print(f"{distance('AACB', 'AABC')=}")
    print(f"{distance('AACB', 'AAC')=}")
