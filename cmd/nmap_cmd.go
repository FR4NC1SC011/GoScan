package cmd

import (
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func Command(open_ports []int) []string {
	cmd := []string{"nmap"}
	ports_value := []string{}
	ports := "-p "

	for p := range open_ports {
		number := open_ports[p]
		text := strconv.Itoa(number)
		ports_value = append(ports_value, text)
	}

	ports = strings.Join(ports_value, ",")
	cmd = append(cmd, ports)

	color.Cyan("Nmap Command -> %#v", cmd)
	return cmd
}
