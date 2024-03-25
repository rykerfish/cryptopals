plain = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"

key = bytes("ICE", encoding="ascii")
plain = bytes(plain, encoding="ascii")

encrypt = []

key_ind = 0
for i in range(0, len(plain)):
    encrypt.append(plain[i] ^ key[key_ind])
    key_ind = key_ind + 1
    if key_ind == len(key):
        key_ind = 0

encrypt = bytes(encrypt).hex()

ans_str = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

assert(encrypt == ans_str)