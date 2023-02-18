package api

import (
	"context"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewDevelopment()
}

func ListJobs() (*batchv1.JobList, error) {
	clientset, err := NewClientSet()
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	jobs, err := clientset.BatchV1().Jobs("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return jobs, nil
}

func NewClientSet() (*kubernetes.Clientset, error) {
	kubeconfigPath, ok := os.LookupEnv("KUBECONFIG")
	if !ok {
		kubeconfigPath = filepath.Join(homedir.HomeDir(), ".kube", "config")
	}
	logger.Info(kubeconfigPath)

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		logger.Warn("fatal kubeconifg")
		config, err = rest.InClusterConfig()
		if err != nil {
			logger.Error("fatal in-cluster config")
			return nil, err
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return clientset, nil
}
