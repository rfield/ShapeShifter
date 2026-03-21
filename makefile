
run: build
	./bin/shapeshifter
	
build: 
	go build -o bin/shapeshifter main.go

.PHONY: clean
clean:
	rm -f bin/shapeshifter
