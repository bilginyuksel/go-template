package errors_test

import (
	"gotemplate/pkg/errors"
	"gotemplate/pkg/errors/mock"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestEchoErrorHandler_ErrTypeHTTPError_SendJSONResponseImmediately(t *testing.T) {
	mockContext := mock.NewMockContext(gomock.NewController(t))
	expectedErr := errors.NewHTTPError(http.StatusBadRequest, "made up error")
	mockContext.EXPECT().
		JSON(http.StatusBadRequest, expectedErr).
		Return(nil)

	errors.EchoErrorHandler(nil, nil)(expectedErr, mockContext)
}

func TestEchoErrorHandler_NativeErrType_MappingExists_SendResponse(t *testing.T) {
	actualErr := errors.New("made up error")
	expectedErr := errors.NewHTTPError(http.StatusBadRequest, "made up error")

	mapper := errors.HTTPErrorMapperFunc(func(err error) *errors.HTTPError {
		if errors.Is(err, actualErr) {
			return expectedErr
		}
		return nil
	})

	mockContext := mock.NewMockContext(gomock.NewController(t))
	mockContext.EXPECT().
		JSON(http.StatusBadRequest, expectedErr).
		Return(nil)

	errors.EchoErrorHandler(nil, mapper)(errors.Wrap(actualErr, "wrapped"), mockContext)
}
