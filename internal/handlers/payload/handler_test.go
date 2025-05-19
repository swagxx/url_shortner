package payload

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"judo/configs"
	"judo/internal/user"
	"judo/pkg/db"
	"judo/pkg/dto"
	"judo/pkg/jwt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func bootstrap() (*AuthHandler, sqlmock.Sqlmock, error) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbMock,
	}))
	if err != nil {
		return nil, nil, err
	}
	myDb := db.DB{
		DB: gormDb,
	}

	userRepo := user.NewUserRepository(&myDb)

	handler := AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: user.NewAuthService(userRepo, &jwt.JWT{}),
	}
	return &handler, mock, nil
}

func TestLoginSuccess(t *testing.T) {
	handler, mock, err := bootstrap()

	hashed, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("test123@gmail.com", string(hashed))
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	if err != nil {
		t.Fatalf("bootstrap() failed: %v", err)
	}

	data, _ := json.Marshal(&dto.LoginRequest{
		Email:    "madk1d302@gmail.com",
		Password: "123456",
	})

	reader := bytes.NewReader(data)
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	req.Header.Set("Content-Type", "application/json")
	handler.Login()(wr, req)

	if wr.Code != http.StatusOK {
		t.Errorf("login() failed: got %v, want %v", wr.Code, http.StatusOK)
	}
}

func TestRegisterSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatalf("bootstrap() failed: %v", err)
	}

	rows := sqlmock.NewRows([]string{"email", "password", "name"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"email"}).AddRow("test123@gmail.com"))
	mock.ExpectCommit()

	data, _ := json.Marshal(&dto.RegisterRequest{
		Username: "Алекс",
		Email:    "madk1d302@gmail.com",
		Password: "123456",
	})

	reader := bytes.NewReader(data)
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/register", reader)
	req.Header.Set("Content-Type", "application/json")
	handler.Register()(wr, req)
	if wr.Code != http.StatusCreated {
		t.Errorf("register() failed: got %v, want %v", wr.Code, http.StatusCreated)
	}
}
