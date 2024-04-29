def pkcs_pad(b1: bytearray, blocklength: int = 16) -> bytearray:
    n = len(b1)
    pad_len = blocklength % n
    pad = [pad_len]*pad_len 

    padded = bytearray(b1)
    padded.extend(pad)

    return padded

def strip_pkcs(b1: bytearray, blocklength: int = 16) -> bytearray:
    n = len(b1)
    if b1 % blocklength == 0:
        return b1
    
    pad_num:int = b1[-1]
    if pad_num > blocklength-1:
        raise Exception("Can't strip PKCS padding: last byte doesn't seem to be padding.")
    
    for i in range(0,pad_num):
        if b1[n-i-1] != pad_num:
            raise Exception("Can't strip PKCS padding: incorrect number of padded bytes")
        
    return b1[0:n-pad_num]

def xor(b1: bytes, b2: bytes) -> bytearray:
    if(len(b1) != len(b2)):
        raise Exception("Arrays aren't same length")

    return  bytearray(a ^ b for a,b in zip(b1, b2))

def rank(decode: str): # ranks based on the number of letters most common in english language
    common: str = "etaoin shrdlu"
    sum_common = 0

    decode = decode.lower()

    for c in common:
        sum_common = sum_common + decode.count(c)

    return sum_common

def getLetterFreq():
    freq_dict = {'a': 0.0651738, 'b': 0.0124248, 'c': 0.0217339, 'd': 0.0349835,
        'e': 0.1041442, 'f': 0.0197881, 'g': 0.0158610, 'h': 0.0492888,
        'i': 0.0558094, 'j': 0.0009033, 'k': 0.0050529, 'l': 0.0331490,
        'm': 0.0202124, 'n': 0.0564513, 'o': 0.0596302, 'p': 0.0137645,
        'q': 0.0008606, 'r': 0.0497563, 's': 0.0515760, 't': 0.0729357,
        'u': 0.0225134, 'v': 0.0082903, 'w': 0.0171272, 'x': 0.0013692,
        'y': 0.0145984, 'z': 0.0007836, ' ': 0.1918182 }
    
    return freq_dict