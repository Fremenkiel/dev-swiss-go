# Install

To build and make usable, do as follow:
Ensure to have .local/bin added to your PATH.

```zsh
go build main.go
mv main ~/.local/devswiss
```

# Completion

Add the following to .zshrc

```
eval "$(devswiss completion zsh)"
```
