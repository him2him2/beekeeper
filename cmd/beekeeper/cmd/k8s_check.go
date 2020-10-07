package cmd

import (
	"flag"
	"fmt"

	"github.com/ethersphere/beekeeper/pkg/k8s"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func (c *command) initK8SCheck() *cobra.Command {
	return &cobra.Command{
		Use:   "check",
		Short: "k8s check",
		Long:  `k8s check.`,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			kubeconfig := flag.String("kubeconfig", c.config.GetString(optionNameK8SConfig), "kubeconfig file")
			flag.Parse()

			config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
			if err != nil {
				fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
				return err
			}

			clientset, err := kubernetes.NewForConfig(config)
			if err != nil {
				fmt.Printf("Client cannot be set: %v\n", err)
				return err
			}

			fmt.Println("k8s")

			return k8s.Check(clientset, c.config.GetString(optionNameK8SNamespace))
		},
		PreRunE: c.k8sPreRunE,
	}
}
