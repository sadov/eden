package expect

import (
	"github.com/dustin/go-humanize"
	"github.com/lf-edge/eden/eserver/api"
	"github.com/lf-edge/eden/pkg/eden"
	"github.com/lf-edge/eve/api/go/config"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

//createImageFile uploads image into EServer from file and calculates size and sha256 of image
func (exp *AppExpectation) createImageFile(id uuid.UUID, dsID string) *config.Image {
	server := &eden.EServer{
		EServerIP:   exp.ctrl.GetVars().EServerIP,
		EServerPort: exp.ctrl.GetVars().EServerPort,
	}
	var fileSize int64
	var status *api.FileInfo
	var err error
	sha256 := ""
	filePath := ""
	if status, err = eden.AddFileIntoEServer(server, exp.appURL); err != nil {
		log.Error(err)
	}
	sha256 = status.Sha256
	fileSize = status.Size
	filePath = status.FileName
	log.Infof("Image uploaded with size %s and sha256 %s", humanize.Bytes(uint64(status.Size)), status.Sha256)
	if filePath == "" {
		log.Fatal("Not uploaded")
	}
	return &config.Image{
		Uuidandversion: &config.UUIDandVersion{
			Uuid:    id.String(),
			Version: "1",
		},
		Name:      filePath,
		Iformat:   exp.imageFormatEnum(),
		DsId:      dsID,
		SizeBytes: fileSize,
		Sha256:    sha256,
	}
}
