import base64
from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes
import utilities as ut


def main():
    f = open("data/8.txt", "r")

    for line in f:
        dat = base64.b16decode(line.strip().upper()) # only works on A-F???
        detected = detectEcb(dat)
        if detected:
            print("ECB DETECTED:", line)

def detectEcb(dat):
    n = len(dat)
    if n % 16 != 0:
        print("err: length wrong for ecb, needs padding")
        exit(-1)

    blocks = n // 16
    block_arr = []

    for i in range(0, blocks):
        start = i*16
        block = dat[start:start+16]
        if block in block_arr:
            return True
        
        block_arr.append(block)

    return False

def aes_ecb(dat, key):
    cipher = Cipher(algorithms.AES128(key), modes.ECB())
    encryptor = cipher.decryptor()
    return encryptor.update(dat) + encryptor.finalize()


if __name__ == "__main__":
    main()