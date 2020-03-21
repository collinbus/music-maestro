package authorization

import "testing"

func TestSuccessSavingAccessCodeWhenAbcIsPassed(t *testing.T) {
	service := NewAccessCodeService(MockFileHandler{})

	success := service.SaveAccessCode("abc")

	if !success {
		t.Error("Saving of the access code failed.")
	}
}

func TestSuccessSavingClientIdWhenAbcIsPassed(t *testing.T) {
	service := NewAccessCodeService(MockFileHandler{})

	success := service.SaveClientId("abc")

	if !success {
		t.Error("Saving of the client id failed.")
	}
}

func TestSuccessSavingClientSecretWhenAbcIsPassed(t *testing.T) {
	service := NewAccessCodeService(MockFileHandler{})

	success := service.SaveClientSecret("abc")

	if !success {
		t.Error("Saving of the client secret failed.")
	}
}

type MockFileHandler struct{}

func (mock MockFileHandler) WriteAccessCodeFile(_ []byte) error {
	return nil
}

func (mock MockFileHandler) WriteClientIdFile(bytes []byte) error {
	return nil
}

func (mock MockFileHandler) WriteClientSecretFile(bytes []byte) error {
	return nil
}
