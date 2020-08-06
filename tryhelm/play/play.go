package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
	"k8s.io/client-go/tools/clientcmd/api"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func Log(format string, v ...interface{}) {
	fmt.Printf(format, v)
}

var CacheDir string= "/tmp/.kubecache/"
var KubeConfig string = filepath.Join(CacheDir, ".kube/config")
var HelmPluginDir string = filepath.Join(CacheDir, ".helm", "data/plugins")
var HelmRegisteryConf string = filepath.Join(CacheDir, ".helm", "conf/registery.json")
var HelmRepoConf string = filepath.Join(CacheDir, ".helm", "conf/repositories.yaml")
var HelmCacheDir string = filepath.Join(CacheDir, ".helm", "cache")

var clientCert string = `LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURLekNDQWhPZ0F3SUJBZ0lVUXNHSmZpUUtFRHdybXZRNnlTRUxjVko5amZVd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0ZURVRNQkVHQTFVRUF4TUthM1ZpWlhKdVpYUmxjekFlRncweE9URXlNVEV4TURFMU1EQmFGdzB5T1RFeQpNVGd4TURFMU1EQmFNRFF4RnpBVkJnTlZCQW9URG5ONWMzUmxiVHB0WVhOMFpYSnpNUmt3RndZRFZRUURFeEJyCmRXSmxjbTVsZEdWekxXRmtiV2x1TUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUEKMUhsdHJkSHpnQyszaGVJa3E2dVhoQldrT3JKOHlGbkpLYVVNTGhmeW93RDVPd1JPakhDNHNONWxhNnZrNUhQSgovb0pwOERwNngvNCtZeHlUTnBRMTVyY1NjTFZxem5rRG5XZ1dSK0tMQWkxM2FLc1ovN3ZiekJFaHRhUk9lMVdMCjdSeXZqYUE4SHZ1QW9ucDVSS3FxaTBobVdTVXh5NE1WMURaUmk5WW9MUnZDeTNHcjhwYXlkczdBTXhxTnliV0UKQnhwNi9uQ0ZheElGY281WG0yNzB5QUpxZlBMUHZTUVRNcmZYQ01peXhLY1JEeVpwT3VDMmJMWkh6QWtwSkVLMQpndHpLTzY3WUREbVF0UGh6QU9helVCLzFzZjdSZjVYeUEvRVZVUHF1ME53VFg5bXFjOG5UZTdUOVpSZEliT1NFCi81ZEttNkJJeWhhajl6UXM5WW9mcXdJREFRQUJvMVF3VWpBT0JnTlZIUThCQWY4RUJBTUNCYUF3RXdZRFZSMGwKQkF3d0NnWUlLd1lCQlFVSEF3SXdEQVlEVlIwVEFRSC9CQUl3QURBZEJnTlZIUTRFRmdRVXZJSlhyZkcvWVk1MApjOGV0Sk4vVTZKSjA0bW93RFFZSktvWklodmNOQVFFTEJRQURnZ0VCQUhBblBzL2V6WldiSlNhOFI1MHcvcUp2CnZUU2p4Sy9PL2NFYzNzQko4UXk3bTBhQzcrT2xjUDl2RlR6TjY4ZFBCbFczYUFqUEZXbklBWmNXV0hKTG8wTHgKOU1TTUMzNWRkTm9JcTNyNU03cXhBdFRxbTl2QUZxdjI4eDJaeGxhWlRCNDJUSjEvbmNBaGJyY0R2M1hXbXNnTQoyckF1cGdQblByMHpFYy9wMllEamczZEJsU2FWMzlYWkdKbVljTlBDZDVjL0N4MDNvdVVHbVY3SHVUS216Mm90CkpEZnJ4YjR5VUlzTDNWRXpoZHFMakc0ZU1RWkNac3h2WHFyditJNEFueXlVWGlLWXFyV0FadHB5RmhkYWZDK0kKVTJ3OG1vQlQ3bDBSRERFRENPMDM5YzNBZG96TXNRaTh4SWlQdHJQeWdBR1B0Z3FSWWZMMjBIRHhlMWxIRENJPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==`
var clientKey string = `LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBMUhsdHJkSHpnQyszaGVJa3E2dVhoQldrT3JKOHlGbkpLYVVNTGhmeW93RDVPd1JPCmpIQzRzTjVsYTZ2azVIUEovb0pwOERwNngvNCtZeHlUTnBRMTVyY1NjTFZxem5rRG5XZ1dSK0tMQWkxM2FLc1oKLzd2YnpCRWh0YVJPZTFXTDdSeXZqYUE4SHZ1QW9ucDVSS3FxaTBobVdTVXh5NE1WMURaUmk5WW9MUnZDeTNHcgo4cGF5ZHM3QU14cU55YldFQnhwNi9uQ0ZheElGY281WG0yNzB5QUpxZlBMUHZTUVRNcmZYQ01peXhLY1JEeVpwCk91QzJiTFpIekFrcEpFSzFndHpLTzY3WUREbVF0UGh6QU9helVCLzFzZjdSZjVYeUEvRVZVUHF1ME53VFg5bXEKYzhuVGU3VDlaUmRJYk9TRS81ZEttNkJJeWhhajl6UXM5WW9mcXdJREFRQUJBb0lCQUM2cWhOWmFlSHE5QURmZApsV1JIQ2I0dTZxVWhzcHBtYmZKQmw4MC80VUMvNC8xOFVxd3h6YjY2K2RlT0ozV3RvTDNQY3VrYVR3RkI2LzNrCnRzaVVBTXF0aHRWT1l5ajZGenNwTTB1U2pDbHNSditzSk1ld1FQUXZCeEpZQmd5OG0wRDJ2ZkFETTNFeWwrd1gKcmk3UXRTaG1QR1J6OE4zYjZwMXpXTEczMVpjaVRVd0JBaXhYN1pKRzl3SHVHYWxYRXBDNmoxVW9xdHl4cG5LcgpFaHdHRkxqZ2NYbEhvTU4wK0ZCM2tBNGpqTkFHdUZkTVhXR3hzQ1ZmK0R2Z0Y4QnRWSXVDVndnN0J1YWFLc1VlClArMmUxZTJRS25xcS9TSkM3d3JqSkdYUlZ2ZVk5NUtzdGZRUlRLZHlLS0FwU29tQTZWUHZTUWRaTG1qU3piSUEKSDJPdzNBRUNnWUVBN2hQMnZEMkZVYUZlNlV4ekhaaDF4d2IzUEd4dG5aWi9sNi9YWjc0QmdTZ0JRampOYlpmUQovUlRIMWNNUW9VV0llTFpEcURYNkxSaUlONEgvVHBiS3BhdkZYdWY4QlVlVnZRL1E4M211ZTRVSExQVXZxNVI2CjJ5S0VSbW9sRHMzN04weHlMRVZiMHpuYzRyRzdiTTgxai9ROGNIa2NJV2RHYXN1QW1MaU13cXNDZ1lFQTVIZ04KaE10VlNKdE9Lc0tRNU5ac3ZJcXFEaGZLb2dSZWxSaUUxdzJOQmw4Ulh3T1E4b3JQdzc3M2dFT09aZ1R4am5WVQpXb296QS8zOXFNZzhnemZVOUZSUnJJUnR5L0g1SXJzVlJIdlMwSmk2UE1EV2pOOVVzTENubm1iMnNRaGtsMjVzCmYwVXNKYnd4R2FDNHJ5T0xhRmsxSVY3N2JQWW83VGF5S3NhQUZ3RUNnWUFOSWQ3SFRXRndPa1dBK3UyU29WbEYKekRXUkNLNFhOamo5aGY2TXcyZFQxNEFTUmczVzBMdi9hY2tJR29WdFQrZVNPUU9NWU1YVy9QWnVrV1lpMEtDRworTzJmS1k2MlA5ZWJvU2EyZDJ6UGRXd0s1ODlrV0lGYVd4SU96RVlWQmZtb2VEZ0lCeDlrakZtSnk2SDZBNGdCCjA4Z05zRENFQ3kxUU1MOThMUm5lbHdLQmdRRFBhTTJkQ0pyYjEySVhINXM2cEE4Qyt6OTJtOUt4VXhZcWRKL2sKTzhFVFQ2c05mc2RQNlFURXg5RVkrRHVHRW1iWTFWdEc5cjlwbUdOM25wQ1E3MWE4bU8rc2xteTNBVzBUUXBRRQpJQVU2cWV3b3kwZjdpNlB1NC9ESHRPR3Z4ZGJNUXNyc24zZDhxbVJJeDhmaDkzclB1R3lyWnZjSjdFOHJTc1JrCllPQllBUUtCZ1FERUs0NjN6Mk5VNVVFeWR6N1ZUOGdVUzhzTlg0bjFkRnlGWHJJdFYwVDlxcTBZc2VVdzByWGQKVmZFd1JvZHNLRVZDME8ydnNuMnBYK2syeUQ0NEJlMXYzU3U4ekRaV0k1bHcxZEtWYUZGU3lvQVVrWnU3LzZHcgpNMWpHN3RvSUFNUzVMSnUvZ3VhWWlIZDl3dEtRdys3eE82SXF2VlM3S1ZtOWtsQnpzMnJkRWc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=`
var clusterCA string = `LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN5RENDQWJDZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRFNU1USXdPVEV5TVRVek1sb1hEVEk1TVRJd05qRXlNVFV6TWxvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTnQ1ClcwRWcrbFc4anl6aXR0anUrK2RTK29DajFKU0lQWUNVVCtFTnhKYTFFdE5DYk1lRXpWNlZYY1FsOERJME5DZWYKUWF6bWlTSVVyRmNiazBuQnFRV0NRczlsWGNEOG9LaHQ3ZzJ2OVo5UXp3MGJYRkJGN0ZtdlM4ekovZm05bVVXTQpEd1ZTQmJibDFPNGZOL29lZzhtOW83dkZJZ2ROeEhVMlRHMVUwZVdud0s5T3BjL0xGRjJ3R1pYaG1XcWZhbUxnCmE1ZHNuWUdGbCt0eTVRMDFZekprZ0s2aUx3cnFyam55Q1NGZ1czUTFmREtvUzhRYzU2YVlJSnYxMDhVSTYxcTIKRlVoREY0clY2ZnFscUpJaCt2KzBMVWpHanVyenhPRFRGRS9CcEUzakcyYXRqZVdiZ001Yjl0NUtsWEJRcTNXOQpZa2d1QmdjNWduOUV3TCtXaXdjQ0F3RUFBYU1qTUNFd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFHMHZ1N0FLVnFuOWZ6RGFDMlhGUlVXaGlxNlUKOHZQRTRoZC9FWVFsb2w2RVNZdUp3Q2hGcGVJWHZ1U3dGbVIyYVJkaWo3azcyWkxGRkg3NDdEcytBSnFIbFRWKwpwd29qWWtRNUxKOVRRK1V3d1RPM0VLUnJtdVRvSUduYWFRY1RwMmU5cVIwODloNzNrOFlFZmxkQXovYys3YzFHCm4zdnhVRHNjOHdNK1BxYWVnK2wvSXZ4djVpb1haY0QzUmY2b1M5NjlUallYUHhuaVp2dVllS25FUTV5TXdjaEgKUytQaWVGeThLeWxleS9pWFhVUzJpd1p0bG1jb29OeVZjRnJiYk5aYXdYUnMzdTNiV1p2UHNhUkhCNmVkRW5rMgpVVU84OWVMd1hPbWtlQmptdU83dWowampoVStKRUVEc1dGMklGWWZ3RTJHUE52VUQ4OFlzeEU2SkJUUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=`
var clusterURL string = `https://master.kubernetes.cluster:6443`

var ChartValues []string = []string{"image.repository=docker.io/bitnami/nginx", "image.registry=harbor.beautytiger.com"}

type ClusterInfo struct {
	ClusterCA string	`json:"certificateAuthorityData"`
	ClusterURL string	`json:"clusterUrl"`
	ClientKey string	`json:"clientKeyData"`
	ClientCert string	`json:"clientCertificateData"`
	RegistryInfo
}

type RegistryInfo struct {
	Username string		`json:"username"`
	Password string		`json:"password"`
	Url string			`json:"url"`
	Repo string			`json:"repo"`
}

type AppInfo struct {
	AppName string
	ChartName string
	ChartValues []string
	NameSpace string
}

func main() {
	templateConf()
	runInstall()
}

func templateConf() {

	info := ClusterInfo{
		ClusterCA: clusterCA,
		ClusterURL: clusterURL,
		ClientKey: clientKey,
		ClientCert: clientCert,
	}
	templ := `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: {{.ClusterCA}}
    server: {{.ClusterURL}}
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: kubernetes-admin
    namespace: myns
  name: kubernetes-admin@kubernetes
current-context: kubernetes-admin@kubernetes
kind: Config
preferences: {}
users:
- name: kubernetes-admin
  user:
    client-certificate-data: {{.ClientCert}}
    client-key-data: {{.ClientKey}}
`
	tmpl, err := template.New("kubeconf").Parse(templ)
	if err != nil { panic(err) }
	err = os.MkdirAll(filepath.Dir(KubeConfig), 0755)
	if err != nil { panic(err) }
	f, err := os.Create(KubeConfig)
	if err != nil { panic(err) }
	err = tmpl.Execute(f, info)
	if err != nil { panic(err) }
}

func newConfig() {
	conf := &api.Config{
		Kind:        "Config",
		APIVersion:  "v1",
		Preferences: *api.NewPreferences(),
		Contexts: map[string]*api.Context{
			"kubernetes-admin@kubernetes": &api.Context{
				Cluster:  "kubernetes",
				AuthInfo: "kubernetes-admin",
			},
		},
		AuthInfos: map[string]*api.AuthInfo{
			"kubernetes-admin": &api.AuthInfo{
				ClientCertificateData: []byte(clientCert),
				ClientKeyData:         []byte(clientKey),
			},
		},
		Clusters: map[string]*api.Cluster{
			"kubernetes": &api.Cluster{
				Server:                   clusterURL,
				CertificateAuthorityData: []byte(clusterCA),
			},
		},
		CurrentContext: "kubernetes-admin@kubernetes",
	}
	fmt.Println(api.IsConfigEmpty(conf))
	fmt.Printf("%+v", *conf)
	d, err := yaml.Marshal(&conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}

func runInstall() {
	settings := new(cli.EnvSettings)
	//settings.KubeAPIServer = clusterURL
	//settings.KubeContext = contextName
	settings.KubeConfig = KubeConfig
	settings.PluginsDirectory = HelmPluginDir
	settings.RegistryConfig = HelmRegisteryConf
	settings.RepositoryConfig = HelmRepoConf
	settings.RepositoryCache = HelmCacheDir

	cfg := new(action.Configuration)
	err := cfg.Init(settings.RESTClientGetter(), settings.Namespace(), "", Log)
	if err != nil {
		fmt.Println("config init err")
		return
	}
	fmt.Println("config init pass")

	valueOpts := new(values.Options)
	valueOpts.Values = append(valueOpts.Values, ChartValues...)
	client := action.NewInstall(cfg)
	fmt.Printf("init options: %+v\n", client.ChartPathOptions)
	chartPathOptions := action.ChartPathOptions{
		Password: "Harbor12345",
		Username: "admin",
		RepoURL:  "http://harbor.beautytiger.com/chartrepo",
	}
	client.ChartPathOptions = chartPathOptions
	fmt.Printf("new options: %+v\n", client.ChartPathOptions)

	name, chart := "n1", "library/nginx"
	client.ReleaseName = name
	cp, err := client.ChartPathOptions.LocateChart(chart, settings)
	if err != nil {
		fmt.Println("chart locate error")
		fmt.Println(err)
		return
	}
	fmt.Println("chart locate pass")

	p := getter.All(settings)
	vals, err := valueOpts.MergeValues(p)
	if err != nil {
		fmt.Println("merge value error")
		fmt.Println(err)
		return
	}
	fmt.Println("merge value pass")
	fmt.Printf("value info: %+v\n", vals)

	// Check chart dependencies to make sure all are present in /charts
	chartRequested, err := loader.Load(cp)
	if err != nil {
		fmt.Println("load error")
		return
	}
	fmt.Println("load pass")
	//fmt.Printf("chart info: %+v\n", *chartRequested)

	if chartRequested.Metadata.Deprecated {
		fmt.Println("WARNING: This chart is deprecated")
	}
	client.Namespace = settings.Namespace()
	fmt.Println("namespace: ", client.Namespace)

	rel, err := client.Run(chartRequested, vals)
	if err != nil {
		fmt.Println("run error")
		fmt.Println(err)
		return
	}
	fmt.Println("run pass")

	fmt.Printf("release info: %+v\n", *rel)
	return
}
