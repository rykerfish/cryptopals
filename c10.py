import utilities as ut
import aes
import base64

def cbc_encrypt(pt: bytearray, key: bytes, iv: bytes, blocklength: int = 16) -> bytearray:

    pt = ut.pkcs_pad(pt)

    n = len(pt)
    n_blocks = int(n / blocklength)

    ct = bytearray(n)
    ct[0:blocklength] = ut.xor(pt[0:blocklength], iv)
    ct[0:blocklength] = aes.ecb_encrypt(ct[0:blocklength], key)

    for i in range(1, n_blocks):
        prev = (i-1)*blocklength
        start = i*blocklength
        end = start + blocklength

        ct[start:end] = ut.xor(ct[prev:start], pt[start:end])
        ct[start:end] = aes.ecb_encrypt(ct[start:end], key)

    return ct


def cbc_decrypt(ct: bytearray, key: bytes, iv: bytes, blocklength: int = 16) -> bytearray:
    n = len(ct)

    n_blocks = int(n / blocklength)

    pt = bytearray(n)
    pt[0:blocklength] = aes.ecb_decrypt(ct[0:blocklength], key)
    pt[0:blocklength] = ut.xor(pt[0:blocklength], iv)


    for i in range(1,n_blocks):
        prev = (i-1)*blocklength
        start = i*blocklength
        end = start + blocklength

        pt[start:end] = aes.ecb_decrypt(ct[start:end], key)
        pt[start:end] = ut.xor(ct[prev:start], pt[start:end])

    return pt




test:bytearray = bytearray("testing testing!123 longer pls", encoding="ascii")
key:bytes = bytes("poggers doggers!", encoding="ascii")
iv = bytes(16)
ct:bytearray = cbc_encrypt(test, key, iv)
pt:bytearray = cbc_decrypt(ct, key, iv)

pt = ut.strip_pkcs(pt)
assert(pt == test)

f = open("data/10.txt", "r")
b64:str = f.read()
dat:bytearray = bytearray(base64.b64decode(b64))

key = bytes("YELLOW SUBMARINE", encoding="ascii")
iv = bytes(16) # all zeros

pt = cbc_decrypt(dat, key, iv)
# print(pt)

f.close()