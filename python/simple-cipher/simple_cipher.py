import re
import string
import random


class Cipher(object):
    def __init__(self, key=None):
        if key is None:
            key = "".join(random.choices(string.ascii_lowercase, k=128))

        if not (key.islower() and key.isalpha()):
            raise ValueError("All items in the key must be chars and lowercase!")

        self.key = key

    def encode(self, text):
        plain_text = re.sub(r"[^a-z]", "", text.lower())
        encoded_text = ""
        for i, char in enumerate(plain_text):
            shift_distance = ord(self.key[i % len(self.key)]) - ord("a")
            encoded_character = chr(
                ord("a")
                + (ord(char) - ord("a") + shift_distance) % (ord("z") - ord("a") + 1)
            )
            encoded_text += encoded_character

        return encoded_text

    def decode(self, text):
        encoded_text = re.sub(r"[^a-z]", "", text.lower())
        decoded_text = ""
        for i, char in enumerate(encoded_text):
            shift_distance = ord(self.key[i % len(self.key)]) - ord("a")
            decoded_character = chr(
                ord("a")
                + (ord(char) - ord("a") - shift_distance) % (ord("z") - ord("a") + 1)
            )
            decoded_text += decoded_character

        return decoded_text


class Caesar(Cipher):
    def __init__(self):
        self.key = "d"


if __name__ == "__main__":
    cipher = Cipher()
    print(f'{cipher.encode("aabbcc")=}')
