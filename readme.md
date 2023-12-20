### Go Modules

```bash
go mod init <project>
go get github.com/joho/godotenv
go mod tidy
```

`<project>` can be a string or a reference to a (GitHub) repo.

Ref:

- https://www.digitalocean.com/community/tutorials/how-to-use-go-modules
- https://www.alexedwards.net/blog/an-introduction-to-packages-imports-and-modules

### Support for multi-module workspaces

```bash
go work init
go work use project-one
go work use project-two
```

ref: https://stackoverflow.com/a/72131169

### ToDos

- add pre-commit checks for Go
