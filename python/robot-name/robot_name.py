import random
import string


class Robot(object):
    def __init__(self):
        _seed = random.Random()
        name = ''.join(_seed.choices(string.ascii_uppercase, k=2))
        name += ''.join(_seed.choices(string.digits, k=3))
        self.name = name

    def reset(self):
        self.__init__()


if __name__ == '__main__':
    robot = Robot()
    print(f"{robot.name=}")
