package fetch

import (
	"fmt"
	"os"
	"path/filepath"
	"context"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/spf13/cobra"
)

func Fetch(cmd *cobra.Command, args []string){
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

	deploymentsClient := clientset.AppsV1().Deployments("default")
	listOptions := metav1.ListOptions{}

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		deployments, err := deploymentsClient.List(context.Background(),listOptions)
		if err != nil {
			return err
		}

		if len(deployments.Items) == 0 {
			fmt.Println("No resources found in default namespace.")
			return nil
		}

		fmt.Printf("%-25s%-10s%-10s\n", "NAME", "READY", "UP-TO-DATE")
		for _, d := range deployments.Items {
			fmt.Printf("%-25s%-10d%-10d\n", d.Name, d.Status.ReadyReplicas, d.Status.UpdatedReplicas)
		}

		return nil
	})

	if retryErr != nil {
		fmt.Printf("Error listing deployments: %v\n", retryErr)
		os.Exit(1)
	}
}
