package defaults

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

const (
	//directories and files
	DefaultDist             = "dist"             //root directory
	DefaultImageDist        = "images"           //directory for images inside dist
	DefaultRedisDist        = "redis"            //directory for volume of redis inside dist
	DefaultAdamDist         = "adam"             //directory for volume of adam inside dist
	DefaultEVEDist          = "eve"              //directory for build EVE inside dist
	DefaultCertsDist        = "certs"            //directory for certs inside dist
	DefaultBinDist          = "bin"              //directory for binaries inside dist
	DefaultEdenHomeDir      = ".eden"            //directory inside HOME directory for configs
	DefaultCurrentDirConfig = "eden-config.yml"  //file for search config in current directory
	DefaultContextFile      = "context.yml"      //file for saving current context inside DefaultEdenHomeDir
	DefaultContextDirectory = "contexts"         //directory for saving contexts inside DefaultEdenHomeDir
	DefaultQemuFileToSave   = "qemu.conf"        //qemu config file inside DefaultEdenHomeDir
	DefaultSSHKey           = "certs/id_rsa.pub" //file for save ssh key
	DefaultConfigHidden     = ".eden-config.yml" //file to save config get --all

	DefaultContext = "default" //default context name

	//domains, ips, ports
	DefaultDomain      = "mydomain.adam"
	DefaultIP          = "192.168.0.1"
	DefaultEVEIP       = "192.168.1.2"
	DefaultEserverPort = 8888
	DefaultTelnetPort  = 7777
	DefaultSSHPort     = 2222
	DefaultEVEHost     = "127.0.0.1"
	DefaultRedisHost   = "localhost"
	DefaultRedisPort   = 6379
	DefaultAdamPort    = 3333

	//tags, versions, repos
	DefaultEVETag            = "5ee6043906449f7fa3447c96fd38dc9a536c5693" //DefaultEVETag tag for EVE image
	DefaultAdamTag           = "0.0.44"
	DefaultRedisTag          = "6"
	DefaultLinuxKitVersion   = "v0.7"
	DefaultImage             = "library/alpine"
	DefaultAdamContainerRef  = "lfedge/adam"
	DefaultRedisContainerRef = "redis"
	DefaultImageTag          = "eden-alpine"
	DefaultEveRepo           = "https://github.com/lf-edge/eve.git"
	DefaultRegistry          = "docker.io"

	//DefaultRepeatCount is repeat count for requests
	DefaultRepeatCount = 20
	//DefaultRepeatTimeout is time wait for next attempt
	DefaultRepeatTimeout         = 5 * time.Second
	DefaultUUID                  = "1"
	DefaultEvePrefixInTar        = "bits"
	DefaultFileToSave            = "./test.tar"
	DefaultIsLocal               = false
	DefaultEVEHV                 = "kvm"
	DefaultQemuCpus              = 4
	DefaultQemuMemory            = 4096
	DefaultEVESerial             = "31415926"
	DefaultImageID               = "1ab8761b-5f89-4e0b-b757-4b87a9fa93ec"
	DefaultDataStoreID           = "eab8761b-5f89-4e0b-b757-4b87a9fa93ec"
	DefaultBaseID                = "22b8761b-5f89-4e0b-b757-4b87a9fa93ec"
	NetDHCPID                    = "6822e35f-c1b8-43ca-b344-0bbc0ece8cf1"
	NetNoDHCPID                  = "6822e35f-c1b8-43ca-b344-0bbc0ece8cf2"
	DefaultTestProg              = ""
	DefaultTestScenario          = ""
	DefaultRootFSVersionPattern  = `^(\d+\.*){2,3}.*-(xen|kvm|acrn)-(amd64|arm64)$`
	DefaultControllerModePattern = `^(?P<Type>(file|proto|adam|zedcloud)):\/\/(?P<URL>.*)$`
	DefaultPodLinkPattern        = `^(?P<TYPE>(docker)):\/\/(?P<TAG>[^:]+):*(?P<VERSION>.*)$`
	DefaultRedisContainerName    = "eden_redis"
	DefaultAdamContainerName     = "eden_adam"
	DefaultDockerNetworkName     = "eden_network"
	DefaultLogLevelToPrint       = log.InfoLevel
	DefaultX509Country           = "RU"
	DefaultX509Company           = "Itmo"
	DefaultLogsRedisPrefix       = "LOGS_EVE_"
	DefaultInfoRedisPrefix       = "INFO_EVE_"

	DefaultAppSubnet = "10.1.0.0/24"

	DefaultZedControl     = "zedcontrol.zededa.net"
	DefaultZedControlDist = "zedcontrol"

	DefaultEVERootCertificate = `-----BEGIN CERTIFICATE-----
MIIF+jCCA+KgAwIBAgIJAPmyNp61XSnEMA0GCSqGSIb3DQEBCwUAMIGJMQswCQYD
VQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTERMA8GA1UEBwwIU2FuIEpvc2Ux
FDASBgNVBAoMC1plZGVkYSBJbmMuMRwwGgYDVQQDDBNaZWRlZGEgSW5jLiBSb290
IENBMR4wHAYJKoZIhvcNAQkBFg9jZXJ0QHplZGVkYS5uZXQwHhcNMTcwMzE3MjAz
ODIyWhcNMzcwMzEyMjAzODIyWjCBiTELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNh
bGlmb3JuaWExETAPBgNVBAcMCFNhbiBKb3NlMRQwEgYDVQQKDAtaZWRlZGEgSW5j
LjEcMBoGA1UEAwwTWmVkZWRhIEluYy4gUm9vdCBDQTEeMBwGCSqGSIb3DQEJARYP
Y2VydEB6ZWRlZGEubmV0MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA
4VAlkiCS5Nfq8E1aQwwYTRnFdrd2zuvsQWm3B63BJulTQaBKrc2/5eTcYCN4L2ib
XwPSvMBbTiy6LIDrrcD6FkeUT/kD4Y7mnRktx7dlUX/sUWPwUOsMBpoUdPRPYNY9
lINfsRxJ0DUr7R95t26LqOlGDvqnR6aAkh8Jtc//juzceJdhudh4+xhXOV5+Fhd3
KPvQSwogo0TAs+tbANzegvr7qct1gUcKJQdiHlj1E8rEtAIgDrn1bA8FrTTHYaOS
y8586dJQ2eEv7VdJDHYuX92CKv9SReh1Xpr3sFJ7l9VBJarJ8t2oEidlM/43LWI8
thslrWzbBNF5CQV6cMJv5hhdf5tq3855ER9QGbvBLThyNv4qTf4vJuiT2sziq4zn
1qVdF++uKI6bDLK8C4K1JaNci/GGzKfKDqRNcExA3cWoiqmBc/d37l7qoP+VX1WO
XxmWjdz7e1FPee5R+Y6BOf4LZabPJFLv5bvnJRUkADBPxS/U3cNB3hMsRitglnBa
APg80h5Y9C3BvmLjfwHO0SEYlciVDQNM26LI+ktfHGGy3dphikfZbyz+/wIPvAea
tFAuEAkIpsnYRXG+8q+nPjnH3zVbZfyw9+l8BB3aIQ2b62lge0E7MVG2XCH719Au
Mubt/hkBlrIbZsXppuQ4ysmhSuPWrEOuI+lodv9WLYECAwEAAaNjMGEwHQYDVR0O
BBYEFOudDwZGRQI5ezmKN+0Z2xk96Z/AMB8GA1UdIwQYMBaAFOudDwZGRQI5ezmK
N+0Z2xk96Z/AMA8GA1UdEwEB/wQFMAMBAf8wDgYDVR0PAQH/BAQDAgGGMA0GCSqG
SIb3DQEBCwUAA4ICAQC3JMCuuJBfYR7JZajl5OzZw0G1/1VpI56/oYSkYc+x1R9g
826BIQpt21i4oj8CbvlKRjqtO61i6g09ze/wCpQmWcIBaBpFkPOVgmCnRiJ7dLG3
BpwyMPlMFUt4gNGEbY8mCKkQuwOt+QosvVA+NQu/qDh4cB+NZpG1apySZjAplyyz
VOuSqKh4jZRM8d3Y0cMVcbjseuu2b6F5O3vr+kTYbHvZZ2XOZ4SHswxCQP3jildO
EtXJinqlcKb087F+yKDtlPG/Y4/vXtSIxdwn+Gu5iL/m7mWhcxMuKhOOZ4dvTk+c
Nbi4W4pUtrhsea/+mVnyYlaM/Hwz3CkZiFBUpEknTZbYBPA9mnevSPTYeL+5hwnG
M86tLcQIEPD5raiuN92wBRTLiiktIdS3ZQ40U9ZTp66mveSynd41hTRhL0/Qcejg
wmOOYyUsot0LXwqsH/MeHbrzRpH27rwDNVdePQM2NdcLZ69ItD9oBre23HuHs2vR
wbA/eg6uWlnzr+u5HfLZyyjdKcsj95n+cSSJxqVUeyn8IaitJH5QmQ0YV7rqgA3K
WO6jgPUBJpoh29RRILb41TpFIE3EwJO4oj/JsyhyWB2mBKjGcrz0o6lF86BQ0h/x
bpDyk/81jDcjmDUY584P7Fqa3jhqfmgVFftcCqswQMYAU6xbF1aRWzIVztFFdg==
-----END CERTIFICATE-----
`
)

var (
	DefaultQemuHostFwd  = map[string]string{strconv.Itoa(DefaultSSHPort): "22"}
	DefaultCobraToViper = map[string]string{
		"redis.dist":  "redis-dist",
		"redis.tag":   "redis-tag",
		"redis.port":  "redis-port",
		"redis.force": "redis-force",

		"adam.dist":         "adam-dist",
		"adam.tag":          "adam-tag",
		"adam.port":         "adam-port",
		"adam.domain":       "domain",
		"adam.ip":           "ip",
		"adam.eve-ip":       "eve-ip",
		"adam.force":        "adam-force",
		"adam.v1":           "api-v1",
		"adam.redis.adam":   "adam-redis-url",
		"adam.remote.redis": "adam-redis",

		"eve.arch":         "eve-arch",
		"eve.os":           "eve-os",
		"eve.accel":        "eve-accel",
		"eve.hv":           "hv",
		"eve.serial":       "eve-serial",
		"eve.pid":          "eve-pid",
		"eve.log":          "eve-log",
		"eve.firmware":     "eve-firmware",
		"eve.repo":         "eve-repo",
		"eve.tag":          "eve-tag",
		"eve.hostfwd":      "eve-hostfwd",
		"eve.dist":         "eve-dist",
		"eve.base-dist":    "eve-base-dist",
		"eve.qemu-config":  "qemu-config",
		"eve.uuid":         "uuid",
		"eve.image-file":   "image-file",
		"eve.dtb-part":     "dtb-part",
		"eve.config-part":  "config-part",
		"eve.base-version": "os-version",

		"eden.images.dist":   "image-dist",
		"eden.images.docker": "docker-yml",
		"eden.images.vm":     "vm-yml",
		"eden.download":      "download",
		"eden.eserver.ip":    "eserver-ip",
		"eden.eserver.port":  "eserver-port",
		"eden.eserver.pid":   "eserver-pid",
		"eden.eserver.log":   "eserver-log",
		"eden.certs-dist":    "certs-dist",
		"eden.bin-dist":      "bin-dist",
		"eden.ssh-key":       "ssh-key",
		"eden.test-bin":      "prog",
		"eden.test-scenario": "scenario",
	}
)
