# go.mod file parser
Easy and lightweight tool to list dependencies of project using project's go.mod

# Usage
```
cd project/dir/with/go.mod/in/it
go run github.com/BiF-kun/go-mod-parser
```

## Flags
```
-indirect
    show indirect dependencies
-version
    show dependencies version
```

## Quirks
You can use `GOFLAGS="-mod=mod"` to ignore vendor folder
