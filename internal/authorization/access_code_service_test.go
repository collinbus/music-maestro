package authorization

import "testing"

type MockFileHandler struct{}

func TestFailureSavingAccessCodeWhenAbcIsPassed(t *testing.T) {
	service := NewAccessCodeService(MockFileHandler{})

	success := service.Save("abc")

	if !success {
		t.Error("Saving of the access code failed.")
	}
}

func (mock MockFileHandler) WriteAccessCodeFile(_ []byte) error {
	return nil
}
