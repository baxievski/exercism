class TriangleError(Exception):
    pass


class Triangle(object):
    def __init__(self, side_a, side_b, side_c):
        [self.a, self.b, self.c] = sorted([side_a, side_b, side_c])

    @property
    def valid(self):
        if self.a <= 0 or self.b <= 0 or self.c <= 0:
            return False
        if self.a + self.b <= self.c:
            return False
        return True

    @property
    def equilateral(self):
        if not self.valid:
            return False
        return self.a == self.b == self.c

    @property
    def isosceles(self):
        if not self.valid:
            return False
        return self.a == self.b or self.b == self.c

    @property
    def scalene(self):
        if not self.valid:
            return False
        return [self.a, self.b, self.c] == list(set([self.a, self.b, self.c]))


def equilateral(sides):
    triangle = Triangle(*sides)
    return triangle.equilateral


def isosceles(sides):
    triangle = Triangle(*sides)
    return triangle.isosceles


def scalene(sides):
    triangle = Triangle(*sides)
    return triangle.scalene


if __name__ == "__main__":
    pass
