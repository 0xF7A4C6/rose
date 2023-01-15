"""
    - This script is used to test if killer successfully bind to the specified port
    - usage: python bind_test.py <port>
"""

import socket, sys

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print('Usage: %s <port>' % sys.argv[0])
        sys.exit(1)

    try:
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.bind(('0.0.0.0', 50023))
        s.listen(5)

        print('Result: ONLINE - Port %s is open.' % sys.argv[1])

        while True:
            c, addr = s.accept()
            c.close()
    except:
        print('Result: OFFLINE - Port %s cant be used.' % sys.argv[1])