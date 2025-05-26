watch: --start-air

build: --build

--start-air:
	@(air)

--build:
	@(go build -o tmp/main cmd/api/main.go)