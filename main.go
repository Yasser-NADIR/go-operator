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
	corev1 "k8s.io/api/core/v1"
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
		fmt.Printf("Pod Name: %s\n", pod.Name)
	}
}

func CreatePod(clientset *kubernetes.Clientset, PodName string, ImageName string, Namespace string){
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: PodName,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name: PodName,
					Image: ImageName,
				},
			},
		},
	}

	createPod, err := clientset.CoreV1().Pods(Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil{
		panic(err.Error())
	}

	fmt.Printf("Created pod: %s\n", createPod.Name)
}

func main(){
	namespace := "default"
	podName := "nginx-1"
	imageName := "nginx"

	config := CreateK8sClientConfig()
	clientset := CreateK8sClientset(config)

	ListPods(clientset, namespace)
	CreatePod(clientset, podName, imageName, namespace) 
	ListPods(clientset, namespace)

}