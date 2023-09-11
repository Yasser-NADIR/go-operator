package main

import (
	"context"
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func CreateK8sClientConfig() (*rest.Config){
	kubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")
	
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil{
		panic(err.Error())
	}

	//this in case of if this code is running inside this cluster
	//rest.InClusterConfig()

	return config
	
}

func CreateK8sClientset(config *rest.Config) (*kubernetes.Clientset){
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func ListPods(clientset *kubernetes.Clientset, Namespace string){
	podList, err := clientset.CoreV1().Pods(Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil{
		panic(err.Error())
	}

	for _, pod := range podList.Items{
		//fmt.Printf("Pod Name: %s\n", pod.Name)
		fmt.Println(pod)
	}
}

func main(){
	config := CreateK8sClientConfig()
	clientset := CreateK8sClientset(config)
	ListPods(clientset, "default")
}