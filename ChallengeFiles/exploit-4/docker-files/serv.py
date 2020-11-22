import socketserver
import http.server
import socket
import threading
import json
from crypto import decrypt
import re

HTTP_PORT = 8000 # 80 is exposed
SOCK_PORT = 8001 # 81 is exposed

rules = []
allowed_functions = ["add_rule", "remove_rule"]

def run_webserver():
	Handler = http.server.SimpleHTTPRequestHandler

	with socketserver.TCPServer(("", HTTP_PORT), Handler) as httpd:
		print("serving at port", HTTP_PORT)
		httpd.serve_forever()

def magic_packet_server():
	server_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
	server_socket.bind(('0.0.0.0', SOCK_PORT))
	print("serving at port", SOCK_PORT)

	while True:
		try:
			message, address = server_socket.recvfrom(1024)
			x = threading.Thread(target=handle_msg, args=(unpack_msg(message),))
			x.start()
			x.join(timeout=5)
		except:
			pass

def handle_msg(msg_json):
	func = msg_json["func"]
	ip = msg_json["ip"]
	port = msg_json["port"]

	if re.match(r".\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}.", ip):
		if re.match(r"\d{1,5}", str(port)):
			if func in allowed_functions:
				eval(f"{func}({ip},{port})")
		
def unpack_msg(msg):
	message = decrypt(msg)
	print(message)
	return json.loads(str(message, "utf-8"))

def add_rule(ip, port):
	rules.append(f"Allow:{ip},{port}")

def remove_rule(ip, port):
	rules.remove(f"Allow:{ip},{port}")

def main():
	while True:
		try:
			x = threading.Thread(target=run_webserver)
			x.setDaemon(True)
			x.start()
			magic_packet_server()
			x.join()
		except:
			pass

if __name__ == "__main__":
	main()

