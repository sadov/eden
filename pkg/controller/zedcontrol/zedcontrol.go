package zedcontrol

import (
	"encoding/json"
	"fmt"
	"github.com/lf-edge/adam/pkg/x509"
	"github.com/lf-edge/eden/pkg/controller/einfo"
	"github.com/lf-edge/eden/pkg/controller/elog"
	"github.com/lf-edge/eden/pkg/defaults"
	"github.com/lf-edge/eden/pkg/utils"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

type Ctx struct {
	url      string
	login    string
	password string
	dir      string
}

//EnvRead use variables from viper for init controller
func (zedcontrol *Ctx) InitWithVars(vars *utils.ConfigVars) error {
	zedcontrol.url = vars.ZedControlAddress
	zedcontrol.login = vars.ZedControlLogin
	zedcontrol.password = vars.ZedControlPassword
	zedcontrol.dir = vars.ZedControlDist
	if zedcontrol.login == "" || zedcontrol.password == "" {
		log.Fatal("login and password must be provided for zedcontrol in config file")
	}
	res, err := zedcontrol.callZCLI(fmt.Sprintf("zcli configure --server=%s --user=%s --password=%s --output=json", zedcontrol.url, zedcontrol.login, zedcontrol.password))
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("zcli configure %s", res)
	res, err = zedcontrol.callZCLI("zcli login")
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("zcli login %s", res)
	return nil
}

//GetDir return dir
func (zedcontrol *Ctx) GetDir() (dir string) {
	return zedcontrol.dir
}

//getLogsRedisStream return info stream for devUUID for load from redis
func (zedcontrol *Ctx) getLogsRedisStream(devUUID uuid.UUID) (dir string) {
	return fmt.Sprintf("%s%s", defaults.DefaultLogsRedisPrefix, devUUID.String())
}

//Register device in zedcontrol
func (zedcontrol *Ctx) Register(eveCert string, eveSerial string) error {
	b, err := ioutil.ReadFile(eveCert)
	switch {
	case err != nil && os.IsNotExist(err):
		log.Printf("cert file %s does not exist", eveCert)
		return err
	case err != nil:
		log.Printf("error reading cert file %s: %v", eveCert, err)
		return err
	}

	cert, err := x509.ParseCert(b)
	if err != nil {
		return err
	}
	//TODO move magic numbers
	res, err := zedcontrol.callZCLI(fmt.Sprintf("zcli edge-node create %s --project=%s --model=%s --onboarding-key=%s --serial=%s", "eden_test", "default-project", "ZedVirtual-4G", cert.Subject.CommonName, eveSerial))
	if err != nil {
		return err
	}
	log.Debugf("zcli edge-node create %s", res)
	res, err = zedcontrol.callZCLI(fmt.Sprintf("zcli edge-node activate %s", "eden_test"))
	if err != nil {
		return err
	}
	log.Debugf("zcli edge-node activate %s", res)
	return nil
}

type adminState int

var (
	adminStateAny        adminState = -1
	adminStatePending    adminState = 0
	adminStateCreated    adminState = 1
	adminStateDeleted    adminState = 2
	adminStateActive     adminState = 3
	adminStateInactive   adminState = 4
	adminStateRegistered adminState = 5
)

type devEl struct {
	ID         string     `json:"devID"`
	AdminState adminState `json:"adminState"`
	OpState    int        `json:"opState"`
}
type devList struct {
	List []devEl `json:"list"`
}

func (zedcontrol *Ctx) deviceListFiltered(adminState adminState) (out []string, err error) {
	nodes, err := zedcontrol.getObj("zcli edge-node show")
	var devices devList
	if err := json.Unmarshal([]byte(nodes), &devices); err != nil {
		return nil, err
	}
	for _, el := range devices.List {
		if adminState != adminStateAny {
			if el.AdminState == adminState {
				out = append(out, el.ID)
			}
		} else {
			out = append(out, el.ID)
		}
	}
	return out, nil
}

//OnBoardList return onboard list
func (zedcontrol *Ctx) OnBoardList() (out []string, err error) {
	return zedcontrol.deviceListFiltered(adminStateAny)
}

//DeviceList return device list
func (zedcontrol *Ctx) DeviceList() (out []string, err error) {
	return nil, nil
	//return zedcontrol.deviceListFiltered(adminStateRegistered) //TODO if return not empty, then cannot onboard
}

//ConfigSet set config for devID
func (zedcontrol *Ctx) ConfigSet(devUUID uuid.UUID, devConfig []byte) (err error) {
	return fmt.Errorf("not implemented now")
}

//ConfigGet get config for devID
func (zedcontrol *Ctx) ConfigGet(devUUID uuid.UUID) (out string, err error) {
	return "", fmt.Errorf("not implemented now")
}

//LogChecker check logs by pattern from existence files with LogLast and use LogWatchWithTimeout with timeout for observe new files
func (zedcontrol *Ctx) LogChecker(devUUID uuid.UUID, q map[string]string, handler elog.HandlerFunc, mode elog.LogCheckerMode, timeout time.Duration) (err error) {
	return fmt.Errorf("not implemented now")
}

//LogLastCallback check logs by pattern from existence files with callback
func (zedcontrol *Ctx) LogLastCallback(devUUID uuid.UUID, q map[string]string, handler elog.HandlerFunc) (err error) {
	return fmt.Errorf("not implemented now")
}

//InfoChecker checks the information in the regular expression pattern 'query' and processes the info.ZInfoMsg found by the function 'handler' from existing files (mode=einfo.InfoExist), new files (mode=einfo.InfoNew) or any of them (mode=einfo.InfoAny) with timeout.
func (zedcontrol *Ctx) InfoChecker(devUUID uuid.UUID, q map[string]string, infoType einfo.ZInfoType, handler einfo.HandlerFunc, mode einfo.InfoCheckerMode, timeout time.Duration) (err error) {
	return fmt.Errorf("not implemented now")
}

//InfoLastCallback check info by pattern from existence files with callback
func (zedcontrol *Ctx) InfoLastCallback(devUUID uuid.UUID, q map[string]string, infoType einfo.ZInfoType, handler einfo.HandlerFunc) (err error) {
	return fmt.Errorf("not implemented now")
}
