from enum import Enum
from threading import Lock


class BankAccount:
    State = Enum("State", ("new", "open", "closed"))

    def __init__(self):
        self._balance = 0
        self._state = BankAccount.State.new
        self._lock = Lock()

    def get_balance(self):
        with self._lock:
            if self._state == BankAccount.State.closed:
                raise ValueError(f"Can't get balance of a closed account.")
            return self._balance

    def open(self):
        with self._lock:
            if self._state == BankAccount.State.open:
                raise ValueError(f"Can't open already opened account.")
            self._state = BankAccount.State.open

    def deposit(self, amount):
        with self._lock:
            if amount < 0:
                raise ValueError(f"Can't deposit negative amount: {amount}.")
            if self._state == BankAccount.State.closed:
                raise ValueError(f"Can't deposit to a closed accocunt.")
            self._balance += amount

    def withdraw(self, amount):
        with self._lock:
            if amount < 0:
                raise ValueError(f"Can't withdraw negative amount: {amount}.")
            if self._state == BankAccount.State.closed:
                raise ValueError(f"Can't withdraw from a closed acocunt.")
            if amount > self._balance:
                raise ValueError(f"Can't withdraw more than current balance.")
            self._balance -= amount

    def close(self):
        with self._lock:
            if self._state == BankAccount.State.closed:
                raise ValueError(f"Can't close already closed account.")
            if self._state == BankAccount.State.new:
                raise ValueError(f"Can't close not opened account.")
            self._state = BankAccount.State.closed
            self._balance = 0


if __name__ == "__main__":
    pass
