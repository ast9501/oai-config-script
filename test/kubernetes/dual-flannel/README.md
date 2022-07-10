

# Test yaml file about oai-gnb, ue and free5GC
## Test environment
* K8s 1.19.0-00
* flannel cni
* free5gc v3.0.5

## Pre-requirements
* flannel cni with pod-network-cidr=10.244.0.0/16

## Install multus
```
git clone https://github.com/k8snetworkplumbingwg/multus-cni.git
cd multus-cni
cat ./deployments/multus-daemonset-thick-plugin.yml | kubectl apply -f -
```

* create network resource
```
kubectl apply -f flannel1.yaml
```

## Apply free5GC network
```
git clone https://github.com/ast9501/free5GC-stage3-on-k8s.git
```
* Replace free5gc-nf-yml/cni/amf.yaml with amf.yaml under this directory
