package environment

import (
	"os"
	"os/user"
	"runtime"
)

type SystemInfo map[string]string

func NewSystemInfo() SystemInfo {
	c := SystemInfo{}
	c = make(map[string]string)
	host, err := os.Hostname()
	if err == nil {
		c.put("Machine Name", host)
	}
	user, err := user.Current()
	if err == nil {
		c.put("User Name", user.Username)
	}
	c.put("Operating System Name", runtime.GOOS)
	//c.put("Operating System Version", os.);
	c.put("Operating System Architecture", runtime.GOARCH)
	//c.put("Golang Name", (<any>process).release.name);
	c.put("Golang Version", runtime.Version())
	return c
}

func (c SystemInfo) put(parameter string, value string) {
	c[parameter] = value
}
