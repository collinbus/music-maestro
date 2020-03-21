package authorization

import (
	"bytes"
	"log"
)

type AccessCodeService struct {
	accessCodeFileUtils AccessCodeFileUtils
}

func (handler AccessCodeService) Save(accessCode string) bool {
	buffer := bytes.NewBufferString(accessCode)
	err := handler.accessCodeFileUtils.WriteAccessCodeFile(buffer.Bytes())
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func NewAccessCodeService(utils AccessCodeFileUtils) *AccessCodeService {
	return &AccessCodeService{
		accessCodeFileUtils: utils,
	}
}
