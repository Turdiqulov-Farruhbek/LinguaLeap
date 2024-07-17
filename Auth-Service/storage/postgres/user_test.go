package postgres

// import (
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	pb "gitlab.com/unknownteam/sub-auth-service/protos/auth"
// )

// func TestRegister(t *testing.T) {
// 	db, err := NewPostgresStorage()
// 	if err != nil {
// 		panic(err)
// 	}
// 	user := &pb.RegisterRequest{
// 		Username: uuid.NewString(),
// 		Email:    uuid.NewString(),
// 		Password: uuid.NewString(),
// 	}
// 	_, err = db.User().Register(user)

// 	assert.NoError(t, err)
// 	//assert.NotNil(t, resp)
// }

// func TestLogin(t *testing.T) {
// 	db, err := NewPostgresStorage()
// 	if err != nil {
// 		panic(err)
// 	}
// 	// Register a user first
// 	user := &pb.RegisterRequest{
// 		Username: "test_user",
// 		Email:    "test_user@example.com",
// 		Password: "test_password",
// 	}
// 	_, err = db.User().Register(user)
// 	assert.NoError(t, err)

// 	// Try to login
// 	loginReq := &pb.LoginRequest{
// 		Username: "test_user",
// 		Password: "test_password",
// 	}
// 	loginResp, err := db.User().Login(loginReq)

// 	assert.NoError(t, err)
// 	assert.True(t, loginResp.Success)
// }

// func TestRefreshToken(t *testing.T) {
// 	db, err := NewPostgresStorage()
// 	if err != nil {
// 		panic(err)
// 	}
// 	// Assume token is already stored in refresh_tokens table
// 	token := &pb.RefreshTokenRequest{
// 		Token: "your_existing_refresh_token",
// 	}
// 	refreshResp, err := db.User().RefreshToken(token)

// 	assert.NoError(t, err)
// 	assert.True(t, refreshResp.Success)
// 	assert.NotEmpty(t, refreshResp.NewRefreshToken)
// }

// func TestLogout(t *testing.T) {
// 	db, err := NewPostgresStorage()

// 	if err != nil {
// 		panic(err)
// 	}
// 	// Assume token is already stored in refresh_tokens table
// 	logoutReq := &pb.LogoutRequest{
// 		Token: "your_existing_refresh_token",
// 	}
// 	logoutResp, err := db.User().Logout(logoutReq)

// 	assert.NoError(t, err)
// 	assert.True(t, logoutResp.Success)
// }
