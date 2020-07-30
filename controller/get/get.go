package get

import "github.com/RomuloDurante/baseNewServer/modelcustomer"

//StartService ...
func StartService() {
	user := modelcustomer.CreateCustomer("Romulo")

	user.Greeting()
}
