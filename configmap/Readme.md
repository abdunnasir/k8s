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
