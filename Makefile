build:
	@go build -o playground

crypto: build
	@./playground --mode "encrypt" test.txt

interface: build
	@./playground --mode "interface"

context: build
	@./playground --mode "context"

run: build
	@./playground
