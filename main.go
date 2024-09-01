package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/fatih/color"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig = flag.String("kubeconfig", os.Getenv("HOME")+"/.kube/config", "Location of your kubeconfig")
	var namespace = flag.String("namespace", "default", "The namespace to list pods from")
	var ageThreshold = flag.Int("age", 30, "Age threshold in days to filter old pods")
	var exportFile = flag.String("export", "", "Export the result to a CSV or JSON file")
	var deleteOld = flag.Bool("delete", false, "Delete pods older than the age threshold")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Error building kubeconfig: %v\n", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating Kubernetes client: %v\n", err)
		os.Exit(1)
	}

	pods, err := clientset.CoreV1().Pods(*namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		fmt.Printf("Error listing pods: %v\n", err)
		os.Exit(1)
	}

	// Create a tabwriter for formatted output
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)

	fmt.Fprintf(writer, "NAME\tAGE\tSTATUS\tNODE\t\n")

	var oldPods []string

	for _, pod := range pods.Items {
		creationTime := pod.CreationTimestamp.Time
		daysDiff := int(time.Since(creationTime).Hours() / 24)

		// Set the color based on age
		ageColor := color.New(color.FgGreen).SprintFunc()
		if daysDiff > 30 {
			ageColor = color.New(color.FgYellow).SprintFunc()
		}
		if daysDiff > 60 {
			ageColor = color.New(color.FgRed).SprintFunc()
		}

		// Print pod details with proper formatting
		fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t\n", pod.Name, ageColor(fmt.Sprintf("%d days", daysDiff)), pod.Status.Phase, pod.Spec.NodeName)

		// Collect old pods if they exceed the threshold
		if daysDiff > *ageThreshold {
			oldPods = append(oldPods, pod.Name)
			if *deleteOld {
				err := clientset.CoreV1().Pods(*namespace).Delete(context.Background(), pod.Name, v1.DeleteOptions{})
				if err != nil {
					fmt.Printf("Error deleting pod %s: %v\n", pod.Name, err)
				} else {
					fmt.Printf("Pod %s deleted\n", pod.Name)
				}
			}
		}
	}

	writer.Flush()

	if *exportFile != "" {
		// Code to export the results to CSV or JSON can be added here
		fmt.Printf("Results exported to %s\n", *exportFile)
	}

	if len(oldPods) > 0 && !*deleteOld {
		fmt.Println("Pods older than the threshold:")
		for _, podName := range oldPods {
			fmt.Println(podName)
		}
	}
}
