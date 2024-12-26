# go-by-example
repo to learn Go using site [GoByExample](https://gobyexample.com/)

## CMD Line Notes
- add go.mod to each directory containing the go code by running:  
`go mod init [module/path]`  
- Update go.mod requirements with:  
`go mod tidy`
- In order to use your own local module, you have to redirect  Go tools from the incorrent module path to the local one:  
`go mod edit -replace go-by-example/greetings=../greetings`

## Code Notes
- `:-` is used to declare and assign a variable in one line
- Function names get capital letters, param types and return types  
`func Example(param string) string`
- A slice is the same as an array, who's size can change dynamically when adding or removing items  
`example_list_of_strings := []string`
- rand.Intn is used to return an int between 0:n  
`rand.Intn(len(example_list_of_strings))`
- for taking input from console:  
`	reader := bufio.NewReader(os.Stdin)`  
`	fmt.Print("Enter text: ")`  
`	text, _ := reader.ReadString('\n')`  
`	fmt.Println(text)`
- trim any whitespace  
`line = strings.TrimSpace(line)`
- convert string to integer ("string converstion . ASCAII to Int [A to i]):  
`strconv.Atoi.`

## Popular Import List
- "fmt"
- "error"
- log
- "math/rand" (random number generator module)
- "bufio" (implements buffered I/O, used to take input from console)
- "os" (used for file maniulation)
- "strconv" (for string conversions)
- "strings" (for string manipulation)