package authorization

import (
	"errors"
	"testing"
)

func TestSuccessSavingAccessCodeWhenAbcIsPassed(t *testing.T) {
	service := NewAccessCodeService(MockSuccessFileHandler{})

	success := service.SaveAccessCode("abc")

	if !success {
		t.Error("Saving of the access code failed.")
	}
}

func TestFailureSavingAccessCodeWhenAbcIsPassed(t *testing.T) {
	service := NewAccessCodeService(MockFailureFileHandler{})

	success := service.SaveAccessCode("abc")

	if success {
		t.Error("Saving of the access code should fail.")
	}
}

func TestSuccessSavingClientIdWhenAbcIsPassed(t *testing.T) {
	service := NewAccessCodeService(MockSuccessFileHandler{})

	success := service.SaveClientId("abc")

	if !success {
		t.Error("Saving of the client id failed.")
	}
}

func TestFailureSavingClientIdWhenAbcIsPassed(t *testing.T) {
	service := NewAccessCodeService(MockFailureFileHandler{})

	success := service.SaveAccessCode("abc")

	if success {
		t.Error("Saving of the access code should fail.")
	}
}

func TestSuccessSavingClientSecretWhenAbcIsPassed(t *testing.T) {
	service := NewAccessCodeService(MockSuccessFileHandler{})

	success := service.SaveClientSecret("abc")

	if !success {
		t.Error("Saving of the client secret failed.")
	}
}

func TestFailureSavingClientSecretWhenAbcIsPassed(t *testing.T) {
	service := NewAccessCodeService(MockFailureFileHandler{})

	success := service.SaveAccessCode("abc")

	if success {
		t.Error("Saving of the access code should fail.")
	}
}

type MockSuccessFileHandler struct{}

func (mock MockSuccessFileHandler) WriteAccessCodeFile(_ []byte) error {
	return nil
}

func (mock MockSuccessFileHandler) WriteClientIdFile(bytes []byte) error {
	return nil
}

func (mock MockSuccessFileHandler) WriteClientSecretFile(bytes []byte) error {
	return nil
}

type MockFailureFileHandler struct{}

func (mock MockFailureFileHandler) WriteAccessCodeFile(_ []byte) error {
	return errors.New("saving access code failed")
}

func (mock MockFailureFileHandler) WriteClientIdFile(bytes []byte) error {
	return errors.New("saving client id failed")
}

func (mock MockFailureFileHandler) WriteClientSecretFile(bytes []byte) error {
	return errors.New("saving client secret failed")
}
