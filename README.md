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

# Cheat sheet

```
devswiss base64-encode [string] - Encodes a given base64 string.
devswiss base64-decode [string] - Decodes a given base64 string.

devswiss uuid [-version 7] [-count 10] - Generates one or more UUIDs.
```
