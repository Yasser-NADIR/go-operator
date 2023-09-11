# go-operator
this repo is for learning go/operator to create an api in go to create custom resource definition 


1. first create an api in go to receive the definition of the new CRD
2. this api will use kubebuilder to create an CRD
3. this api will apply this CRD to k8s vis API k8s
4. this api will use kubebuilder to create controller manager for this CRD
5. apply this controller manager to k8s

for the moment we are going to learn how to communicate with k8s using Golang eg Pod:

- listing elements
- create element
- update element
- delete element
- then testing for other elements (deployment, service, ingress)
- trying to deploy whole application using just go

then trying with operators:

- first trying to create a CRD
- then trying with controller manager, it will be easy now to understand if we tried to manipulate k8s with go


after we try each part, now we will try to combine both of them:

- create functions to create CRD using kubebuilder
- create functions to insert CRD into k8s
- create function to create controller manager
- create function to build and run controller manager

some links to understand operators:

- https://www.youtube.com/watch?v=8Ex7ybi273g&ab_channel=VMware%7Bcode%7D
- https://www.youtube.com/watch?v=KBTXBUVNF2I&ab_channel=CNCF%5BCloudNativeComputingFoundation%5D
- https://pres.metamagical.dev/kubecon-us-2019/#slide-22
- https://github.com/DirectXMan12/kubebuilder-workshops/blob/kubecon-us-2019/api/v1/guestbook_types.go
- https://book-v1.book.kubebuilder.io/getting_started/installation_and_setup.html
- https://www.sobyte.net/post/2022-08/go-k8s-operators-part1/
