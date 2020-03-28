class Rational(object):
    def __init__(self, numer, denom):
        g = self.gcd(numer, denom)
        if denom < 0:
            self.numer = -1 * numer // g
            self.denom = abs(denom) // g
        else:
            self.numer = numer // g
            self.denom = denom // g

    def gcd(self, a, b):
        for i in range(min(abs(a), abs(b)), 1, -1):
            if a % i == 0 and b % i == 0:
                return i
        return 1

    def __eq__(self, other):
        if not isinstance(self, Rational) or not isinstance(other, Rational):
            return False
        if self.numer == 0 and other.numer == 0:
            return True
        return self.numer == other.numer and self.denom == other.denom

    def __repr__(self):
        return f'{self.numer}/{self.denom}'

    def __add__(self, other):
        res_numer = self.numer * other.denom + other.numer * self.denom
        res_denom = self.denom * other.denom
        return Rational(res_numer, res_denom)

    def __sub__(self, other):
        res_numer = self.numer * other.denom - other.numer * self.denom
        res_denom = self.denom * other.denom
        return Rational(res_numer, res_denom)

    def __mul__(self, other):
        res_numer = self.numer * other.numer
        res_denom = self.denom * other.denom
        return Rational(res_numer, res_denom)

    def __truediv__(self, other):
        res_numer = self.numer * other.denom
        res_denom = self.denom * other.numer
        return Rational(res_numer, res_denom)

    def __abs__(self):
        return Rational(abs(self.numer), abs(self.denom))

    def __pow__(self, power):
        return Rational(self.numer ** power, self.denom ** power)

    def __rpow__(self, base):
        return base ** (self.numer / self.denom)


if __name__ == '__main__':
    r = Rational(1, 2)
    p = Rational(3, 4)
    print(r / p)