package authorization

import "io/ioutil"

type AccessCodeFileUtils interface {
	WriteAccessCodeFile(bytes []byte) error
}

type AccessCodeFileHandler struct{}

func (AccessCodeFileHandler) WriteAccessCodeFile(bytes []byte) error {
	return ioutil.WriteFile("persistence/access_code", bytes, 0644)
}
