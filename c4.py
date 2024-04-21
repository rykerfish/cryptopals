import utilities as ut
import conversions as conv

f = open("data/4.txt", "r")
lines = f.readlines()

best_score = 0
best_str = ""

for hex_str in lines:

    b_str = bytes.fromhex(hex_str)

    for key in range(0,128):
        decode = bytes((b & 0b01111111) ^ key for b in b_str) # discards upper bit that isn't ascii?
        d_str = conv.b_to_ascii(decode).lower()

        score = ut.rank(d_str)

        if score > best_score:
            best_score = score
            best_str = d_str

print(best_str)
print(best_score)