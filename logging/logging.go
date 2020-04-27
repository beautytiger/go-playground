package main

import (
	"flag"
	"k8s.io/klog"
)


func main() {
	klog.InitFlags(nil)
	flag.Parse()
	str := "message"
	klog.Error("klog.Error\n")
	klog.Errorf("klog.Errorf: %s\n", str)
	klog.Errorln("klog.Errorln\n")
	klog.Warning("klog.Warning\n")
	klog.Warningf("klog.Warningf\n")
	klog.Warningln("klog.Warningln\n")
	klog.Info("klog.Info\n")
	klog.Infof("klog.Infof\n")
	klog.Infoln("klog.Infoln\n")
	klog.V(0).Infof("klog.V(0).Infof\n")
	klog.V(1).Infof("klog.V(1).Infof\n")
	klog.V(2).Infof("klog.V(2).Infof\n")
	klog.V(3).Infof("klog.V(3).Infof\n")
	klog.V(4).Infof("klog.V(4).Infof\n")
	klog.V(5).Infof("klog.V(5).Infof")
	klog.V(6).Infof("klog.V(6).Infof")
	klog.V(7).Infof("klog.V(7).Infof")
	klog.V(8).Infof("klog.V(8).Infof")
	klog.V(9).Infof("klog.V(9).Infof")
	klog.Info(`Invalid value for cacheReservation.
				The minimum percentage is %d and maximum percentage is %d.`)
}
