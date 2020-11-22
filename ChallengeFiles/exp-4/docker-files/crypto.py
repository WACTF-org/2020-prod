import sys
sys.path.append("/opt/")
import aes

key = b'\xca\x94\xbe\x15\xda\xfe\x57\x8f\x75\x3a\xf3\xfd\xab\x85\x84\x91'
iv = b'\xff' * 16

def decrypt(msg):
    return aes.AES(key).decrypt_ctr(msg, iv)
