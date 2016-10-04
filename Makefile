bench:
	go test -v -bench=. .

profile:
	go test -v -bench=. . -cpuprofile cpu.out
	go tool pprof -text ./ui.test.exe cpu.out
