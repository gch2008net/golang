module example.com/hello

go 1.22.4

// replace example.com/greetings => ./../greetings

// require example.com/greetings v0.0.0-00010101000000-000000000000

require github.com/gch2008net/common-go v1.0.5

require golang.org/x/example/hello v0.0.0-20240716161537-39e772fc2670 // indirect
