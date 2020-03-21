package authorization

import (
	"bytes"
)

type ApplicationDataService struct {
	accessCodeFileUtils ApplicationDataFileUtils
}

func (handler ApplicationDataService) SaveAccessCode(accessCode string) bool {
	buffer := bytes.NewBufferString(accessCode)
	err := handler.accessCodeFileUtils.WriteAccessCodeFile(buffer.Bytes())
	if err != nil {
		return false
	}
	return true
}

func (handler ApplicationDataService) SaveClientId(id string) bool {
	buffer := bytes.NewBufferString(id)
	err := handler.accessCodeFileUtils.WriteClientIdFile(buffer.Bytes())
	if err != nil {
		return false
	}
	return true
}

func (handler ApplicationDataService) SaveClientSecret(secret string) bool {
	buffer := bytes.NewBufferString(secret)
	err := handler.accessCodeFileUtils.WriteClientSecretFile(buffer.Bytes())
	if err != nil {
		return false
	}
	return true
}

func NewAccessCodeService(utils ApplicationDataFileUtils) *ApplicationDataService {
	return &ApplicationDataService{
		accessCodeFileUtils: utils,
	}
}
