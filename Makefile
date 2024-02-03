OUT_FILE_NAME = "forum-Z"

run:
	@go run .

kill_and_run:
	@killport 8899

watch:
	@watchexec --restart --ignore docs --exts go make run

release:
	@go build -o ${OUT_FILE_NAME}

start:
	@yarn --cwd web-ui dev

