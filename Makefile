build:
	go build -o build/ main.go
brun: 
	./build/main
run:
	go run main.go $@

clean: 
	rm -rf build

add:
	@cobra-cli add $@