// Save this as main.go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    v1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
)

func main() {
    config, err := rest.InClusterConfig()
    if err != nil {
        log.Fatal(err.Error())
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        log.Fatal(err.Error())
    }

    for {
        podList, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{
            FieldSelector: "status.phase=Pending",
        })
        if err != nil {
            log.Fatal(err.Error())
        }

        for _, pod := range podList.Items {
            if pod.Spec.SchedulerName == "my-custom-scheduler" {
                fmt.Printf("Scheduling pod: %s\n", pod.Name)

                nodeList, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
                if err != nil {
                    log.Fatal(err.Error())
                }

                if len(nodeList.Items) == 0 {
                    fmt.Printf("No available nodes for scheduling pod: %s\n", pod.Name)
                    continue
                }

                node := nodeList.Items[0]

                bind := &v1.Binding{
                    ObjectMeta: metav1.ObjectMeta{Name: pod.Name},
                    Target: v1.ObjectReference{
                        Kind: "Node",
                        Name: node.Name,
                    },
                }

                err = clientset.CoreV1().Pods(pod.Namespace).Bind(context.TODO(), bind, metav1.CreateOptions{})
                if err != nil {
                    log.Fatal(err.Error())
                }

                fmt.Printf("Pod %s scheduled to node %s\n", pod.Name, node.Name)
            }
        }

        time.Sleep(10 * time.Second)
    }
}

