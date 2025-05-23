watch: --start-air

restart: --build --start-bin

--start-air:
	@(air)

--build:
	@(go build -o tmp/main cmd/api/main.go)

--start-bin:
	@(tmp/main)