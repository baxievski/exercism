import math


class ComplexNumber:
    def __init__(self, real, imaginary):
        self.real = real
        self.imaginary = imaginary
    
    def __repr__(self):
        return f"{self.real} + i * {self.imaginary}"

    def __eq__(self, other):
        if self.real != other.real:
            return False
        if self.imaginary != other.imaginary:
            return False
        return True

    def __add__(self, other):
        real = self.real + other.real
        imaginary = self.imaginary + other.imaginary
        return ComplexNumber(real, imaginary)

    def __mul__(self, other):
        real = self.real * other.real - self.imaginary * other.imaginary
        imaginary = self.real * other.imaginary + self.imaginary * other.real
        return ComplexNumber(real, imaginary)

    def __sub__(self, other):
        real = self.real - other.real
        imaginary = self.imaginary - other.imaginary
        return ComplexNumber(real, imaginary)

    def __truediv__(self, other):
        return self * other._reciprocal()

    def __abs__(self):
        return (self.real ** 2 + self.imaginary ** 2) ** 0.5

    def conjugate(self):
        return ComplexNumber(self.real, -1*self.imaginary)

    def exp(self):
        real = math.pow(math.e, self.real) * math.cos(self.imaginary)
        imaginary = math.pow(math.e, self.real) * math.sin(self.imaginary)
        return ComplexNumber(real, imaginary)

    def _reciprocal(self):
        real = self.real / (abs(self) ** 2)
        imaginary = -1 * self.imaginary / (abs(self) ** 2)
        return ComplexNumber(real, imaginary)

if __name__ == "__main__":
    print(f"{ComplexNumber(1, 2)=}")