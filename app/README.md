# GO gRPC DB Server

## Dependencies

### App: [Dep](https://github.com/golang/dep)


### Tools: [Retool](https://github.com/twitchtv/retool)

> retool helps manage the versions of tools that you use with your repository. These are executables that are a crucial part of your development environment, but aren't imported by any of your code, so they don't get scooped up by glide or godep (or any other vendoring tool).

```
usage: retool (add | remove | upgrade | sync | do | build | help)

use retool with a subcommand:

add will add a tool
remove will remove a tool
upgrade will upgrade a tool
sync will synchronize your _tools with tools.json, downloading if necessary
build will compile all the tools in _tools
do will run stuff using your installed tools

help [command] will describe a command in more detail
version will print the installed version of retool
```

Add a tool dependency:

    retool add github.com/jteeuwen/go-bindata/go-bindata origin/master
               ^ tool                                    ^ commit

Prepend `retool do` to use tools installed by retool:

    retool do go-bindata -pkg testdata -o ./testdata/testdata.go ./testdata/data.json

### Testing

    docker run --link a8f182c53bdc:localdev --net docker-go-postgres_default grpc_cli ls localdev:8080 --enable_ssl
