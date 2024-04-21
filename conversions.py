import base64

def hexToB64(h: str) -> str:
    b_hex: bytes = bytes.fromhex(h)
    b_b64: bytes = base64.b64encode(b_hex)

    return b_b64.decode("ascii")

def b_to_ascii(b: bytes) -> str:
    return b.decode("ascii")