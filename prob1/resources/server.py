import http.server
import socketserver

PORT = 9090
uref = http.server.SimpleHTTPRequestHandler

with socketserver.TCPServer(("", PORT), uref) as httpd_luis:
    print("listening at port %d " % (PORT))
    httpd_luis.serve_forever()

