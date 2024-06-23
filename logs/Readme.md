# Common log check commands

### Get the status of major components
```
 kubectl get componentstatuses
```
Which lists the component status as shown below.

NAME                 STATUS    MESSAGE   ERROR
controller-manager   Healthy   ok        
scheduler            Healthy   ok        
etcd-0               Healthy   ok  


### Get logs from the master and node components

##### master components

* Kube APIServe
```
sudo journalctl -u kube-apiserver
```
* Kube Scheduler
```
sudo journalctl -u kube-scheduler
```
* Kube Controller Manager
```
sudo journalctl -u kube-controller-manager
```


##### worker node components
CNI (Container Network Interface) logs
```
sudo journalctl -u calico-node
```
Kube-Proxy logs
```
sudo journalctl -u kube-proxy
```
Kubelet log
```
sudo journalctl -u kubelet
```
Container Runtime log
```
sudo journalctl -u containerd
```
