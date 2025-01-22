import random

import aes
import utilities as ut

def main():
    # do later
    pass



def encryption_oracle(text: bytearray) -> bytearray:
    
    prepend_count = random.randint(5,10)
    append_count = random.randint(5,10)

    prepend_text = bytearray([random.randint(0,255) for i in range(prepend_count)])
    append_text = bytearray([random.randint(0,255) for i in range(append_count)])

    text = prepend_text + text + append_text

    mode = random.randint(0,1)
    key = bytes([random.randint(0,255) for i in range(16)])

    text = ut.pkcs_pad(text)

    if mode == 0:
        ct = aes.ecb_encrypt(text, key)
    else:
        iv = bytes([random.randint(0,255) for i in range(16)])
        ct = aes.cbc_encrypt(text, key, iv)

    return ct

if __name__ == "__main__":
    main()