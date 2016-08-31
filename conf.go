/*
Conf package allows you to write a config file in JSON format and call
the stored values.
The stored values can be Int, Float64, Bool, String.


Example usage:

	package main

	import "git.tele2.net/nms/conf"
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
*/
package conf

import "encoding/json"
import "os"
import "flag"
import "log"

type config struct {
	conf        map[string]interface{}
	initialized bool
}

var c config = config{
	initialized: false,
}

func init() {
	read()
}

func read() {
	fileName := flag.String("conf", "", "Configuration file name with absolute path.")
	flag.Parse()
	if *fileName == "" {
		log.Fatal("Use -conf command argument to include a configuration file.")
	}

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal("Error opening config file ", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c.conf)
	if err != nil {
		log.Fatal("Error parsing config file: ", err)
	}
	c.initialized = true
	return
}

// GetInt returns an integer value and boolean status.
// If the status is false, the value is unreliable.
func GetInt(param string) (int, bool) {
	v, ok := c.conf[param]
	if !ok || !c.initialized {
		return 0, false
	}
	switch v := v.(type) {
	case int:
		return v, true
	case float64:
		return int(v), true
	default:
		return 0, false
	}
}

// GetFloat returns a float64 value and boolean status.
// If the status is false, the value is unreliable.
func GetFloat(param string) (float64, bool) {
	v, ok := c.conf[param]
	if !ok || !c.initialized {
		return 0, false
	}
	switch v := v.(type) {
	case float64:
		return v, true
	default:
		return 0, false
	}
}

// GetString returns a string value and boolean status.
// If the status is false, the value is unreliable.
func GetString(param string) (string, bool) {
	v, ok := c.conf[param]
	if !ok || !c.initialized {
		return "", false
	}
	switch v := v.(type) {
	case string:
		return v, true
	default:
		return "", false
	}
}

// GetBool returns a boolean value and boolean status.
// If the status is false, the value is unreliable.
func GetBool(param string) (bool, bool) {
	v, ok := c.conf[param]
	if !ok || !c.initialized {
		return false, false
	}
	switch v := v.(type) {
	case bool:
		return v, true
	default:
		return false, false
	}
}
