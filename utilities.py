import base64

def hexToNB64(h: str) -> str:
    b_hex: bytes = bytes.fromhex(h)
    b_b64: bytes = base64.b64encode(b_hex)

    return b_b64.decode("ascii")

def b_to_ascii(b: bytes) -> str:
    return b.decode("ascii")
    
def xor(b1: bytes, b2: bytes) -> bytes:
    if(len(b1) != len(b2)):
        raise Exception("Arrays aren't same length")

    return  bytes(a ^ b for a,b in zip(b1, b2))

def getLetterFreq():
    freq_dict = {'e': 0.111607, 'm': 0.030129, 'a': 0.084966, 'h': 0.030034, 'r': 0.075809, 
             'g': 0.024705, 'i': 0.075448, 'b': 0.02072, 'o': 0.071635, 'f': 0.018121, 
             't': 0.069509, 'y': 0.017779, 'n': 0.066544, 'w': 0.012899, 's': 0.057351, 'k': 0.011016, 
             'l': 0.054893, 'v': 0.010074, 'c': 0.045388, 'x': 0.002902, 'u': 0.036308, 
             'z': 0.002721, 'd': 0.033844, 'j': 0.001965, 'p': 0.031671, 'q': 0.001962}
    
    return freq_dict