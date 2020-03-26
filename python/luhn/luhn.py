class Luhn(object):
    def __init__(self, number):
        self.number = str(number).replace(' ', '')

    def calculate_luhn(self):
        num_list = [int(x) for x in self.number]
        not_doubled = [x for x in num_list[-1::-2]]
        doubled = [2 * x if 2 * x <= 9 else 2 * x - 9 for x in num_list[-2::-2]]
        return sum(doubled + not_doubled)

    def valid(self):
        if not self.number.isdigit():
            return False
        if len(self.number) <= 1:
            return False
        if self.calculate_luhn() % 10 != 0:
            return False
        return True


if __name__ == '__main__':
    print(f'{Luhn("123").valid()=}')
