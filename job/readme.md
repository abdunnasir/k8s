
Run the job
kubectl apply -f job.yaml


restartPolicy: Never
ensures that the container does not restart on failure.


backoffLimit: 4
specifies the number of retries before considering the Job as failed.


To see the list of Jobs:
kubectl get jobs

To see the status of the Job and its pods:
kubectl describe job pi-calculation

To view logs from the pods created by the Job (once they're completed):
kubectl logs <pod_name>
