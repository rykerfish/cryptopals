import base64

def hexToNB64(h: str) -> str:
    b_hex: bytes = bytes.fromhex(h)
    b_b64: bytes = base64.b64encode(b_hex)

    return b_b64.decode("ascii")
    
def xor(b1: bytes, b2: bytes) -> bytes:
    if(len(b1) != len(b2)):
        raise Exception("Arrays aren't same length")

    return  bytes(a ^ b for a,b in zip(b1, b2))