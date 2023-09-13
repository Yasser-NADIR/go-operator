package main

import (
	gok8s "myproject/goK8sOperator/go-k8s"
)


func main(){
	namespace := "default"
	podname := "nginx-1"
	imagename := "nginx"
	newImageName := "nginx:1.21"

	config := gok8s.CreateK8sClientConfig()
	clientset := gok8s.CreateK8sClientset(config)

	gok8s.ListPods(clientset, namespace)
	gok8s.CreatePod(clientset, podname, imagename, namespace) 
	gok8s.ListPods(clientset, namespace)
	gok8s.UpdatePod(clientset, podname, newImageName, namespace)
	gok8s.ListPods(clientset, namespace)
	gok8s.DeletePod(clientset, podname, namespace) 
	//here after deleting pod the list pod will show the the pod didn't delete
	//it because the state of deleted pod didn't change in etcd
	//after we checked using kubectl we found that the pod deleted
	//as I checked in pod.status.condition, and after several tests:
	//the conditions is an array, each part of this array describe an state with a message
	//the message show either the pod is ready or pod is completed
	//if the pod is completed, it is in his destroying phase which it is translate to be deleted
	//it is normal that listpod shows the name of the pod, that is because it isn't deleted but it is destroying
	//it will take a litte bit of time to will not shows again
	//etcd keeps the name of the pod because it not deleted 100%
	//but if we shows the status of the pod, it will be clear the is in phase of destroying
	gok8s.ListPods(clientset, namespace)
}