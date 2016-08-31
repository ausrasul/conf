# Configuration parser package for Golang
Conf package allows you to write a config file in JSON format and call
the stored values.
The stored values can be Int, Float64, Bool, String.

### Installation
```
go get github.com/ausrasul/conf
```

### Usage example:

```
	package main

	import "github.com/ausrasul/conf"
	import "fmt"

	func main() {
		a, ok := conf.GetInt("testInt")
		if !ok {
			fmt.Println("int not ok")
		}
		b, ok := conf.GetFloat("testFloat")
		if !ok {
			fmt.Println("float not ok")
		}
		c, ok := conf.GetString("testString")
		if !ok {
			fmt.Println("string not ok")
		}
		d, ok := conf.GetBool("testBool")
		if !ok {
			fmt.Println("bool not ok")
		}
		fmt.Println(a, b, c, d)
	}
```
