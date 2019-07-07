package main

import "fmt"

type Container {
	id string
}
// This `person` struct type has `name` and `age` fields.
type NetworkSettings struct {
	ports map[string]map[string]string
}

// map[80/tcp:[map[HostIp:0.0.0.0 HostPort:8000]]]
func main() {
	var portsinstance = make(map[string]map[string]string)
	portsinstance["80/tcp"] = make(map[string]string)
	portsinstance["80/tcp"]["HostIp"] = "0.0.0.0"
	portsinstance["80/tcp"]["HostPort"] = "8000"
	networksettings := NetworkSettings{ports: portsinstance}
	fmt.Println(networksettings)

}

//  '{{range $p, $conf := .NetworkSettings.Ports}} {{$p}} -> {{(index $conf 0).HostPort}} {{end}}' $INSTANCE_ID
