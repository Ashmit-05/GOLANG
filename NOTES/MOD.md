**`mod`** is one of the most useful commands in go. It performs a variety of operations. When downloading a module as a dependency, there might be a chance that it is corrupted. To avoid that, go creates a **go.sum** file, while stores the hashes of that version of the module. In case any discrepancy is found, it alerts you.
**Note that these operations are very expensive. Use them carefully in your CI/CD pipeline**

## Important Commands

1. `go get -u <github repo link>` --> used to install a go module as a dependency into your application
2. `go mod verify` --> Verifies the hashes of the modules
3. `go mod tidy` --> It adds any missing module requirements necessary to build the current module's packages and dependencies, if there are some not used dependencies go mod tidy will remove those from go. mod accordingly. It also adds any missing entries to go. sum and removes unnecessary entries.
4. `go mod init <something>` --> The go mod init command **creates a go.mod file to track your code's dependencies**. So far, the file includes only the name of your module and the Go version your code supports. But as you add dependencies, the go.mod file will list the versions your code depends on.
5. `go list` --> It will list the dependencies of your current go file
	1. `go list all` --> this will list all the dependencies installed in your system
	2. `go list -m all` --> all the dependencies used in your file
	3. `go list -m -versions <dependency name>` --> this will list out all the available versions of a dependency
6. `go mod why <dependency name>` --> this will list out all the modules which depend on the given module
7. `go mod graph` --> shows a graph of the modules and their dependencies
8. `go mod edit` --> used to edit the mod file
	1. `-go` --> to change the go version used by the module
	2. `-module` --> to change the module name
9. `go mod vendor` --> **Not very advisable**. This creates a vendor file, which is like node modules in node. If you wish to make any changes to the dependency code, you can do it by accessing it here & publish your own version. **Note that if you use to `go run main.go` command after making any changes to the vendor file, it will just run the code and dependencies from the web or the cache. In order to reflect the changes, use `go run -mod=vendor main.go`**