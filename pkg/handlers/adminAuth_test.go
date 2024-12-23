package handlers

import (
	"context"
	"testing"

	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAdminService is a mock type for the AdminServiceInter interface
type MockAdminService struct {
	mock.Mock
}

// AddMaterialService implements interfaces.AdminServiceInter.
func (m *MockAdminService) AddMaterialService(p *pb.AdminMaterial) (*pb.AdminResponse, error) {
	args := m.Called(p)
	return args.Get(0).(*pb.AdminResponse), args.Error(1)
}

// BlockUserService implements interfaces.AdminServiceInter.
func (m *MockAdminService) BlockUserService(p *pb.AdID) (*pb.AdminResponse, error) {
	args := m.Called(p)
	return args.Get(0).(*pb.AdminResponse), args.Error(1)
}

// EditMaterialService implements interfaces.AdminServiceInter.
func (m *MockAdminService) EditMaterialService(p *pb.AdminMaterial) (*pb.AdminMaterial, error) {
	args := m.Called(p)
	return args.Get(0).(*pb.AdminMaterial), args.Error(1)
}

// FindAllItemService implements interfaces.AdminServiceInter.
func (m *MockAdminService) FindAllItemService(p *pb.AdminItemNoParams) (*pb.AdminItemList, error) {
	args := m.Called(p)
	return args.Get(0).(*pb.AdminItemList), args.Error(1)
}

// FindAllMaterialService implements interfaces.AdminServiceInter.
func (m *MockAdminService) FindAllMaterialService(p *pb.AdminItemNoParams) (*pb.AdminMaterialList, error) {
	args := m.Called(p)
	return args.Get(0).(*pb.AdminMaterialList), args.Error(1)
}

// FindAllOrdersSvc implements interfaces.AdminServiceInter.
func (m *MockAdminService) FindAllOrdersSvc(p *pb.AdminItemNoParams) (*pb.AdminOrderList, error) {
	args := m.Called(p)
	return args.Get(0).(*pb.AdminOrderList), args.Error(1)
}

// FindMaterialByIDService implements interfaces.AdminServiceInter.
func (m *MockAdminService) FindMaterialByIDService(p *pb.AdminMaterialID) (*pb.AdminMaterial, error) {
	args := m.Called(p)
	return args.Get(0).(*pb.AdminMaterial), args.Error(1)
}

// FindOrderSvc implements interfaces.AdminServiceInter.
func (m *MockAdminService) FindOrderSvc(p *pb.AdminItemID) (*pb.AdminOrder, error) {
	args := m.Called(p)
	return args.Get(0).(*pb.AdminOrder), args.Error(1)
}

// FindOrdersByUserSvc implements interfaces.AdminServiceInter.
func (m *MockAdminService) FindOrdersByUserSvc(p *pb.AdminItemID) (*pb.AdminOrderList, error) {
	args := m.Called(p)
	return args.Get(0).(*pb.AdminOrderList), args.Error(1)
}

// RemoveMaterialService implements interfaces.AdminServiceInter.
func (m *MockAdminService) RemoveMaterialService(p *pb.AdminMaterialID) (*pb.AdminResponse, error) {
	args := m.Called(p)
	return args.Get(0).(*pb.AdminResponse), args.Error(1)
}

// UnblockUserService implements interfaces.AdminServiceInter.
func (m *MockAdminService) UnblockUserService(p *pb.AdID) (*pb.AdminResponse, error) {
	args := m.Called(p)
	return args.Get(0).(*pb.AdminResponse), args.Error(1)
}

func (m *MockAdminService) LoginService(p *pb.AdminLogin) (*pb.AdminResponse, error) {
	args := m.Called(p)
	return args.Get(0).(*pb.AdminResponse), args.Error(1)
}

func (m *MockAdminService) AdminViewProfileService(p *pb.AdID) (*pb.AdminProfile, error) {
	args := m.Called(p)
	return args.Get(0).(*pb.AdminProfile), args.Error(1)
}

// TestAdminLoginRequest tests the AdminLoginRequest handler method
func TestAdminLoginRequest(t *testing.T) {
	mockService := new(MockAdminService)

	// Create a mock response
	mockResponse := &pb.AdminResponse{Status: pb.AdminResponse_OK, Message: "Login successful"}

	// Mock the behavior of the LoginService method
	mockService.On("LoginService", mock.Anything).Return(mockResponse, nil)

	handler := NewAdminHandler(mockService)

	// Call the method
	loginRequest := &pb.AdminLogin{Email: "admin@example.com", Password: "password"}
	response, err := handler.AdminLoginRequest(context.Background(), loginRequest)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, pb.AdminResponse_OK, response.Status)
	assert.Equal(t, "Login successful", response.Message)

	// Assert that the mock method was called
	mockService.AssertExpectations(t)
}

// TestAdminViewProfile tests the AdminViewProfile handler method
func TestAdminViewProfile(t *testing.T) {
	mockService := new(MockAdminService)

	// Create a mock response
	mockProfile := &pb.AdminProfile{Admin_ID: 1, Email: "admin@example.com"}

	// Mock the behavior of the AdminViewProfileService method
	mockService.On("AdminViewProfileService", mock.Anything).Return(mockProfile, nil)

	handler := NewAdminHandler(mockService)

	// Call the method
	adminID := &pb.AdID{ID: 1}
	profile, err := handler.AdminViewProfile(context.Background(), adminID)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, profile)
	assert.Equal(t, uint32(1), profile.Admin_ID)
	assert.Equal(t, "admin@example.com", profile.Email)

	// Assert that the mock method was called
	mockService.AssertExpectations(t)
}
