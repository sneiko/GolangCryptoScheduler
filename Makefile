current_dir := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))
app_base = ${current_dir}"/cmd/app"

run:
	go run ${app_base}/main.go

clean:
	go clean