## Config map and secrets
ConfigMaps are typically used for non-sensitive configuration data, while Secrets are used for storing sensitive information <br/>
ConfigMaps stores data as key-value pairs, whereas Secrets stores data as base64-encoded data<br/>

### Understand the configmap.yaml sample file

In the sample configmap.yaml, we have
ConfigMap, Secret and a Deployment

How to apply the yaml file
```
kubectl apply -f example.yaml
kubectl get configmap
kubectl get secret
kubectl get deployment
```

Access the logs of the deployed container to verify that the environment variables are correctly injected:
```
kubectl logs -l app=myapp
```

How to inspect the environment variables directly?
```
kubectl exec -it <pod-name> -- /bin/bash
echo $APP_NAME
echo $DB_USERNAME
echo $DB_PASSWORD
```
#### Significance of | (pipe)
In the below example ,
Pipe (|): Indicates that the value is a block scalar, preserving the newlines and formatting.
Multi-line Content: Each line after the pipe is part of the value for the key app.properties.
```
  app.properties: |
    app.name=Test App
    app.version=1.0
    app.env=production
```

#### config map from a file.

You can also use config map from a file
First create two files, named host.txt and port.txt
in a folder temp

Then run
```
kubectl create configmap cmap-from-file-host --from-file=temp/host.txt
kubectl create configmap cmap-from-file-port --from-file=temp/port.txt
```

Then you can check weather they are created using the command
```
kubectl get configmap cmap-from-file-host
kubectl get configmap cmap-from-file-port
```

Yaml file of them can be seen by

```
kubectl get configmap cmap-from-file-host -o yaml
kubectl get configmap cmap-from-file-port -o yaml
```

You will see yaml like
```
apiVersion: v1
data:
  host.txt: |
    http://localhost
kind: ConfigMap
metadata:
  creationTimestamp: "2024-06-23T06:49:30Z"
  name: cmap-from-file-host
  namespace: default
  resourceVersion: "352"
  uid: 12ae001f-eab1-46c3-b5ae-f90f2f9efd77
```

So the config map is created with the key host.txt,
same as the file name.

```
data:
  host.txt: |
```
See the file configmap-from-file.yaml, to see how the
config map is used in it.

#### Secret from cli

We can create two secret from cli as shown below

```
kubectl create secret generic db-credentials --from-literal=db-username=dbuser --from-literal=db-password=mypass
```

Then if you check the yaml file by running

```
kubectl get secret db-credentials -o yaml
```

You will get out put like

```
apiVersion: v1
data:
  db-password: bXlwYXNz
  db-username: ZGJ1c2Vy
kind: Secret
```

Then see secret-from-cli.yaml, how it's used in a pod.

Then run
```
kubectl exec -it my-pod-new -- /bin/bash
```
After that, from the container run
```
echo $DB_USERNAME
```
That will show  `dbuser`

#### Pass env values to a pod
See the yaml file
pass-env.yaml

We can easily pass them to a pod using

```
env:
    - name: Variable1
      value: "somevalue"
    - name: Variable2
      value: "someothervalue"
```
Then from the pod we can print it like echo $Variable1, echo $Variable2
