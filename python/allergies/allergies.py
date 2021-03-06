class Allergies(object):
    names = {
        1: "eggs",
        2: "peanuts",
        4: "shellfish",
        8: "strawberries",
        16: "tomatoes",
        32: "chocolate",
        64: "pollen",
        128: "cats",
    }

    def __init__(self, score):
        _score = score % 256
        self.lst = [
            Allergies.names[i]
            for i in (2 ** x for x in range(8))
            if i & _score
        ]

    def allergic_to(self, item):
        if item in self.lst:
            return True
        return False
