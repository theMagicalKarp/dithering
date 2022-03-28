BINARY_NAME=dither
BUILD_DIR=build

build:
	go build -o ${BUILD_DIR}/${BINARY_NAME} main.go

clean:
	go clean
	rm -rf ${BUILD_DIR}

test:
	go test ./...

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all

examples: build
	${BUILD_DIR}/${BINARY_NAME} --in examples/michelangelo/in.png --out examples/michelangelo/out.png --out-format png
	${BUILD_DIR}/${BINARY_NAME} --in examples/baby/in.png --out examples/baby/out.png --out-format png
	${BUILD_DIR}/${BINARY_NAME} --in examples/american_gothic/in.jpg --out examples/american_gothic/out.jpg --out-format jpeg
	${BUILD_DIR}/${BINARY_NAME} --in examples/milkyway/in.png --out examples/milkyway/out.png --out-format png
	${BUILD_DIR}/${BINARY_NAME} --in examples/gopher/in.png --out examples/gopher/out.png --out-format png
	${BUILD_DIR}/${BINARY_NAME} --in examples/factors/in.png --out examples/factors/out_1.png --out-format png --factor 1
	${BUILD_DIR}/${BINARY_NAME} --in examples/factors/in.png --out examples/factors/out_2.png --out-format png --factor 2
	${BUILD_DIR}/${BINARY_NAME} --in examples/factors/in.png --out examples/factors/out_3.png --out-format png --factor 3
	${BUILD_DIR}/${BINARY_NAME} --in examples/factors/in.png --out examples/factors/out_4.png --out-format png --factor 4

.PHONY: build run clean test test_coverage dep vet lint
