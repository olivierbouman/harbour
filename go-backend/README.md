# back-end in Go

### LOCAL SETUP

# Add path to go storage location

run the following command from the location where you want to use Go: `export PATH=$PATH:/usr/local/go/bin`
Check with `echo $PATH` whether the path has been added

# Add go module file

go mod init go-backend

# Download all the used Go modules

go mod tidy

# Run go-backend

- Regular compiling: go run cmd/web/\*.go
- Live reloading: using the air package: `https://github.com/cosmtrek/air`
  -> Simply run: `air` (and check if `export PATH=$PATH:$(go env GOPATH)/bin`)
- Kill port in use: `https://stackoverflow.com/questions/11583562/how-to-kill-a-process-running-on-particular-port-in-linux`

### GO

# Pointers

For extra info
`https://www.youtube.com/watch?v=sTFJtxJXkaY&list=PL0q7mDmXPZm625OqhkdCPRkY9XAq6X7kW&index=3`

Variable:

- name: foo // fmt.Println(foo) will print out the value of foo
- type: int
- value: 23
- address: 0xc00 // "&foo" can be read as "address of foo"

If we assign `p := &foo` we can say p is a pointer to foo
Dereferencing: If we then say `*p` we get the value that the pointer 'p' is pointing towards; `*p` would then give `23` as the value of.
Assigning a new value to `*p` will change the value of `foo`
`func xxx (a *int){...}` Means that `a` will be a pointer of int type/struct
We can then subsequently call this function by `xxx(&foo)`, passing through a pointer to `foo`

# Receivers

# Interfaces
