from threading import Lock


class BankAccount:
    def __init__(self):
        self.balance = 0
        self.is_open = None
        self.is_closed = None
        self.__lock = Lock()

    def get_balance(self):
        with self.__lock:
            if self.is_closed:
                raise ValueError(f"Can't get balance of a closed account.")
            return self.balance

    def open(self):
        with self.__lock:
            if self.is_open:
                raise ValueError(f"Can't open already opened account.")
            self.is_open = True
            self.is_closed = False

    def deposit(self, amount):
        with self.__lock:
            if amount < 0:
                raise ValueError(f"Can't deposit negative amount: {amount}.")
            if self.is_closed:
                raise ValueError(f"Can't deposit to a closed acocunt.")
            self.balance += amount

    def withdraw(self, amount):
        with self.__lock:
            if amount < 0:
                raise ValueError(f"Can't withdraw negative amount: {amount}.")
            if self.is_closed:
                raise ValueError(f"Can't withdraw from a closed acocunt.")
            if amount > self.get_balance():
                raise ValueError(f"Can't withdraw more than current balance.")
            self.balance -= amount

    def close(self):
        with self.__lock:
            if self.is_closed:
                raise ValueError(f"Can't close already closed account.")
            if not self.is_open:
                raise ValueError(f"Can't close not opened account.")
            self.is_closed = True
            self.is_open = False
            self.balance = 0


if __name__ == "__main__":
    pass