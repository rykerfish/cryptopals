import utilities as ut
import conversions as conv

def rank(decode: str): # ranks based on the number of letters most common in english language
    common: str = "etaoin shrdlu"
    sum_common = 0

    decode = decode.lower()

    for c in common:
        sum_common = sum_common + decode.count(c)

    return sum_common

freq = ut.getLetterFreq()

hex_str = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
b_str = bytes.fromhex(hex_str)

decode_freq = {}

best_score = 0
best_str = ""

for key in range(0,128):

    decode = bytes(b ^ key for b in b_str)
    d_str = conv.b_to_ascii(decode).lower()

    score = rank(d_str)

    if score > best_score:
        best_score = score
        best_str = d_str

print(best_str)
print(best_score)