import utilities as ut

str = "YELLOW SUBMARINE"

a = ut.pkcs_pad(bytearray(str, encoding="ascii"), blocklength=20)

ans = bytes("YELLOW SUBMARINE\x04\x04\x04\x04", encoding="ascii")

assert(a == ans)
print(a)
