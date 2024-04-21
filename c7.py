import base64
from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes
import utilities as ut
import conversions as conv

def main():

    key = b"YELLOW SUBMARINE"
    f = open("data/7.txt", "r")
    b64:str = f.read()

    byte_dat:bytes = base64.b64decode(b64)

    text = aes_ecb(byte_dat, key)

    print(conv.b_to_ascii(text))


def aes_ecb(dat, key):
    cipher = Cipher(algorithms.AES(key), modes.ECB())
    encryptor = cipher.decryptor()
    return encryptor.update(dat) + encryptor.finalize()


if __name__ == "__main__":
    main()