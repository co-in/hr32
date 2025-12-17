web-mem: profile
	go tool pprof -http=:8080 mem.prof

web-cpu: profile
	go tool pprof -http=:8081 cpu.prof

profile:
	go test -benchtime=60s -bench='^Benchmark_.*HR32$$' -cpuprofile=cpu.prof -memprofile=mem.prof

benchmark:
	go test -benchtime=5s -bench=.

test:
	go test .

web:
	@$(shell which xdg-open >/dev/null 2>&1 && echo xdg-open || echo open) index.html