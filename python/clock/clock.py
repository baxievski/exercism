class Clock(object):
    def __init__(self, hour, minute):
        self.hour = (hour + minute // 60) % 24
        self.minute = minute % 60

    def __repr__(self):
        return f"{self.hour:02d}:{self.minute:02d}"

    def __add__(self, other):
        return Clock(self.hour, self.minute + other)

    def __sub__(self, other):
        return Clock(self.hour, self.minute - other)

    def __eq__(self, other):
        return self.hour == other.hour and self.minute == other.minute


if __name__ == "__main__":
    print(f"{Clock(20, 20)}")
