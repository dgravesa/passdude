#
all:
	mkdir -p bin
	go build -o bin ./...

clean:
	rm -r bin

