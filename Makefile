bench:
	go test -v -bench=. .

profile:
	go test -v -bench=. . -cpuprofile cpu.out
	go tool pprof -text ./ui.test.exe cpu.out

cover:
	go test -coverprofile=cover.out .
	#go tool cover -func=cover.out
	#rm cover.out

test:
	go test ./...

race:
	go test -race ./...
