package modelcustomer

//Customer ...
type Customer struct {
	Name string
}

//CreateCustomer ...
func CreateCustomer(name string) *Customer {
	return &Customer{
		Name: name,
	}
}
