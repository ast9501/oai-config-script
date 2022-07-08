# CU
```
 "Mounts": [
            {
                "Type": "bind",
                "Source": "/dev/cpu_dma_latency",
                "Destination": "/dev/cpu_dma_latency",
                "Mode": "",
                "RW": true,
                "Propagation": "rprivate"
            },
            {
                "Type": "bind",
                "Source": "/sys/devices/system/cpu",
                "Destination": "/sys/devices/system/cpu",
                "Mode": "",
                "RW": true,
                "Propagation": "rprivate"
            }
        ],
"ExposedPorts": {
                "2152/tcp": {},
                "600/tcp": {},
                "601/tcp": {}
            },
```

# DU
```
 "Mounts": [
            {
                "Type": "bind",
                "Source": "/dev/cpu_dma_latency",
                "Destination": "/dev/cpu_dma_latency",
                "Mode": "",
                "RW": true,
                "Propagation": "rprivate"
            }
        ],
"ExposedPorts": {
                "2152/tcp": {},
                "4043/tcp": {},
                "600/tcp": {},
                "601/tcp": {}
            },
```

# UE
```
 "Mounts": [
            {
                "Type": "bind",
                "Source": "/dev/cpu_dma_latency",
                "Destination": "/dev/cpu_dma_latency",
                "Mode": "",
                "RW": true,
                "Propagation": "rprivate"
            }
        ],
"ExposedPorts": {
                "2152/tcp": {},
                "600/tcp": {},
                "601/tcp": {}
            },
```
