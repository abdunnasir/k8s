# k8s
k8s tips


Get logs on a container within a pod
kubectl logs <pod_name> -c <container_name>
eg:  kubectl logs  nginx-77b4fdf86c-d68pz  -c  nginx


How can you get last 10 lines of logs from container nginx running inside pod nginx-7db9fccd9b-kvxrs?
kubectl logs nginx-7db9fccd9b-kvxrs -c nginx --namespace prod --tail=10


Which command can be used to list logs of all container with pods having label run=nginx in namespace prod?
kubectl logs -lrun=nginx --all-containers=true --namespace prod

Which command will list events in namespace prod and sort them by timestamp?
kubectl get events --sort-by='.metadata.creationTimestamp' --namespace prod


Which command(s) should be used to capture all the cluster information for debugging in a file name /var/logs/test-cluster.log?
kubectl cluster-info dump --output-directory=test.log



### Kind

How to create a cluster named kind-2?
```
kind create cluster --name kind-2
```

How to delete a cluster named kind-2
```
kind delete clusters kind-2
```
