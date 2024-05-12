import utilities as ut
import aes
import base64


test:bytearray = bytearray("testing testing!123 longer pls", encoding="ascii")
key:bytes = bytes("poggers doggers!", encoding="ascii")
iv = bytes(16)
ct:bytearray = aes.cbc_encrypt(test, key, iv)
pt:bytearray = aes.cbc_decrypt(ct, key, iv)

pt = ut.strip_pkcs(pt)
assert(pt == test)

f = open("data/10.txt", "r")
b64:str = f.read()
dat:bytearray = bytearray(base64.b64decode(b64))

key = bytes("YELLOW SUBMARINE", encoding="ascii")
iv = bytes(16) # all zeros

pt = aes.cbc_decrypt(dat, key, iv)
pt = ut.strip_pkcs(pt)
print(pt)

f.close()