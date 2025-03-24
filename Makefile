run:
	go run .

split:
	tmux split-window -h "go run ./cmd/test"
	tmux split-window -v "go run ./cmd/test"
	go run ./cmd/test
