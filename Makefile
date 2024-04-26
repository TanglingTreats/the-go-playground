build:
	@go build -o playground

crypto: build
	@./playground test.txt

run: build
	@./playground
