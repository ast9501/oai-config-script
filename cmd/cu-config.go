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
	amfDomainName := "free5gc-amf-svc"
	duDomainName := "oai-du-svc"
	localInterfaceName := "eth0"
	filePath := "openairinterface5g/targets/PROJECTS/GENERIC-NR-5GC/CONF/cu_gnb.conf"

	// config mnc
	mod := "20c\\                  mnc = 93;"
	cmd := exec.Command("sed", "-i", mod, filePath)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[Err]", err)
	} else {
		fmt.Println("[LOG]Set mnc: 93")
	}

	// config protc
	mod = "43c\\    local_s_portc   = 600;"
	cmd = exec.Command("sed", "-i", mod, filePath)
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[Err]", err)
	} else {
		fmt.Println("[LOG]Set local_s_protc: 600")
	}

	mod = "45c\\    remote_s_portc  = 601;"
	cmd = exec.Command("sed", "-i", mod, filePath)
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[Err]", err)
	} else {
		fmt.Println("[LOG]Set remote_s_protc: 601")
	}

	// config local ip address
	ip, err := GetInterfaceIpv4Addr(localInterfaceName)
	if err != nil {
		fmt.Println("[Err]", err)
	}
	//fmt.Println("Config local interface ip: ", ip)
	mod = "41c\\    local_s_address = \"" + ip + "\";"
	cmd = exec.Command("sed", "-i", mod, filePath)
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[Err]", err)
	} else {
		fmt.Println("[LOG]Set interface ip: ", ip)
	}

	// TODO: Add fixed virtual ip as ip
	ip = "10.244.1.2"
	mod = "204c\\        GNB_IPV4_ADDRESS_FOR_NG_AMF              = \"" + ip + "/16\";"
	cmd = exec.Command("sed", "-i", mod, filePath)
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[Err]", err)
	} else {
		fmt.Println("[LOG]Set virtual ip: ", ip)
	}

	mod = "206c\\        GNB_IPV4_ADDRESS_FOR_NGU                 = \"" + ip + "/16\";"
	cmd = exec.Command("sed", "-i", mod, filePath)
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[Err]", err)
	} else {
		fmt.Println("[LOG]Set virtual ip: ", ip)
	}

	// config local interface name
	mod = "40c\\    local_s_if_name = \"" + localInterfaceName + "\";"
	cmd = exec.Command("sed", "-i", mod, filePath)
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[Err]", err)
	} else {
		fmt.Println("[LOG]Set local interface name: ", localInterfaceName)
	}

	// config AMF ip address
	iprecords, _ := net.LookupIP(amfDomainName)
 	for _, amfIp := range iprecords {
		modAmfAddr := "193c\\        amf_ip_address      = ( { ipv4       = \"" + amfIp.String() + "\";"
		//fmt.Println(modAmfAddr)
  		cmd = exec.Command("sed", "-i", modAmfAddr, filePath)
		_, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("[Err]", err)
		} else {
			fmt.Println("[LOG]Config AMF ip: ", amfIp)
		}
		
 	}

	// config du ip address
	iprecords, _ = net.LookupIP(duDomainName)
 	for _, duIp := range iprecords {
		modDuAddr := "42c\\    remote_s_address = \"" + duIp.String() + "\";"
		//fmt.Println(modAmfAddr)
  		cmd = exec.Command("sed", "-i", modDuAddr, filePath)
		_, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("[Err]", err)
		} else {
			fmt.Println("[LOG]Config DU ip: ", duIp)
		}
 	}
}