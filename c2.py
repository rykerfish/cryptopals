import utilities as ut

h1 = "1c0111001f010100061a024b53535009181c"
h2 = "686974207468652062756c6c277320657965"

b1: bytes = bytes.fromhex(h1)
b2: bytes = bytes.fromhex(h2)

b3: bytes = ut.xor(b1, b2)

h3 = b3.hex()
ans_str = "746865206b696420646f6e277420706c6179"

assert(h3 == ans_str)