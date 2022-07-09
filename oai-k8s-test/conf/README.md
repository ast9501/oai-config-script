## ip addr on amf
```
root@free5gc-amf-deployment-67776595d-b5mh6:/free5gc# ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
4: eth0@if528: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UP group default
    link/ether ce:13:5e:6c:5e:b0 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 10.244.0.33/24 brd 10.244.0.255 scope global eth0
       valid_lft forever preferred_lft forever
6: net1@if5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether c6:80:28:cd:8f:a5 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 10.244.0.201/24 brd 10.244.0.255 scope global net1
       valid_lft forever preferred_lft forever
```

## ip addr on cu
```
root@oai-cu-deployment-9968c4c48-vz8th:/oai/openairinterface5g/cmake_targets/ran_build/build# ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
4: eth0@if529: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UP group default
    link/ether be:4a:30:29:fd:69 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 10.244.0.34/24 brd 10.244.0.255 scope global eth0
       valid_lft forever preferred_lft forever
6: net1@if530: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 8a:29:b0:b4:43:68 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 10.244.0.202/24 brd 10.244.0.255 scope global net1
       valid_lft forever preferred_lft forever
```
