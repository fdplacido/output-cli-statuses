# TMUX Status Info

## provide mappings

- Update the `gcloud.json` and `kubectl.json` files with own values.

- For gcloud check from:
```shell
gcloud projects list
```

- For kubectl check from:
```shell
kubectl config get-contexts
```

## tmux conf

```shell
# set the environment so it picks the commands
set-environment -g PATH "~/apps/google-cloud-sdk/bin:~/go/bin/:/usr/local/bin:/usr/bin:/bin:/opt/homebrew/bin/"

# show it somewhere
set -g status-right "#(tmux-status-info)"

# Update the status bar every n seconds
set -g status-interval 5
```

# Compile to your local machine

```shell
go build -o $GOPATH/bin/tmux-status-info ./cmd/main.go
```