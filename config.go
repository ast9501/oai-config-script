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
	amfDomainName := "sample.com"
	duDomainName := "sample.com"
	localInterfaceName := "ens3"

	// config local ip address
	ip, err := GetInterfaceIpv4Addr(localInterfaceName)
	if err != nil {
		fmt.Println("[Err]", err)
	}
	fmt.Println("Config local interface ip: ", ip)
	mod := "41c\\    local_s_address = \"" + ip + "\";"
	cmd := exec.Command("sed", "-i", mod, "cu_gnb.conf")
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[Err]", err)
	}

	mod = "204c\\        GNB_IPV4_ADDRESS_FOR_NG_AMF              = \"" + ip + "\";"
	cmd = exec.Command("sed", "-i", mod, "cu_gnb.conf")
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[Err]", err)
	}

	mod = "206c\\        GNB_IPV4_ADDRESS_FOR_NGU                 = \"" + ip + "\";"
	cmd = exec.Command("sed", "-i", mod, "cu_gnb.conf")
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[Err]", err)
	}

	// config local interface name
	mod = "40c\\    local_s_if_name = \"" + localInterfaceName + "\";"
	cmd = exec.Command("sed", "-i", mod, "cu_gnb.conf")
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[Err]", err)
	}

	// config AMF ip address
	iprecords, _ := net.LookupIP(amfDomainName)
 	for _, amfIp := range iprecords {
		modAmfAddr := "193c\\        amf_ip_address      = ( { ipv4       = \"" + amfIp.String() + "\";"
		//fmt.Println(modAmfAddr)
  		cmd = exec.Command("sed", "-i", modAmfAddr, "cu_gnb.conf")
		_, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("[Err]", err)
		}
		fmt.Println("Config AMF ip: ", amfIp)
 	}

	// config du ip address
	iprecords, _ = net.LookupIP(duDomainName)
 	for _, duIp := range iprecords {
		modDuAddr := "42c\\    remote_s_address = \"" + duIp.String() + "\";"
		//fmt.Println(modAmfAddr)
  		cmd = exec.Command("sed", "-i", modDuAddr, "cu_gnb.conf")
		_, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("[Err]", err)
		}
		fmt.Println("Config DU ip: ", duIp)
 	}
}