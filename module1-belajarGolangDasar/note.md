# Notes

Development: go run {filename}.go

Production / When Finished: go build → this will compile it to binary, you can run it by doing ./{moduleName}

Golang couldn't have multiple functions with the same name.

```
# module1-belajarGolangDasar
.\sample.go:5:6: main redeclared in this block
        .\helloWorld.go:5:6: other declaration of main
```

Learning solution:
Just run it with go run, but later on in production never do this, instead just make 1 function, and make the others unique.

# Go Data Types:
- Integer
  - int8
  - int16
  - int32
  - int64
  - uint8 (unsigned integer, min 0, no negatives)
  - uint16
  - uint32
  - uint64
- Floating Point
  - float32
  - float64
  - complex64 (mostly for imaginary numbers, it's barely used)
  - complex128
- Boolean (bool)
  - true
  - false
- String
  - "{x}"

# Go Aliases
- byte : uint8
- rune : int32
- int : Minimal int32
- uint : Minimal uint32

# String commands
- len("{x}") =  count the character count
- "{x}"[number] = fetch character at the number, refer to [string.go](string.go)

# Variables
unlike python or javascript or php, we can't store multiple data types in a variable in golang, we use var refer to [variable.go](variable.go)