## Build script

```go

go build passgen.go

# GOOS='windows' go build passgen.go

```

- Instalation on Linux

```bash
git clone https://github.com/jesherdevsk8/gerador-senha.git /tmp/gerador-senha \
&& cd /tmp/gerador-senha \
&& go build passgen.go \
&& sudo mv passgen /usr/local/bin/gerarsenha 

```
