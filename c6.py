import utilities as ut
import base64
import numpy as np
import scipy

def main():
    test_a: bytes = bytes("this is a test", encoding="ascii")
    test_b: bytes = bytes("wokka wokka!!!", encoding="ascii")

    h_dist: int = hammingDist(test_a, test_b)
    assert(h_dist == 37)

    f = open("6.txt", "r")
    b64:str = f.read()
    byte_dat:bytes = base64.b64decode(b64)

    # tests different keysizes and gets the one with the smallest normalized hamming distance
    leastAvgDist = avgHammingDist(byte_dat, 1)
    keysize = 1
    for k in range(2, 40):
        avgDist = avgHammingDist(byte_dat, k)

        if avgDist < leastAvgDist:
            leastAvgDist = avgDist
            keysize = k

    print(keysize)

    chunks = splitDat(byte_dat, keysize)

    key = []
    for c in range(0,keysize):

        dat = chunks[c]
        key.append(solveSingleKeyXOR(dat))

    decrypt = []

    key_ind = 0
    for i in range(0, len(byte_dat)):
        decrypt.append(byte_dat[i] ^ key[key_ind])
        key_ind = key_ind + 1
        if key_ind == len(key):
            key_ind = 0

    decrypt = bytes(decrypt)
    decrypt = ut.b_to_ascii(decrypt)

    print(decrypt)
    key = [chr(x) for x in key]
    key = ''.join(key)
    print(key)

    

def solveSingleKeyXOR(byte_arr:bytes) -> int:

    best_score = -1
    best_key = -1

    for key in range(0,128):
        decode = bytes(b ^ key for b in byte_arr)
        d_str = ut.b_to_ascii(decode).lower()

        score = scoreLetterFreq(d_str)

        if score > best_score:
            best_score = score
            best_key = key

    return best_key

# def scoreLetterFreq(decode:str) -> float:

#     exp_freq = ut.getLetterFreq()
#     obs_freq = {}
#     for char in exp_freq.keys():
#         obs_freq[char] = decode.count(char)/len(decode)

#     exp_vals = list(exp_freq.values())
#     obs_vals = list(obs_freq.values())

#     hellinger = (1/np.sqrt(2)) * np.linalg.norm(np.sqrt(exp_vals) - np.sqrt(obs_vals))

#     return 1-hellinger*hellinger

def scoreLetterFreq(decode:str) -> float:

    exp_freq = ut.getLetterFreq()
    obs_freq = {}
    for char in exp_freq.keys():
        obs_freq[char] = decode.count(char)/len(decode)

    exp_vals = list(exp_freq.values())
    obs_vals = list(obs_freq.values())

    bhatt_dist = 0
    for i in range(0,len(exp_vals)):
        bhatt_dist += np.sqrt(exp_vals[i]*obs_vals[i])

    return bhatt_dist

def splitDat(dat: bytes, keysize:int) -> list:

    chunks = [dat[x::keysize] for x in range(0,keysize)]

    return chunks

def avgHammingDist(byteArr:bytes, keysize:int, n_chunks:int=40) -> float:

    if n_chunks % 2 != 0:
        raise IndexError("Number of chunks to average over needs to be divisible by 2")

    avgDist = 0
    for i in range(0,n_chunks):
        a = byteArr[i*keysize:(i+1)*keysize]
        b = byteArr[(i+1)*keysize:(i+2)*keysize]

        avgDist += hammingDist(a,b)/keysize # normalize by keysize

    avgDist /= int(n_chunks) # number of avg change in bits per byte

    return avgDist



def hammingDist(a:bytes, b:bytes) -> int:

    xor_ab = ut.xor(a,b)

    ones_count = 0
    for i in range(0, len(xor_ab)):
        bin_n = bin(xor_ab[i])
        ones_count = ones_count + bin_n.count("1")

    return ones_count


if __name__ == "__main__":
    main()