# Solution for CU DU connect to free5GC
## Environment
- Golang 1.15.8
- kubernetes v1.23.7 (rke v1.3.12)
- free5GC v3.0.5

## Directory
### cmd
source code of cu, du setup exec

### deploy
yaml file for deploy on kubernetes

### configmap
configmap of cu, du conf

### sample-config
original conf file

### test
include cu, du, ue conf on container environment. CU, DU, UE on container and free5GC on k8s.

## Deploy
### Pre-requirements
Install golang; Setup kubernetes (with flannel cni, pod-network-cidr=10.244.0.0/16)

### Deploy multus cni
```
git clone https://github.com/k8snetworkplumbingwg/multus-cni.git
cd mulus-cni
cat ./deployments/multus-daemonset-thick-plugin.yml | kubectl apply -f -
```

* Create dual flannel cni
```
kubectl apply -f deploy/flannel-dual.yaml
```

### Deploy free5GC v3.0.5
* Setup NFS server as backend storage of MongoDB
```
sudo apt install -y nfs-kernel-server
sudo mkdir -p /var/share/free5gc
echo "/var/share/free5gc *(rw,sync,no_root_squash)" | sudo tee -a /etc/exports
sudo exportfs -r
## check if directory is exported correctly
sudo showmount -e
```

* Install gtp5g Kernel module
```
sudo apt install -y build-essential
# for free5gc v3.0.5
git clone -b v0.2.1 https://github.com/free5gc/gtp5g.git
cd gtp5g
make
sudo make install
# check if gtp5g is loaded
lsmod | grep gtp5g
```

* Deploy free5GC v3.0.5
For connection between CU and AMF, we need to insert secondary network interface into AMF, CU pods.
```
# Replace deploy/free5GC-kubernetes/free5gc-nf-yml/cni/amf.yaml with deploy/amf.yaml
mv deploy/free5GC-kubernetes/free5gc-nf-yml/cni/amf.yaml.old
cp deploy/amf.yaml deploy/free5GC-kubernetes/free5gc-nf-yml/cni

# Deploy free5GC v3.0.5
kubectl apply -f deploy/free5GC-kubernetes/free5gc-nf-yml/cni
```

* Add UE info through free5GC web console
Connect to <Node-IP>:30300
```
user: admin
passwd: free5gc
```
Click Registration UE with default UE setting.

### Deploy OAI-CU, DU
* Apply DU first
```
kubectl apply -f deploy/du-gnb.yaml
```
while pods du turn into running status, deploy cu immediately.
```
NAME                                 READY   STATUS    RESTARTS   AGE
oai-du-deployment-5d8f8669c4-wq9zw   1/1     Running   0          3s
```

* Apply CU
```
kubectl apply -f deploy/cu-gnb.yaml
```

### Run ue in container
```
docker run -ti --name ue --privileged alan0415/oai-ue:v0.1 bash
```

### Config UE
* Edit conf file
```
nano openairinterface5g/targets/PROJECTS/GENERIC-NR-5GC/CONF/ue.conf
```

```
# ue.conf
uicc0 = {
imsi = "208930000000003";
key = "8baf473f2f8fd09487cccbd7097c6862";
opc= "8e27b6af0e692e750f32667a3b14605d";
dnn= "internet";
nssai_sst=1;
nssai_sd=0x010203;
}
rfsimulator: {
    serverport: 30043;
}
```

* Install tools
```
apt update && apt install iputils-ping screen -y
```

* Start up UE
```
cd openairinterface5g/cmake_targets/ran_build/build
screen

# inside screen
RFSIMULATOR=<Node-IP> ./nr-uesoftmodem -r 106 --numerology 1 --band 78 -C 3619200000 --rfsim --sa --nokrnmod -O ../../../targets/PROJECTS/GENERIC-NR-5GC/CONF/ue.conf

# exit screen
Ctrl + A + D
```
Now you can ping to DN through UE

```
# show oaitun_ue1
ip addr

# ping to DN
ping -I oaitun_ue1 8.8.8.8
```

### Build code
* Generate exec file uesd to setup cu and du while running cu/du pods
```
make all
```
The Dockerfile will ADD the exec file
