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
		fmt.Println("Pod Name: ", pod.Name)
		fmt.Println("Pod Image: ", pod.Spec.Containers[0].Image)
		fmt.Println()
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

	fmt.Printf("Created pod: %s\n\n", createPod.Name)
}

func UpdatePod(clientset *kubernetes.Clientset, podname string, newimage string, namespace string){
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podname, metav1.GetOptions{})
	if err != nil{
		panic(err.Error())
	}

	pod.Spec.Containers[0].Image = newimage

	Updatedpod, err := clientset.CoreV1().Pods(namespace).Update(context.TODO(), pod, metav1.UpdateOptions{})
	if err != nil{
		panic(err.Error())
	}

	fmt.Printf("updatd pod: %s\n\n", Updatedpod.Name)
}

func DeletePod(clientset *kubernetes.Clientset, podname string, namespace string){
	err := clientset.CoreV1().Pods(namespace).Delete(context.TODO(), podname, metav1.DeleteOptions{})
	if err != nil{
		panic(err.Error())
	}

	fmt.Printf("Deleted Pod: %s\n", podname)
}

func main(){
	namespace := "default"
	podName := "nginx-1"
	imageName := "nginx"
	newImageName := "nginx:1.21"

	config := CreateK8sClientConfig()
	clientset := CreateK8sClientset(config)

	ListPods(clientset, namespace)
	CreatePod(clientset, podName, imageName, namespace) 
	ListPods(clientset, namespace)
	//in updated function an error will appear because the pod didn't complite its creation
	//in future steps I should fix it
	UpdatePod(clientset, podName, newImageName, namespace)
	ListPods(clientset, namespace)
	DeletePod(clientset, podName, namespace) 
	//here after deleting pod the list pod will show the the pod didn't delete
	//it because the state of deleted pod didn't change in etcd
	//after we checked using kubectl we found that the pod deleted
	ListPods(clientset, namespace)
}