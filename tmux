#!/bin/bash
set -e

session=etch

if ! tmux ls | grep -q "$session"; then
	tmux new-session -d -s $session

	tmux rename-window vim
	tmux split-window -v

	tmux send-keys -t $session:1.1 "vim" C-m

	tmux resize-pane -t $session:1.2 -D 20

	tmux send-keys -t $session:1.2 "modd" C-m
fi

tmux select-window -t $session:1
tmux select-pane -L
exec tmux attach-session -d -t $session
