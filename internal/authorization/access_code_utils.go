package authorization

import "io/ioutil"

type ApplicationDataFileUtils interface {
	WriteAccessCodeFile(bytes []byte) error
	WriteClientIdFile(bytes []byte) error
	WriteClientSecretFile(bytes []byte) error
}

type ApplicationDataFileHandler struct{}

func (ApplicationDataFileHandler) WriteAccessCodeFile(bytes []byte) error {
	return writeFile("persistence/access_code", bytes)
}

func (ApplicationDataFileHandler) WriteClientIdFile(bytes []byte) error {
	return writeFile("persistence/client_id", bytes)
}

func (ApplicationDataFileHandler) WriteClientSecretFile(bytes []byte) error {
	return writeFile("persistence/client_secret", bytes)
}

func writeFile(path string, bytes []byte) error {
	return ioutil.WriteFile(path, bytes, 0644)
}
