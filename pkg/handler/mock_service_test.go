// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/instill-ai/model-backend/pkg/service (interfaces: Service)

// Package handler_test is a generated GoMock package.
package handler_test

import (
	context "context"
	reflect "reflect"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	redis "github.com/go-redis/redis/v9"
	uuid "github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
	resource "github.com/instill-ai/model-backend/internal/resource"
	datamodel "github.com/instill-ai/model-backend/pkg/datamodel"
	repository "github.com/instill-ai/model-backend/pkg/repository"
	service "github.com/instill-ai/model-backend/pkg/service"
	utils "github.com/instill-ai/model-backend/pkg/utils"
	taskv1alpha "github.com/instill-ai/protogen-go/common/task/v1alpha"
	mgmtv1alpha "github.com/instill-ai/protogen-go/core/mgmt/v1alpha"
	modelv1alpha "github.com/instill-ai/protogen-go/model/model/v1alpha"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CheckModel mocks base method.
func (m *MockService) CheckModel(arg0 context.Context, arg1 uuid.UUID) (*modelv1alpha.Model_State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckModel", arg0, arg1)
	ret0, _ := ret[0].(*modelv1alpha.Model_State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckModel indicates an expected call of CheckModel.
func (mr *MockServiceMockRecorder) CheckModel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckModel", reflect.TypeOf((*MockService)(nil).CheckModel), arg0, arg1)
}

// ConvertOwnerNameToPermalink mocks base method.
func (m *MockService) ConvertOwnerNameToPermalink(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertOwnerNameToPermalink", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertOwnerNameToPermalink indicates an expected call of ConvertOwnerNameToPermalink.
func (mr *MockServiceMockRecorder) ConvertOwnerNameToPermalink(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertOwnerNameToPermalink", reflect.TypeOf((*MockService)(nil).ConvertOwnerNameToPermalink), arg0)
}

// ConvertOwnerPermalinkToName mocks base method.
func (m *MockService) ConvertOwnerPermalinkToName(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertOwnerPermalinkToName", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertOwnerPermalinkToName indicates an expected call of ConvertOwnerPermalinkToName.
func (mr *MockServiceMockRecorder) ConvertOwnerPermalinkToName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertOwnerPermalinkToName", reflect.TypeOf((*MockService)(nil).ConvertOwnerPermalinkToName), arg0)
}

// CreateUserModelAsync mocks base method.
func (m *MockService) CreateUserModelAsync(arg0 context.Context, arg1 *datamodel.Model) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserModelAsync", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserModelAsync indicates an expected call of CreateUserModelAsync.
func (mr *MockServiceMockRecorder) CreateUserModelAsync(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserModelAsync", reflect.TypeOf((*MockService)(nil).CreateUserModelAsync), arg0, arg1)
}

// DBToPBModel mocks base method.
func (m *MockService) DBToPBModel(arg0 context.Context, arg1 *datamodel.ModelDefinition, arg2 *datamodel.Model) (*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBToPBModel", arg0, arg1, arg2)
	ret0, _ := ret[0].(*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DBToPBModel indicates an expected call of DBToPBModel.
func (mr *MockServiceMockRecorder) DBToPBModel(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBToPBModel", reflect.TypeOf((*MockService)(nil).DBToPBModel), arg0, arg1, arg2)
}

// DBToPBModelDefinition mocks base method.
func (m *MockService) DBToPBModelDefinition(arg0 context.Context, arg1 *datamodel.ModelDefinition) (*modelv1alpha.ModelDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBToPBModelDefinition", arg0, arg1)
	ret0, _ := ret[0].(*modelv1alpha.ModelDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DBToPBModelDefinition indicates an expected call of DBToPBModelDefinition.
func (mr *MockServiceMockRecorder) DBToPBModelDefinition(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBToPBModelDefinition", reflect.TypeOf((*MockService)(nil).DBToPBModelDefinition), arg0, arg1)
}

// DBToPBModelDefinitions mocks base method.
func (m *MockService) DBToPBModelDefinitions(arg0 context.Context, arg1 []*datamodel.ModelDefinition) ([]*modelv1alpha.ModelDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBToPBModelDefinitions", arg0, arg1)
	ret0, _ := ret[0].([]*modelv1alpha.ModelDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DBToPBModelDefinitions indicates an expected call of DBToPBModelDefinitions.
func (mr *MockServiceMockRecorder) DBToPBModelDefinitions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBToPBModelDefinitions", reflect.TypeOf((*MockService)(nil).DBToPBModelDefinitions), arg0, arg1)
}

// DBToPBModels mocks base method.
func (m *MockService) DBToPBModels(arg0 context.Context, arg1 []*datamodel.Model) ([]*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBToPBModels", arg0, arg1)
	ret0, _ := ret[0].([]*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DBToPBModels indicates an expected call of DBToPBModels.
func (mr *MockServiceMockRecorder) DBToPBModels(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBToPBModels", reflect.TypeOf((*MockService)(nil).DBToPBModels), arg0, arg1)
}

// DeleteResourceState mocks base method.
func (m *MockService) DeleteResourceState(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResourceState", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteResourceState indicates an expected call of DeleteResourceState.
func (mr *MockServiceMockRecorder) DeleteResourceState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResourceState", reflect.TypeOf((*MockService)(nil).DeleteResourceState), arg0, arg1)
}

// DeleteUserModel mocks base method.
func (m *MockService) DeleteUserModel(arg0 context.Context, arg1 resource.Namespace, arg2 uuid.UUID, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserModel", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserModel indicates an expected call of DeleteUserModel.
func (mr *MockServiceMockRecorder) DeleteUserModel(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserModel", reflect.TypeOf((*MockService)(nil).DeleteUserModel), arg0, arg1, arg2, arg3)
}

// DeployUserModelAsyncAdmin mocks base method.
func (m *MockService) DeployUserModelAsyncAdmin(arg0 context.Context, arg1 uuid.UUID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployUserModelAsyncAdmin", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeployUserModelAsyncAdmin indicates an expected call of DeployUserModelAsyncAdmin.
func (mr *MockServiceMockRecorder) DeployUserModelAsyncAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployUserModelAsyncAdmin", reflect.TypeOf((*MockService)(nil).DeployUserModelAsyncAdmin), arg0, arg1)
}

// GetMgmtPrivateServiceClient mocks base method.
func (m *MockService) GetMgmtPrivateServiceClient() mgmtv1alpha.MgmtPrivateServiceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMgmtPrivateServiceClient")
	ret0, _ := ret[0].(mgmtv1alpha.MgmtPrivateServiceClient)
	return ret0
}

// GetMgmtPrivateServiceClient indicates an expected call of GetMgmtPrivateServiceClient.
func (mr *MockServiceMockRecorder) GetMgmtPrivateServiceClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMgmtPrivateServiceClient", reflect.TypeOf((*MockService)(nil).GetMgmtPrivateServiceClient))
}

// GetModelByUID mocks base method.
func (m *MockService) GetModelByUID(arg0 context.Context, arg1, arg2 uuid.UUID, arg3 modelv1alpha.View) (*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelByUID", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelByUID indicates an expected call of GetModelByUID.
func (mr *MockServiceMockRecorder) GetModelByUID(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelByUID", reflect.TypeOf((*MockService)(nil).GetModelByUID), arg0, arg1, arg2, arg3)
}

// GetModelByUIDAdmin mocks base method.
func (m *MockService) GetModelByUIDAdmin(arg0 context.Context, arg1 uuid.UUID, arg2 modelv1alpha.View) (*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelByUIDAdmin", arg0, arg1, arg2)
	ret0, _ := ret[0].(*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelByUIDAdmin indicates an expected call of GetModelByUIDAdmin.
func (mr *MockServiceMockRecorder) GetModelByUIDAdmin(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelByUIDAdmin", reflect.TypeOf((*MockService)(nil).GetModelByUIDAdmin), arg0, arg1, arg2)
}

// GetModelDefinition mocks base method.
func (m *MockService) GetModelDefinition(arg0 context.Context, arg1 string) (*modelv1alpha.ModelDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelDefinition", arg0, arg1)
	ret0, _ := ret[0].(*modelv1alpha.ModelDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelDefinition indicates an expected call of GetModelDefinition.
func (mr *MockServiceMockRecorder) GetModelDefinition(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelDefinition", reflect.TypeOf((*MockService)(nil).GetModelDefinition), arg0, arg1)
}

// GetModelDefinitionByUID mocks base method.
func (m *MockService) GetModelDefinitionByUID(arg0 context.Context, arg1 uuid.UUID) (*modelv1alpha.ModelDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelDefinitionByUID", arg0, arg1)
	ret0, _ := ret[0].(*modelv1alpha.ModelDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelDefinitionByUID indicates an expected call of GetModelDefinitionByUID.
func (mr *MockServiceMockRecorder) GetModelDefinitionByUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelDefinitionByUID", reflect.TypeOf((*MockService)(nil).GetModelDefinitionByUID), arg0, arg1)
}

// GetOperation mocks base method.
func (m *MockService) GetOperation(arg0 context.Context, arg1 string) (*longrunningpb.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperation", arg0, arg1)
	ret0, _ := ret[0].(*longrunningpb.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperation indicates an expected call of GetOperation.
func (mr *MockServiceMockRecorder) GetOperation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockService)(nil).GetOperation), arg0, arg1)
}

// GetRedisClient mocks base method.
func (m *MockService) GetRedisClient() *redis.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRedisClient")
	ret0, _ := ret[0].(*redis.Client)
	return ret0
}

// GetRedisClient indicates an expected call of GetRedisClient.
func (mr *MockServiceMockRecorder) GetRedisClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRedisClient", reflect.TypeOf((*MockService)(nil).GetRedisClient))
}

// GetRepository mocks base method.
func (m *MockService) GetRepository() repository.Repository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepository")
	ret0, _ := ret[0].(repository.Repository)
	return ret0
}

// GetRepository indicates an expected call of GetRepository.
func (mr *MockServiceMockRecorder) GetRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepository", reflect.TypeOf((*MockService)(nil).GetRepository))
}

// GetResourceState mocks base method.
func (m *MockService) GetResourceState(arg0 context.Context, arg1 uuid.UUID) (*modelv1alpha.Model_State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceState", arg0, arg1)
	ret0, _ := ret[0].(*modelv1alpha.Model_State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceState indicates an expected call of GetResourceState.
func (mr *MockServiceMockRecorder) GetResourceState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceState", reflect.TypeOf((*MockService)(nil).GetResourceState), arg0, arg1)
}

// GetRscNamespaceAndNameID mocks base method.
func (m *MockService) GetRscNamespaceAndNameID(arg0 string) (resource.Namespace, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRscNamespaceAndNameID", arg0)
	ret0, _ := ret[0].(resource.Namespace)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRscNamespaceAndNameID indicates an expected call of GetRscNamespaceAndNameID.
func (mr *MockServiceMockRecorder) GetRscNamespaceAndNameID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRscNamespaceAndNameID", reflect.TypeOf((*MockService)(nil).GetRscNamespaceAndNameID), arg0)
}

// GetRscNamespaceAndPermalinkUID mocks base method.
func (m *MockService) GetRscNamespaceAndPermalinkUID(arg0 string) (resource.Namespace, uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRscNamespaceAndPermalinkUID", arg0)
	ret0, _ := ret[0].(resource.Namespace)
	ret1, _ := ret[1].(uuid.UUID)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRscNamespaceAndPermalinkUID indicates an expected call of GetRscNamespaceAndPermalinkUID.
func (mr *MockServiceMockRecorder) GetRscNamespaceAndPermalinkUID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRscNamespaceAndPermalinkUID", reflect.TypeOf((*MockService)(nil).GetRscNamespaceAndPermalinkUID), arg0)
}

// GetTritonEnsembleModel mocks base method.
func (m *MockService) GetTritonEnsembleModel(arg0 context.Context, arg1 uuid.UUID) (*datamodel.TritonModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTritonEnsembleModel", arg0, arg1)
	ret0, _ := ret[0].(*datamodel.TritonModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTritonEnsembleModel indicates an expected call of GetTritonEnsembleModel.
func (mr *MockServiceMockRecorder) GetTritonEnsembleModel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTritonEnsembleModel", reflect.TypeOf((*MockService)(nil).GetTritonEnsembleModel), arg0, arg1)
}

// GetTritonModels mocks base method.
func (m *MockService) GetTritonModels(arg0 context.Context, arg1 uuid.UUID) ([]*datamodel.TritonModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTritonModels", arg0, arg1)
	ret0, _ := ret[0].([]*datamodel.TritonModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTritonModels indicates an expected call of GetTritonModels.
func (mr *MockServiceMockRecorder) GetTritonModels(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTritonModels", reflect.TypeOf((*MockService)(nil).GetTritonModels), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockService) GetUser(arg0 context.Context) (string, uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(uuid.UUID)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUser indicates an expected call of GetUser.
func (mr *MockServiceMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockService)(nil).GetUser), arg0)
}

// GetUserModelByID mocks base method.
func (m *MockService) GetUserModelByID(arg0 context.Context, arg1 resource.Namespace, arg2 uuid.UUID, arg3 string, arg4 modelv1alpha.View) (*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserModelByID", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserModelByID indicates an expected call of GetUserModelByID.
func (mr *MockServiceMockRecorder) GetUserModelByID(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserModelByID", reflect.TypeOf((*MockService)(nil).GetUserModelByID), arg0, arg1, arg2, arg3, arg4)
}

// ListModelDefinitions mocks base method.
func (m *MockService) ListModelDefinitions(arg0 context.Context, arg1 modelv1alpha.View, arg2 int, arg3 string) ([]*modelv1alpha.ModelDefinition, string, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListModelDefinitions", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*modelv1alpha.ModelDefinition)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(int64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ListModelDefinitions indicates an expected call of ListModelDefinitions.
func (mr *MockServiceMockRecorder) ListModelDefinitions(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModelDefinitions", reflect.TypeOf((*MockService)(nil).ListModelDefinitions), arg0, arg1, arg2, arg3)
}

// ListModels mocks base method.
func (m *MockService) ListModels(arg0 context.Context, arg1 uuid.UUID, arg2 modelv1alpha.View, arg3 int, arg4 string, arg5 bool) ([]*modelv1alpha.Model, string, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListModels", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]*modelv1alpha.Model)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(int64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ListModels indicates an expected call of ListModels.
func (mr *MockServiceMockRecorder) ListModels(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModels", reflect.TypeOf((*MockService)(nil).ListModels), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ListModelsAdmin mocks base method.
func (m *MockService) ListModelsAdmin(arg0 context.Context, arg1 modelv1alpha.View, arg2 int, arg3 string, arg4 bool) ([]*modelv1alpha.Model, string, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListModelsAdmin", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*modelv1alpha.Model)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(int64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ListModelsAdmin indicates an expected call of ListModelsAdmin.
func (mr *MockServiceMockRecorder) ListModelsAdmin(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModelsAdmin", reflect.TypeOf((*MockService)(nil).ListModelsAdmin), arg0, arg1, arg2, arg3, arg4)
}

// ListUserModels mocks base method.
func (m *MockService) ListUserModels(arg0 context.Context, arg1 resource.Namespace, arg2 uuid.UUID, arg3 modelv1alpha.View, arg4 int, arg5 string, arg6 bool) ([]*modelv1alpha.Model, string, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserModels", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].([]*modelv1alpha.Model)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(int64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ListUserModels indicates an expected call of ListUserModels.
func (mr *MockServiceMockRecorder) ListUserModels(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserModels", reflect.TypeOf((*MockService)(nil).ListUserModels), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// PBToDBModel mocks base method.
func (m *MockService) PBToDBModel(arg0 context.Context, arg1 *modelv1alpha.Model) *datamodel.Model {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PBToDBModel", arg0, arg1)
	ret0, _ := ret[0].(*datamodel.Model)
	return ret0
}

// PBToDBModel indicates an expected call of PBToDBModel.
func (mr *MockServiceMockRecorder) PBToDBModel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PBToDBModel", reflect.TypeOf((*MockService)(nil).PBToDBModel), arg0, arg1)
}

// PublishUserModel mocks base method.
func (m *MockService) PublishUserModel(arg0 context.Context, arg1 resource.Namespace, arg2 uuid.UUID, arg3 string) (*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishUserModel", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishUserModel indicates an expected call of PublishUserModel.
func (mr *MockServiceMockRecorder) PublishUserModel(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishUserModel", reflect.TypeOf((*MockService)(nil).PublishUserModel), arg0, arg1, arg2, arg3)
}

// RenameUserModel mocks base method.
func (m *MockService) RenameUserModel(arg0 context.Context, arg1 resource.Namespace, arg2 uuid.UUID, arg3, arg4 string) (*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameUserModel", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenameUserModel indicates an expected call of RenameUserModel.
func (mr *MockServiceMockRecorder) RenameUserModel(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameUserModel", reflect.TypeOf((*MockService)(nil).RenameUserModel), arg0, arg1, arg2, arg3, arg4)
}

// TriggerUserModel mocks base method.
func (m *MockService) TriggerUserModel(arg0 context.Context, arg1 uuid.UUID, arg2 service.InferInput, arg3 taskv1alpha.Task) ([]*modelv1alpha.TaskOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerUserModel", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*modelv1alpha.TaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerUserModel indicates an expected call of TriggerUserModel.
func (mr *MockServiceMockRecorder) TriggerUserModel(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerUserModel", reflect.TypeOf((*MockService)(nil).TriggerUserModel), arg0, arg1, arg2, arg3)
}

// TriggerUserModelTestMode mocks base method.
func (m *MockService) TriggerUserModelTestMode(arg0 context.Context, arg1 uuid.UUID, arg2 service.InferInput, arg3 taskv1alpha.Task) ([]*modelv1alpha.TaskOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerUserModelTestMode", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*modelv1alpha.TaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerUserModelTestMode indicates an expected call of TriggerUserModelTestMode.
func (mr *MockServiceMockRecorder) TriggerUserModelTestMode(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerUserModelTestMode", reflect.TypeOf((*MockService)(nil).TriggerUserModelTestMode), arg0, arg1, arg2, arg3)
}

// UndeployUserModelAsyncAdmin mocks base method.
func (m *MockService) UndeployUserModelAsyncAdmin(arg0 context.Context, arg1, arg2 uuid.UUID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UndeployUserModelAsyncAdmin", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UndeployUserModelAsyncAdmin indicates an expected call of UndeployUserModelAsyncAdmin.
func (mr *MockServiceMockRecorder) UndeployUserModelAsyncAdmin(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UndeployUserModelAsyncAdmin", reflect.TypeOf((*MockService)(nil).UndeployUserModelAsyncAdmin), arg0, arg1, arg2)
}

// UnpublishUserModel mocks base method.
func (m *MockService) UnpublishUserModel(arg0 context.Context, arg1 resource.Namespace, arg2 uuid.UUID, arg3 string) (*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpublishUserModel", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnpublishUserModel indicates an expected call of UnpublishUserModel.
func (mr *MockServiceMockRecorder) UnpublishUserModel(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpublishUserModel", reflect.TypeOf((*MockService)(nil).UnpublishUserModel), arg0, arg1, arg2, arg3)
}

// UpdateResourceState mocks base method.
func (m *MockService) UpdateResourceState(arg0 context.Context, arg1 uuid.UUID, arg2 modelv1alpha.Model_State, arg3 *int32, arg4 *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResourceState", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateResourceState indicates an expected call of UpdateResourceState.
func (mr *MockServiceMockRecorder) UpdateResourceState(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResourceState", reflect.TypeOf((*MockService)(nil).UpdateResourceState), arg0, arg1, arg2, arg3, arg4)
}

// UpdateUserModel mocks base method.
func (m *MockService) UpdateUserModel(arg0 context.Context, arg1 resource.Namespace, arg2 uuid.UUID, arg3 *modelv1alpha.Model) (*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserModel", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserModel indicates an expected call of UpdateUserModel.
func (mr *MockServiceMockRecorder) UpdateUserModel(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserModel", reflect.TypeOf((*MockService)(nil).UpdateUserModel), arg0, arg1, arg2, arg3)
}

// UpdateUserModelState mocks base method.
func (m *MockService) UpdateUserModelState(arg0 context.Context, arg1 resource.Namespace, arg2 uuid.UUID, arg3 *modelv1alpha.Model, arg4 modelv1alpha.Model_State) (*modelv1alpha.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserModelState", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*modelv1alpha.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserModelState indicates an expected call of UpdateUserModelState.
func (mr *MockServiceMockRecorder) UpdateUserModelState(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserModelState", reflect.TypeOf((*MockService)(nil).UpdateUserModelState), arg0, arg1, arg2, arg3, arg4)
}

// WriteNewDataPoint mocks base method.
func (m *MockService) WriteNewDataPoint(arg0 context.Context, arg1 utils.UsageMetricData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteNewDataPoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteNewDataPoint indicates an expected call of WriteNewDataPoint.
func (mr *MockServiceMockRecorder) WriteNewDataPoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteNewDataPoint", reflect.TypeOf((*MockService)(nil).WriteNewDataPoint), arg0, arg1)
}
