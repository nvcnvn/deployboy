package main

import (
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	println(clientcmd.RecommendedConfigPathEnvVar)
}