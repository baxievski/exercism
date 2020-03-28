def latest(scores):
    return scores[-1]


def personal_best(scores):
    return max(scores)


def personal_top_three(scores):
    return sorted(scores, reverse=True)[:3]


if __name__ == "__main__":
    print(f"{latest([10, 30, 90, 30, 100, 20, 10, 0, 30, 40, 40, 70, 70])=}")
    print(f"{personal_best([10, 30, 90, 30, 100, 20, 10, 0, 30, 40, 40, 70, 70])=}")
    print(f"{personal_top_three([10, 30, 90, 30, 100, 20, 10, 0, 30, 40, 40, 70, 70])=}")
