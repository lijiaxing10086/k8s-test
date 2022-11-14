package client

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

type Client struct {
	Clientset kubernetes.Interface
}

func (c Client) CreatePod(pod *v1.Pod) (*v1.Pod, error) {
	pod, err := c.Clientset.CoreV1().Pods(pod.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		klog.Errorf("Error occured while creating pod %s: %s", pod.Name, err.Error())
		return nil, err
	}

	klog.Infof("Pod %s is succesfully created", pod.Name)
	return pod, nil
}
func (c Client) GetPod() (*v1.PodList, error) {
	podlist, err := c.Clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return podlist, err
}
