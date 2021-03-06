package mqtt

import (
	"os"
	"strconv"
	"time"
	"zabbix.com/pkg/conf"
	"zabbix.com/pkg/plugin"
)

type Options struct {
	//Interval int `conf:"optional,range=0:10,default=10"`
	Username string `conf:"optional"` // MQTT username to use
	Password string `conf:"optional"` // MQTT password to use
	ClientID string `conf:"optional"`
	Timeout  int    `conf:"optional,range=0:30"` //MQTT Connect timeout
}

func (p *Plugin) Configure(global *plugin.GlobalOptions, private interface{}) {

	//p.options.Interval = 10
	p.options.Username = ""
	p.options.Password = ""
	p.options.Timeout = global.Timeout

	hostname, _ := os.Hostname()
	p.options.ClientID = hostname + strconv.Itoa(time.Now().Second())

	if err := conf.Unmarshal(private, &p.options); err != nil {
		p.Warningf("cannot unmarshal configuration options: %s", err)
	}

	p.Debugf("configure: timeout=%d", p.options.Timeout)
	p.Debugf("configure: username=%s", p.options.Username)
	p.Debugf("configure: clientid=%s", p.options.ClientID)
	//p.Debugf("configure: password is set")

}

func (p *Plugin) Validate(private interface{}) (err error) {
	return
}
