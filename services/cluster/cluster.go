package cluster

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"github.com/briandowns/spinner"
)

func Cluster(appName string, servicePort int32) {
	// create a new spinner with custom settings
	s := spinner.New(spinner.CharSets[20], 100*time.Millisecond)
	s.Suffix = " Deploying on the cluster..."
	s.Start()
	defer s.Stop()

	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Printf("Error building config from flags: %v\n", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating kubernetes client: %v\n", err)
		os.Exit(1)
	}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: appName+"-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": appName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": appName,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  appName+"-container",
							Image: "itsdarshankumar/"+appName+":latest",
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: servicePort,
									HostPort:      9000,
									Protocol:      corev1.ProtocolTCP,
								},
							},
						},
					},
				},
			},
		},
	}
	// Port forwarding
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: appName+"-service",
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": appName,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       9000,
					TargetPort: intstr.FromInt(int(servicePort)),
				},
			},
			Type: corev1.ServiceTypeNodePort,
		},
	}
	svcClient := clientset.CoreV1().Services("default")
	svc, err = svcClient.Create(context.Background(), svc, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating service: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Service created: %s:%d\n", svc.Spec.ClusterIP, svc.Spec.Ports[0].NodePort)

	deploymentsClient := clientset.AppsV1().Deployments("default")
	_, err = deploymentsClient.Create(context.Background(), deployment, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating deployment: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Deployment created!")
}

func int32Ptr(i int32) *int32 { return &i }
