import base64
from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes
import utilities as ut
import conversions as conv
import aes

def main():

    key = b"YELLOW SUBMARINE"
    f = open("data/7.txt", "r")
    b64:str = f.read()

    byte_dat:bytes = base64.b64decode(b64)

    text = aes.ecb_decrypt(byte_dat, key)

    print(conv.b_to_ascii(text))


if __name__ == "__main__":
    main()