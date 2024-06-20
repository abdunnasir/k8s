# Cluster IP Service in k8s

A ClusterIP service is a type of Kubernetes service that exposes a set of pods running on the cluster as a single, stable IP address. This IP address is reachable only from within the Kubernetes cluster.It is primarily used for internal communication between different parts of an application or between different applications running within the same Kubernetes cluster.

## How to access a clusetr IP service

From another pod in the same cluster, something like this.
```
wget -qO- http://clusterip-service:8080
```
No public access possible.

### How to make sure the test pod was able to connect ?
```
kubectl logs test-pod
```
It will show nginx default page content

# LoadBalancer Service type

It will provide EXTERNAL-IP when checking
```
kubectl get svc nginx-service
```
And then
curl http://<EXTERNAL-IP>:8080
Should show the nginx page content.

To get an External IP, You should use this
with a cloud service provider like
GKE (Google Kubernetes Engine), AKS (Azure Kubernetes Service), EKS (Amazon Elastic Kubernetes Service),

# Node Port service Type

NodePort services are useful for scenarios where you need to access your service externally but don't want to use an external load balancer.
```
kubectl get node -o wide
```

NAME                 STATUS   ROLES           AGE     VERSION   INTERNAL-IP   EXTERNAL-IP   OS-IMAGE                         KERNEL-VERSION       CONTAINER-RUNTIME
kind-control-plane   Ready    control-plane   3h28m   v1.30.0   172.24.0.2    <none>        Debian GNU/Linux 12 (bookworm)   4.15.0-213-generic   containerd://1.7.15

From here, you can see INTERNAL-IP = 172.24.0.2
```
kubectl get svc
```
nginx-service       NodePort    10.96.166.131   <none>        30020:30020/TCP   3h23m


So we can take http://172.24.0.2:30020/ on a browser

# Ingress Resource
