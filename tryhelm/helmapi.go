package main

import (
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/kube"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/repo"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

func main() {
	//getRepoIndexAndSaveToFile()
	//getReleaseHistory()
	//getReleaseStatus()
	//getAllRelease()
	//getReleaseInfo()
	getReleaseValue()
}

// 返回基本的chartrepo配置
func httpChartRepo() *repo.ChartRepository {
	//r := repo.NewFile()
	entry := &repo.Entry{
		Name:                  "harbor",
		URL:                   "http://harbor.beautytiger.com/chartrepo",
		Username:              "admin",
		Password:              "Harbor12345",
		InsecureSkipTLSverify: true,
	}
	//r.Add(entry)
	var httpProvider = []getter.Provider{
		{
			Schemes: []string{"http", "https"},
			New:     getter.NewHTTPGetter,
		},
	}
	chartRepo, err := repo.NewChartRepository(entry, httpProvider)
	if err != nil {
		panic(err)
	}
	return chartRepo
}

//返回action config 的基本配置
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

//打印release对象
func releasePrint(rel *release.Release) {
	fmt.Printf("name: %s version: %d\n", rel.Name, rel.Version)
	fmt.Printf("notes: %v\n", rel.Info.Notes)
	fmt.Printf("status: %v\n", rel.Info.Status)
	fmt.Printf("description: %v\n", rel.Info.Description)
	fmt.Printf("first deploy: %v\n", rel.Info.FirstDeployed)
	fmt.Printf("last deploy: %v\n", rel.Info.LastDeployed)
	fmt.Printf("deleted: %v\n", rel.Info.Deleted)
	fmt.Println()
}

//获取repo的索引文件，存到本地，输出文件名
//helm repo update
//pass 获取所有的charts，harbor有相关接口
func getRepoIndexAndSaveToFile() {
	chartRepo := httpChartRepo()
	f, err := chartRepo.DownloadIndexFile()
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	// 输出文件地址
	// /home/matrix/.cache/helm/repository/harbor-index.yaml
}

//展示chart详情，比如获取values文件
//helm show
//pass, 获取chart详情，harbor有接口

//根据名称搜索chart
//helm search
//pass，可使用harbor接口遍历再搜索，部署文件不会很多

//下载chart文件
//helm pull
//pass，无使用场景

//安装chart
//helm install

//升级chart
//helm upgrade

//查看release的发布历史
//helm history
func getReleaseHistory() {
	rName := "r1"
	rNs := "default"
	actionConfig := getActionConfig(rNs)
	histCmd := action.NewHistory(actionConfig)
	rel, err := histCmd.Run(rName)
	if err != nil {
		fmt.Printf("release not found: %v", err)
	}
	for _, r := range rel {
		releasePrint(r)
	}
}

//检查部署状态
//helm status
func getReleaseStatus() {
	rName := "r1"
	rNs := "default"
	actionConfig := getActionConfig(rNs)
	statusCmd := action.NewStatus(actionConfig)
	rel, err := statusCmd.Run(rName)
	if err != nil {
		fmt.Printf("release not found: %v", err)
	}
	releasePrint(rel)
}


//获取release详细信息
//helm get
func getReleaseInfo() {
	rName := "r1"
	rNs := "default"
	actionConfig := getActionConfig(rNs)
	getCmd := action.NewGet(actionConfig)
	rel, err := getCmd.Run(rName)
	if err != nil {
		fmt.Printf("release not found: %v", err)
	}
	releasePrint(rel)
}
func getReleaseValue() {
	rName := "r1"
	rNs := "default"
	actionConfig := getActionConfig(rNs)
	getCmd := action.NewGetValues(actionConfig)
	values, err := getCmd.Run(rName)
	if err != nil {
		fmt.Printf("release not found: %v", err)
	}
	fmt.Println(values)
}


//获取所有release
//helm list
func getAllRelease() {
	rNs := "default"
	actionConfig := getActionConfig(rNs)
	listCmd := action.NewList(actionConfig)
	rel, err := listCmd.Run()
	if err != nil {
		fmt.Printf("release not found: %v", err)
	}
	for _, r := range rel {
		releasePrint(r)
	}
}

//获取helm的底层资源对象
//helm get manifest myhelmredis | kubectl get -f -
//略难，最后实现

//卸载应用
//helm uninstall

//发布回滚
//helm rollback

//helm driver sql
//https://github.com/helm/helm/pull/5371
//https://github.com/helm/helm/pull/7635
