class Rational(object):
    def __init__(self, numerator, denominator):
        g = self.gcd(numerator, denominator)
        if denominator < 0:
            self.numerator = -1 * numerator // g
            self.denominator = abs(denominator) // g
        else:
            self.numerator = numerator // g
            self.denominator = denominator // g

    def gcd(self, a, b):
        for i in range(min(abs(a), abs(b)), 1, -1):
            if a % i == 0 and b % i == 0:
                return i
        return 1

    def __eq__(self, other):
        if not isinstance(self, Rational):
            return False
        if not isinstance(other, Rational):
            return False
        if self.numerator == 0 and other.numerator == 0:
            return True
        return self.numerator == other.numerator and self.denominator == other.denominator

    def __repr__(self):
        return f'{self.numerator}/{self.denominator}'

    def __add__(self, other):
        res_numer = self.numerator * other.denominator + other.numerator * self.denominator
        res_denom = self.denominator * other.denominator
        return Rational(res_numer, res_denom)

    def __sub__(self, other):
        res_numer = self.numerator * other.denominator - other.numerator * self.denominator
        res_denom = self.denominator * other.denominator
        return Rational(res_numer, res_denom)

    def __mul__(self, other):
        res_numer = self.numerator * other.numerator
        res_denom = self.denominator * other.denominator
        return Rational(res_numer, res_denom)

    def __truediv__(self, other):
        res_numer = self.numerator * other.denominator
        res_denom = self.denominator * other.numerator
        return Rational(res_numer, res_denom)

    def __abs__(self):
        return Rational(abs(self.numerator), abs(self.denominator))

    def __pow__(self, power):
        return Rational(self.numerator ** power, self.denominator ** power)

    def __rpow__(self, base):
        return base ** (self.numerator / self.denominator)


if __name__ == '__main__':
    r = Rational(1, 2)
    p = Rational(3, 4)
    print(r / p)