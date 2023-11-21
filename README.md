# go.mod file parser
Easy and lightweight tool to list dependencies of project using project's go.mod

# Usage
```
cd project/dir/with/go.mod/in/it
go run github.com/BiF-kun/go-mod-parser@latest
```

## Flags
```
-indirect
    show indirect dependencies
-version
    show dependencies version
```

## Quirks
Not using `@version` implies that you need to add this module to `go.mod` (and vendor folder if you are using vendoring scheme) and this might not be the desired approach (it will be removed on the next `go mod tidy` run)  
Anyway, you can bypass this by using `GOFLAGS="-mod=mod"`, which automatically implies `go get`, which downloads all packages and slows down execution in clean (e.g. docker) environments  
So yeah, it is clearly superior to use `@version` approach
