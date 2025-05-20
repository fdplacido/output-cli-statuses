# Compile

```shell
# In tmux, set the environment so it picks the built command
set-environment -g PATH "~/apps/google-cloud-sdk/bin:~/go/bin/:/usr/local/bin:/usr/bin:/bin:/opt/homebrew/bin/"
```

```shell
go build -o ~/go/bin/tmux-status-info ./cmd/main.go
```