#### First let's create pv with a single node

Persistent Volume (PV): A storage resource in a Kubernetes cluster, provisioned by an administrator. PVs are cluster-wide resources that have specific storage capacities and access modes.

To get the details
```
kubectl get pv my-pv
```

Then the status should be Available
If not, check for errors using the coomand

```
kubectl describe pv my-pv
```


### pvc

Persistent Volume Claim (PVC): A request for storage by a user. PVCs consume PV resources, and their access modes and size requests must be met by available PVs.

First check the pvc details by running the command

```
kubectl describe pvc my-pvc
```

check for the status, it should be Bound

Use the below command to debug it.
```
kubectl describe pvc my-pvc
```
