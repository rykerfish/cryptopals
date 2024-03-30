import base64
from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes
import utilities as ut

def main():

    key = b"YELLOW SUBMARINE"
    f = open("7.txt", "r")
    b64:str = f.read()

    byte_dat:bytes = base64.b64decode(b64)

    cipher = Cipher(algorithms.AES128(key), modes.ECB())
    decryptor = cipher.decryptor()
    text = decryptor.update(byte_dat) + decryptor.finalize()

    print(ut.b_to_ascii(text))


if __name__ == "__main__":
    main()