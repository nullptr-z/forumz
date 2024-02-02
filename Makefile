run:
	@go run .

kill_and_run:
	@killport 8899

watch:
	@watchexec --restart --ignore docs --exts go make run
