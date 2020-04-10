import re


class PhoneNumber(object):
    def __init__(self, phone_number):
        number = re.sub(r"[\(\)\.\-\+\s]", "", phone_number)
        self._number = number
        if not self._validate():
            raise ValueError(f"'{phone_number}' - not a valid north american phone number.")

    def _validate(self):
        if len(self._number) > 11:
            return False
        if len(self._number) == 11 and not self._number[0] == "1":
            return False
        if not self._number.isdigit():
            return False
        if self.area_code[0] in "01":
            return False
        if self.exchange_code[0] in "01":
            return False
        return True

    @property
    def area_code(self):
        n = self._number
        if len(n) == 11:
            n = n[1:]
        return n[:3]

    @property
    def exchange_code(self):
        n = self._number
        if len(n) == 11:
            n = n[1:]
        return n[3:6]

    @property
    def subscriber_number(self):
        n = self._number
        if len(n) == 11:
            n = n[1:]
        return n[6:]

    @property
    def number(self):
        return self.area_code + self.exchange_code + self.subscriber_number

    def pretty(self):
        return f"({self.area_code}) {self.exchange_code}-{self.subscriber_number}"


if __name__ == "__main__":
    pass
