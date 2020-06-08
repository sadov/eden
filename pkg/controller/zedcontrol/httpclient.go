package zedcontrol

import (
	"fmt"
	"github.com/lf-edge/eden/pkg/defaults"
	"github.com/lf-edge/eden/pkg/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (zedcontrol *Ctx) callZCLI(command string) (out string, err error) {
	configPath := filepath.Join(zedcontrol.dir, "config")
	if err = os.MkdirAll(configPath, 0755); err != nil {
		return "", err
	}
	volumeMap := map[string]string{"/root/.config/zededa": filepath.Join(zedcontrol.dir, "config")}
	u, err := utils.RunDockerCommand("zededa/zcli", command, volumeMap)
	if err != nil {
		log.Printf("error callZCLI: %v", err)
		return "", err
	}
	return u, nil
}

func (zedcontrol *Ctx) getObj(path string) (out string, err error) {
	return zedcontrol.callZCLI(path)
}

func (zedcontrol *Ctx) getList(path string) (out []string, err error) {
	buf, err := zedcontrol.getObj(path)
	return strings.Fields(buf), err
}

func (zedcontrol *Ctx) postObj(path string, obj []byte) (err error) {
	return fmt.Errorf("not supported")
}

func (zedcontrol *Ctx) putObj(path string, obj []byte) (err error) {
	return fmt.Errorf("not supported")
}

func repeatableAttempt(client *http.Client, req *http.Request) (response *http.Response, err error) {
	maxRepeat := defaults.DefaultRepeatCount
	delayTime := defaults.DefaultRepeatTimeout

	for i := 0; i < maxRepeat; i++ {
		timer := time.AfterFunc(2*delayTime, func() {
			i = 0
		})
		resp, err := client.Do(req)
		if err == nil {
			return resp, nil
		}
		log.Debugf("error %s URL %s: %v", req.Method, req.RequestURI, err)
		timer.Stop()
		log.Infof("Attempt to re-establish connection with controller (%d) of (%d)", i, maxRepeat)
		time.Sleep(delayTime)
	}
	return nil, fmt.Errorf("all connection attempts failed")
}
