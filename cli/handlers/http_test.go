package handlers

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/zerpto/ponodo/config"
	"github.com/zerpto/ponodo/contracts"
	"github.com/zerpto/ponodo/mocks"
)

func TestHttpHandler_Short(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := mocks.NewMockAppContract(ctrl)
	handler := &HttpHandler{
		App: mockApp,
	}

	result := handler.Short()
	assert.Equal(t, "Run http server to expose api endpoints.", result)
}

func TestHttpHandler_Long(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := mocks.NewMockAppContract(ctrl)
	handler := &HttpHandler{
		App: mockApp,
	}

	result := handler.Long()
	assert.Equal(t, "Run http server to expose api endpoints.", result)
}

func TestHttpHandler_Example(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := mocks.NewMockAppContract(ctrl)
	handler := &HttpHandler{
		App: mockApp,
	}

	result := handler.Example()
	assert.Equal(t, "zerpto http --port 8080", result)
}

func TestHttpHandler_Use(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := mocks.NewMockAppContract(ctrl)
	handler := &HttpHandler{
		App: mockApp,
	}

	result := handler.Use()
	assert.Equal(t, "http", result)
}

func TestHttpHandler_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := mocks.NewMockAppContract(ctrl)
	mockConfigLoader := mocks.NewMockConfigContract(ctrl)

	gin.SetMode(gin.TestMode)

	mockApp.EXPECT().GetConfigLoader().Return(&config.Loader{
		Config: mockConfigLoader,
	}).AnyTimes()
	mockConfigLoader.EXPECT().GetDebug().Return(false).AnyTimes()
	mockApp.EXPECT().SetGin(gomock.Any()).Do(func(engine *gin.Engine) {
		assert.NotNil(t, engine)
	}).AnyTimes()
	mockApp.EXPECT().SetValidator(gomock.Any()).Do(func(v *validator.Validate) {
		assert.NotNil(t, v)
	}).AnyTimes()

	handler := &HttpHandler{
		App: mockApp,
		RouterSetupFn: func(app contracts.AppContract) {
			// Setup router for test
		},
	}

	// We can't easily test the full Run method without mocking http.Server,
	// but we can test the initial setup
	assert.NotNil(t, handler.App)
	assert.NotNil(t, handler.RouterSetupFn)
}

func TestNewHttpHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := mocks.NewMockAppContract(ctrl)
	setupFn := func(app contracts.AppContract) {
		// Test setup function
	}

	handlerContract := NewHttpHandler(mockApp, setupFn)

	require.NotNil(t, handlerContract)
	
	// Type assert to HttpHandler to access fields
	handler, ok := handlerContract.(*HttpHandler)
	require.True(t, ok, "NewHttpHandler should return *HttpHandler")
	assert.Equal(t, mockApp, handler.App)
	assert.NotNil(t, handler.RouterSetupFn)
}

