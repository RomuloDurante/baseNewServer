package config

//package with the config parans and API keys///
const (
	//PortToDeploy ...
	domain         = "www.bunkerlibrary.com"
	domainTest     = "localhost:5500"
	digitalOceanIP = "159.89.236.62"
	PORThttps      = ":443"
	PORThttp       = ":80"
	testIP         = "localhost" //"192.168.15.97"
	PORTtest       = ":5500"
	//->PRODUCTION
	//->PortToDeploy ...
	//PortToDeploy = digitalOceanIP + PORThttps
	//->Addr ...
	//Addr = "https://" + domain + "/"

	//->for test
	//->PortToDeploy ...
	PortToDeploy = testIP + PORTtest
	//->Addr ...
	Addr = "http://" + domainTest + "/"
)
