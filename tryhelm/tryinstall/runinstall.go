package main

import (
	"fmt"
	"os"
	"path/filepath"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/kube"

	"k8s.io/client-go/util/homedir"
	//_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func getActionConfig(releaseNS string) *action.Configuration {
	actionConfig := new(action.Configuration)
	kubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")
	if err := actionConfig.Init(kube.GetConfig(kubeconfigPath, "", releaseNS), releaseNS, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
		fmt.Printf(format, v)
	}); err != nil {
		panic(err)
	}
	return actionConfig
}

func main() {
	chartPath := "harbor/library/nginx"
	chart, err := loader.Load(chartPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(chart)
	releaseName := "n1"
	releaseNamespace := "default"
	actionConfig := getActionConfig(releaseNamespace)
	iCli := action.NewInstall(actionConfig)
	iCli.Namespace = releaseNamespace
	iCli.ReleaseName = releaseName
	rel, err := iCli.Run(chart, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully installed release: ", rel.Name)
}
