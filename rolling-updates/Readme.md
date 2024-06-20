# Explanation of the yaml file and commands

##### How to get public IP to show nginx default page
**Service expose the deploymnet to the an IP.**

​```
kubectl get nodes -o wide
​```

The above command gave INTERNAL-IP as  172.24.0.3

​```
kubectl get svc
​```
> This is a blockquote.

gave an out put
nginx-service   NodePort    10.96.169.150   <none>        80:31118/TCP   22m
So the IP to be used is
31118
So took http://172.24.0.3:31118/ in browser

```
kubectl get pods -l app=nginx
```
It lists the pods under the label app=nginx

```
kubectl get pods --watch
```
The above command printed the chnages during the deployment
