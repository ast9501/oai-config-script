package main

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
)

func GetInterfaceIpv4Addr(interfaceName string) (addr string, err error) {
	var (
		ief			*net.Interface
		addrs 		[]net.Addr
		ipv4Addr 	net.IP
	)

	if ief, err = net.InterfaceByName(interfaceName); err != nil {
		return
	}
	if addrs, err = ief.Addrs(); err != nil {
		return
	}
	for _, addr := range addrs {
		if ipv4Addr = addr.(*net.IPNet).IP.To4(); ipv4Addr != nil {
			break
		}
	}
	if ipv4Addr == nil {
		return "", errors.New(fmt.Sprintf("Interface %s dont't have ipv4 address\n", interfaceName))
	}
	return ipv4Addr.String(), nil
}

func main() {
	//amfDomainName := "sample.com"
	cuDomainName := "sample.com"
	localInterfaceName := "ens3"

	// config local ip address
	ip, err := GetInterfaceIpv4Addr(localInterfaceName)
	if err != nil {
		fmt.Println("[Err]", err)
	}
	fmt.Println("Config local interface ip: ", ip)
	mod := "193c\\    local_n_address = \"" + ip + "\";"
	cmd := exec.Command("sed", "-i", mod, "du_gnb.conf")
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[Err]", err)
	} else {
		fmt.Println("[LOG]Set local_n_address: ", ip)
	}

	// config local interface name
	mod = "192c\\    local_n_if_name = \"" + localInterfaceName + "\";"
	cmd = exec.Command("sed", "-i", mod, "du_gnb.conf")
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[Err]", err)
	} else {
		fmt.Println("[LOG]Set local_n_if_name: ", localInterfaceName)
	}

	// config cu ip address
	iprecords, _ := net.LookupIP(cuDomainName)
 	for _, cuIp := range iprecords {
		modCuAddr := "194c\\    remote_n_address = \"" + cuIp.String() + "\";"
		//fmt.Println(modAmfAddr)
  		cmd = exec.Command("sed", "-i", modCuAddr, "du_gnb.conf")
		_, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("[Err]", err)
		} else {
			fmt.Println("Config CU ip: ", cuIp)
		}
 	}
}