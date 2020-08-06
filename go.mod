module github.com/beautytiger/go-playground

go 1.13

require (
	github.com/astaxie/beego v1.12.1
	github.com/coreos/etcd v3.3.13+incompatible
	github.com/gin-gonic/gin v1.6.2
	github.com/go-redis/redis/v8 v8.0.0-beta.6
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gophercloud/gophercloud v0.11.0 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/olivere/elastic/v7 v7.0.15
	github.com/prometheus/client_golang v1.5.1
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/sirupsen/logrus v1.5.0
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.3
	github.com/stretchr/testify v1.6.1
	github.com/urfave/cli v1.22.4
	github.com/urfave/negroni v1.0.0
	go.uber.org/zap v1.15.0
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a
	google.golang.org/grpc/examples v0.0.0-20200721210703-a5a36bd3f0bb // indirect
	gopkg.in/yaml.v2 v2.2.8
	helm.sh/helm/v3 v3.2.4
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v0.18.2
	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20200414100711-2df71ebbae66 // indirect
	rsc.io/letsencrypt v0.0.3 // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.29.0
