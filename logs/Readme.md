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


#### Get the logs of a deployment

For a deployment that has labels as shown below
```
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
```

Get pods of it by running
```
kubectl get pods -l app=nginx
```

That will list three pods with names like
nginx-deployment-576c6b7b6-52j2c

#### Get events of a specific pod

```
kubectl describe pod nginx-deployment-576c6b7b6-52j2c
```


### Get events from the ReplicaSet

First list the replica set by running

```
kubectl get replicasets -l app=nginx
```
NAME                         DESIRED   CURRENT   READY   AGE
nginx-deployment-576c6b7b6   3         3         3       5m24s

Then get the events of it
```
kubectl describe replicaset nginx-deployment-576c6b7b6
```

### Gather Events from the Service

We have a service file nginx-service.yaml,
create a service using that
```
kubectl apply -f nginx-service.yaml
```

Now get the service details by running
```
kubectl get services
```
nginx-service   ClusterIP   10.96.177.26   <none>        80/TCP    11s

Now get the details of the service by running
```
kubectl describe service nginx-service
```

Get the end points of the service by running
```
kubectl get endpoints nginx-service
```

###### How are the Endpoints Determined?

Endpoints are determined by the labels defined in the Service's selector. Kubernetes automatically selects the pods that match the selector labels and includes their IP addresses in the endpoint list for the service.

### Display All the Logs from the Pods Created from the Deployment

It will list the logs
```
kubectl logs -l app=nginx --all-containers=true
```

It will list the logs and watch for the logs
```
kubectl logs -f -l app=nginx --all-containers=true
```
