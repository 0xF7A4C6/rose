debug:
	clear
	go run compile.go debug
	sudo ./build/rose_debug debug
	rm -rf build/*

test:
	clear
	clear && gcc -std=c99 -static -pthread -D DEBUG ../src/lib/attack/*.c ../src/lib/network/*.c ../src/lib/security/*.c ../src/lib/update/*.c ../src/lib/utils/*.c ../src/*.c -o bin
	./bin debug
	rm -rf bin

release:
	clear
	go run compile.go release

cmd:
	clear
	go run compile.go cmd