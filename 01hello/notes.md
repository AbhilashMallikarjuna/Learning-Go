Detailed description for a specific command
```go
go help <command_name>
```
### Lexer
* The job of the lexer is to understand that the user is following the 
grammer of the language.
* We don't need to manually put the semicolons in the code. The lexer puts them before
compiling

### Types
* Case sensitive
* Variable type should be known in advance
* Everything is a type
* The first intial captial letter will provide whether it is public or private
```go
	fmt.Println("Hi")
```
Here Println is exported publicly in whichever module it is present
#### Various types:
* String
* Bool
* integer : Signed and unsigned
  * uint8 
  * uint64 
  * int8 
  * int64 
  * uintptr
* Floating : 
  * float32 
  * float64
* Complex
* Arrays
* Slices
* Maps
* Structs
* Pointers