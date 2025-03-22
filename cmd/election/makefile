run:
	tmux split-window -h "sleep 2 && go run ./worker"
	tmux split-window -v "sleep 2 && go run ./worker"
	go run ./cache/cmd
