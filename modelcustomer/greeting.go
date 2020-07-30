package modelcustomer

import "fmt"

//Greeting ...
func (c *Customer) Greeting() {
	fmt.Println(c.Name)
}
