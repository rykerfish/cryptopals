from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes
import utilities as ut

def ecb_encrypt(dat, key):
    cipher = Cipher(algorithms.AES(key), modes.ECB())
    encryptor = cipher.encryptor()
    return encryptor.update(dat) + encryptor.finalize()

def ecb_decrypt(dat, key):
    cipher = Cipher(algorithms.AES(key), modes.ECB())
    decryptor = cipher.decryptor()
    return decryptor.update(dat) + decryptor.finalize()

def cbc_encrypt(pt: bytearray, key: bytes, iv: bytes, blocklength: int = 16) -> bytearray:

    pt = ut.pkcs_pad(pt)

    n = len(pt)
    n_blocks = int(n / blocklength)

    ct = bytearray(n)
    ct[0:blocklength] = ut.xor(pt[0:blocklength], iv)
    ct[0:blocklength] = ecb_encrypt(ct[0:blocklength], key)

    for i in range(1, n_blocks):
        prev = (i-1)*blocklength
        start = i*blocklength
        end = start + blocklength

        ct[start:end] = ut.xor(ct[prev:start], pt[start:end])
        ct[start:end] = ecb_encrypt(ct[start:end], key)

    return ct

def cbc_decrypt(ct: bytearray, key: bytes, iv: bytes, blocklength: int = 16) -> bytearray:
    n = len(ct)

    n_blocks = int(n / blocklength)

    pt = bytearray(n)
    pt[0:blocklength] = ecb_decrypt(ct[0:blocklength], key)
    pt[0:blocklength] = ut.xor(pt[0:blocklength], iv)


    for i in range(1,n_blocks):
        prev = (i-1)*blocklength
        start = i*blocklength
        end = start + blocklength

        pt[start:end] = ecb_decrypt(ct[start:end], key)
        pt[start:end] = ut.xor(ct[prev:start], pt[start:end])

    return pt