// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/labstack/echo/v4 (interfaces: Context)

// Package mock is a generated GoMock package.
package mock

import (
	io "io"
	multipart "mime/multipart"
	http "net/http"
	url "net/url"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
)

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// Attachment mocks base method.
func (m *MockContext) Attachment(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attachment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Attachment indicates an expected call of Attachment.
func (mr *MockContextMockRecorder) Attachment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attachment", reflect.TypeOf((*MockContext)(nil).Attachment), arg0, arg1)
}

// Bind mocks base method.
func (m *MockContext) Bind(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bind", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bind indicates an expected call of Bind.
func (mr *MockContextMockRecorder) Bind(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*MockContext)(nil).Bind), arg0)
}

// Blob mocks base method.
func (m *MockContext) Blob(arg0 int, arg1 string, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Blob", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Blob indicates an expected call of Blob.
func (mr *MockContextMockRecorder) Blob(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Blob", reflect.TypeOf((*MockContext)(nil).Blob), arg0, arg1, arg2)
}

// Cookie mocks base method.
func (m *MockContext) Cookie(arg0 string) (*http.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cookie", arg0)
	ret0, _ := ret[0].(*http.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cookie indicates an expected call of Cookie.
func (mr *MockContextMockRecorder) Cookie(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cookie", reflect.TypeOf((*MockContext)(nil).Cookie), arg0)
}

// Cookies mocks base method.
func (m *MockContext) Cookies() []*http.Cookie {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cookies")
	ret0, _ := ret[0].([]*http.Cookie)
	return ret0
}

// Cookies indicates an expected call of Cookies.
func (mr *MockContextMockRecorder) Cookies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cookies", reflect.TypeOf((*MockContext)(nil).Cookies))
}

// Echo mocks base method.
func (m *MockContext) Echo() *echo.Echo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Echo")
	ret0, _ := ret[0].(*echo.Echo)
	return ret0
}

// Echo indicates an expected call of Echo.
func (mr *MockContextMockRecorder) Echo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Echo", reflect.TypeOf((*MockContext)(nil).Echo))
}

// Error mocks base method.
func (m *MockContext) Error(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", arg0)
}

// Error indicates an expected call of Error.
func (mr *MockContextMockRecorder) Error(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockContext)(nil).Error), arg0)
}

// File mocks base method.
func (m *MockContext) File(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "File", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// File indicates an expected call of File.
func (mr *MockContextMockRecorder) File(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "File", reflect.TypeOf((*MockContext)(nil).File), arg0)
}

// FormFile mocks base method.
func (m *MockContext) FormFile(arg0 string) (*multipart.FileHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormFile", arg0)
	ret0, _ := ret[0].(*multipart.FileHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormFile indicates an expected call of FormFile.
func (mr *MockContextMockRecorder) FormFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormFile", reflect.TypeOf((*MockContext)(nil).FormFile), arg0)
}

// FormParams mocks base method.
func (m *MockContext) FormParams() (url.Values, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormParams")
	ret0, _ := ret[0].(url.Values)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormParams indicates an expected call of FormParams.
func (mr *MockContextMockRecorder) FormParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormParams", reflect.TypeOf((*MockContext)(nil).FormParams))
}

// FormValue mocks base method.
func (m *MockContext) FormValue(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormValue", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// FormValue indicates an expected call of FormValue.
func (mr *MockContextMockRecorder) FormValue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormValue", reflect.TypeOf((*MockContext)(nil).FormValue), arg0)
}

// Get mocks base method.
func (m *MockContext) Get(arg0 string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockContextMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockContext)(nil).Get), arg0)
}

// HTML mocks base method.
func (m *MockContext) HTML(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTML", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HTML indicates an expected call of HTML.
func (mr *MockContextMockRecorder) HTML(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTML", reflect.TypeOf((*MockContext)(nil).HTML), arg0, arg1)
}

// HTMLBlob mocks base method.
func (m *MockContext) HTMLBlob(arg0 int, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTMLBlob", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HTMLBlob indicates an expected call of HTMLBlob.
func (mr *MockContextMockRecorder) HTMLBlob(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTMLBlob", reflect.TypeOf((*MockContext)(nil).HTMLBlob), arg0, arg1)
}

// Handler mocks base method.
func (m *MockContext) Handler() echo.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handler")
	ret0, _ := ret[0].(echo.HandlerFunc)
	return ret0
}

// Handler indicates an expected call of Handler.
func (mr *MockContextMockRecorder) Handler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handler", reflect.TypeOf((*MockContext)(nil).Handler))
}

// Inline mocks base method.
func (m *MockContext) Inline(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Inline", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Inline indicates an expected call of Inline.
func (mr *MockContextMockRecorder) Inline(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inline", reflect.TypeOf((*MockContext)(nil).Inline), arg0, arg1)
}

// IsTLS mocks base method.
func (m *MockContext) IsTLS() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTLS")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTLS indicates an expected call of IsTLS.
func (mr *MockContextMockRecorder) IsTLS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTLS", reflect.TypeOf((*MockContext)(nil).IsTLS))
}

// IsWebSocket mocks base method.
func (m *MockContext) IsWebSocket() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsWebSocket")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsWebSocket indicates an expected call of IsWebSocket.
func (mr *MockContextMockRecorder) IsWebSocket() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsWebSocket", reflect.TypeOf((*MockContext)(nil).IsWebSocket))
}

// JSON mocks base method.
func (m *MockContext) JSON(arg0 int, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSON", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// JSON indicates an expected call of JSON.
func (mr *MockContextMockRecorder) JSON(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSON", reflect.TypeOf((*MockContext)(nil).JSON), arg0, arg1)
}

// JSONBlob mocks base method.
func (m *MockContext) JSONBlob(arg0 int, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSONBlob", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// JSONBlob indicates an expected call of JSONBlob.
func (mr *MockContextMockRecorder) JSONBlob(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSONBlob", reflect.TypeOf((*MockContext)(nil).JSONBlob), arg0, arg1)
}

// JSONP mocks base method.
func (m *MockContext) JSONP(arg0 int, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSONP", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// JSONP indicates an expected call of JSONP.
func (mr *MockContextMockRecorder) JSONP(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSONP", reflect.TypeOf((*MockContext)(nil).JSONP), arg0, arg1, arg2)
}

// JSONPBlob mocks base method.
func (m *MockContext) JSONPBlob(arg0 int, arg1 string, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSONPBlob", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// JSONPBlob indicates an expected call of JSONPBlob.
func (mr *MockContextMockRecorder) JSONPBlob(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSONPBlob", reflect.TypeOf((*MockContext)(nil).JSONPBlob), arg0, arg1, arg2)
}

// JSONPretty mocks base method.
func (m *MockContext) JSONPretty(arg0 int, arg1 interface{}, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSONPretty", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// JSONPretty indicates an expected call of JSONPretty.
func (mr *MockContextMockRecorder) JSONPretty(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSONPretty", reflect.TypeOf((*MockContext)(nil).JSONPretty), arg0, arg1, arg2)
}

// Logger mocks base method.
func (m *MockContext) Logger() echo.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(echo.Logger)
	return ret0
}

// Logger indicates an expected call of Logger.
func (mr *MockContextMockRecorder) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockContext)(nil).Logger))
}

// MultipartForm mocks base method.
func (m *MockContext) MultipartForm() (*multipart.Form, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultipartForm")
	ret0, _ := ret[0].(*multipart.Form)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultipartForm indicates an expected call of MultipartForm.
func (mr *MockContextMockRecorder) MultipartForm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultipartForm", reflect.TypeOf((*MockContext)(nil).MultipartForm))
}

// NoContent mocks base method.
func (m *MockContext) NoContent(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NoContent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// NoContent indicates an expected call of NoContent.
func (mr *MockContextMockRecorder) NoContent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoContent", reflect.TypeOf((*MockContext)(nil).NoContent), arg0)
}

// Param mocks base method.
func (m *MockContext) Param(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Param", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Param indicates an expected call of Param.
func (mr *MockContextMockRecorder) Param(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Param", reflect.TypeOf((*MockContext)(nil).Param), arg0)
}

// ParamNames mocks base method.
func (m *MockContext) ParamNames() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParamNames")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ParamNames indicates an expected call of ParamNames.
func (mr *MockContextMockRecorder) ParamNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParamNames", reflect.TypeOf((*MockContext)(nil).ParamNames))
}

// ParamValues mocks base method.
func (m *MockContext) ParamValues() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParamValues")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ParamValues indicates an expected call of ParamValues.
func (mr *MockContextMockRecorder) ParamValues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParamValues", reflect.TypeOf((*MockContext)(nil).ParamValues))
}

// Path mocks base method.
func (m *MockContext) Path() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

// Path indicates an expected call of Path.
func (mr *MockContextMockRecorder) Path() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockContext)(nil).Path))
}

// QueryParam mocks base method.
func (m *MockContext) QueryParam(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryParam", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// QueryParam indicates an expected call of QueryParam.
func (mr *MockContextMockRecorder) QueryParam(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryParam", reflect.TypeOf((*MockContext)(nil).QueryParam), arg0)
}

// QueryParams mocks base method.
func (m *MockContext) QueryParams() url.Values {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryParams")
	ret0, _ := ret[0].(url.Values)
	return ret0
}

// QueryParams indicates an expected call of QueryParams.
func (mr *MockContextMockRecorder) QueryParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryParams", reflect.TypeOf((*MockContext)(nil).QueryParams))
}

// QueryString mocks base method.
func (m *MockContext) QueryString() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryString")
	ret0, _ := ret[0].(string)
	return ret0
}

// QueryString indicates an expected call of QueryString.
func (mr *MockContextMockRecorder) QueryString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryString", reflect.TypeOf((*MockContext)(nil).QueryString))
}

// RealIP mocks base method.
func (m *MockContext) RealIP() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RealIP")
	ret0, _ := ret[0].(string)
	return ret0
}

// RealIP indicates an expected call of RealIP.
func (mr *MockContextMockRecorder) RealIP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RealIP", reflect.TypeOf((*MockContext)(nil).RealIP))
}

// Redirect mocks base method.
func (m *MockContext) Redirect(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Redirect", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Redirect indicates an expected call of Redirect.
func (mr *MockContextMockRecorder) Redirect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Redirect", reflect.TypeOf((*MockContext)(nil).Redirect), arg0, arg1)
}

// Render mocks base method.
func (m *MockContext) Render(arg0 int, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Render", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Render indicates an expected call of Render.
func (mr *MockContextMockRecorder) Render(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Render", reflect.TypeOf((*MockContext)(nil).Render), arg0, arg1, arg2)
}

// Request mocks base method.
func (m *MockContext) Request() *http.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request")
	ret0, _ := ret[0].(*http.Request)
	return ret0
}

// Request indicates an expected call of Request.
func (mr *MockContextMockRecorder) Request() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockContext)(nil).Request))
}

// Reset mocks base method.
func (m *MockContext) Reset(arg0 *http.Request, arg1 http.ResponseWriter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", arg0, arg1)
}

// Reset indicates an expected call of Reset.
func (mr *MockContextMockRecorder) Reset(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockContext)(nil).Reset), arg0, arg1)
}

// Response mocks base method.
func (m *MockContext) Response() *echo.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Response")
	ret0, _ := ret[0].(*echo.Response)
	return ret0
}

// Response indicates an expected call of Response.
func (mr *MockContextMockRecorder) Response() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Response", reflect.TypeOf((*MockContext)(nil).Response))
}

// Scheme mocks base method.
func (m *MockContext) Scheme() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scheme")
	ret0, _ := ret[0].(string)
	return ret0
}

// Scheme indicates an expected call of Scheme.
func (mr *MockContextMockRecorder) Scheme() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scheme", reflect.TypeOf((*MockContext)(nil).Scheme))
}

// Set mocks base method.
func (m *MockContext) Set(arg0 string, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0, arg1)
}

// Set indicates an expected call of Set.
func (mr *MockContextMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockContext)(nil).Set), arg0, arg1)
}

// SetCookie mocks base method.
func (m *MockContext) SetCookie(arg0 *http.Cookie) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCookie", arg0)
}

// SetCookie indicates an expected call of SetCookie.
func (mr *MockContextMockRecorder) SetCookie(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCookie", reflect.TypeOf((*MockContext)(nil).SetCookie), arg0)
}

// SetHandler mocks base method.
func (m *MockContext) SetHandler(arg0 echo.HandlerFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHandler", arg0)
}

// SetHandler indicates an expected call of SetHandler.
func (mr *MockContextMockRecorder) SetHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHandler", reflect.TypeOf((*MockContext)(nil).SetHandler), arg0)
}

// SetLogger mocks base method.
func (m *MockContext) SetLogger(arg0 echo.Logger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogger", arg0)
}

// SetLogger indicates an expected call of SetLogger.
func (mr *MockContextMockRecorder) SetLogger(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogger", reflect.TypeOf((*MockContext)(nil).SetLogger), arg0)
}

// SetParamNames mocks base method.
func (m *MockContext) SetParamNames(arg0 ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "SetParamNames", varargs...)
}

// SetParamNames indicates an expected call of SetParamNames.
func (mr *MockContextMockRecorder) SetParamNames(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParamNames", reflect.TypeOf((*MockContext)(nil).SetParamNames), arg0...)
}

// SetParamValues mocks base method.
func (m *MockContext) SetParamValues(arg0 ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "SetParamValues", varargs...)
}

// SetParamValues indicates an expected call of SetParamValues.
func (mr *MockContextMockRecorder) SetParamValues(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParamValues", reflect.TypeOf((*MockContext)(nil).SetParamValues), arg0...)
}

// SetPath mocks base method.
func (m *MockContext) SetPath(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPath", arg0)
}

// SetPath indicates an expected call of SetPath.
func (mr *MockContextMockRecorder) SetPath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPath", reflect.TypeOf((*MockContext)(nil).SetPath), arg0)
}

// SetRequest mocks base method.
func (m *MockContext) SetRequest(arg0 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRequest", arg0)
}

// SetRequest indicates an expected call of SetRequest.
func (mr *MockContextMockRecorder) SetRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRequest", reflect.TypeOf((*MockContext)(nil).SetRequest), arg0)
}

// SetResponse mocks base method.
func (m *MockContext) SetResponse(arg0 *echo.Response) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetResponse", arg0)
}

// SetResponse indicates an expected call of SetResponse.
func (mr *MockContextMockRecorder) SetResponse(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResponse", reflect.TypeOf((*MockContext)(nil).SetResponse), arg0)
}

// Stream mocks base method.
func (m *MockContext) Stream(arg0 int, arg1 string, arg2 io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stream", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stream indicates an expected call of Stream.
func (mr *MockContextMockRecorder) Stream(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockContext)(nil).Stream), arg0, arg1, arg2)
}

// String mocks base method.
func (m *MockContext) String(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockContextMockRecorder) String(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockContext)(nil).String), arg0, arg1)
}

// Validate mocks base method.
func (m *MockContext) Validate(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockContextMockRecorder) Validate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockContext)(nil).Validate), arg0)
}

// XML mocks base method.
func (m *MockContext) XML(arg0 int, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XML", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// XML indicates an expected call of XML.
func (mr *MockContextMockRecorder) XML(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XML", reflect.TypeOf((*MockContext)(nil).XML), arg0, arg1)
}

// XMLBlob mocks base method.
func (m *MockContext) XMLBlob(arg0 int, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XMLBlob", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// XMLBlob indicates an expected call of XMLBlob.
func (mr *MockContextMockRecorder) XMLBlob(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XMLBlob", reflect.TypeOf((*MockContext)(nil).XMLBlob), arg0, arg1)
}

// XMLPretty mocks base method.
func (m *MockContext) XMLPretty(arg0 int, arg1 interface{}, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XMLPretty", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// XMLPretty indicates an expected call of XMLPretty.
func (mr *MockContextMockRecorder) XMLPretty(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XMLPretty", reflect.TypeOf((*MockContext)(nil).XMLPretty), arg0, arg1, arg2)
}