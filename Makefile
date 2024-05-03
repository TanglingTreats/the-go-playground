build:
	@go build -o playground

crypto: build
	@./playground --mode encrypt test.txt

run: build
	@./playground
