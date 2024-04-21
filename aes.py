from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes

def ecb_encrypt(dat, key):
    cipher = Cipher(algorithms.AES(key), modes.ECB())
    encryptor = cipher.encryptor()
    return encryptor.update(dat) + encryptor.finalize()

def ecb_decrypt(dat, key):
    cipher = Cipher(algorithms.AES(key), modes.ECB())
    decryptor = cipher.decryptor()
    return decryptor.update(dat) + decryptor.finalize()