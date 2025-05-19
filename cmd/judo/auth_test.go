package main

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"judo/configs"
	typesimpo "judo/internal/types"
	"judo/pkg/db"
	"judo/pkg/dto"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func initCfg() *configs.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN not found")
	}
	return &configs.Config{
		DB: configs.DBConfig{
			DSN: dsn,
		},
	}
}

func initData(db *db.DB) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	db.Create(&typesimpo.User{
		Email:    "test123@gmail.com",
		Password: string(hashed),
		Name:     "Майкл",
	})
}

func initDb(cfg *configs.Config) *db.DB {
	gormDB, err := gorm.Open(postgres.Open(cfg.DB.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to the database")
	}
	return &db.DB{
		DB: gormDB,
	}
}

func TestLogin(t *testing.T) {
	cfg := initCfg()
	db := initDb(cfg)

	db.AutoMigrate(&typesimpo.User{})
	if err := db.Exec("DELETE FROM users;").Error; err != nil {
		log.Fatal("failed to clean up users: ", db)
	}

	t.Cleanup(func() {
		if err := db.Exec("DELETE FROM users;").Error; err != nil {
			log.Fatal("failed to clean up users: ", db)
		}
	})

	initData(db)

	ts := httptest.NewServer(App(db, cfg))
	defer ts.Close()

	test := []struct {
		name       string
		email      string
		password   string
		expectCode int
		err        bool
	}{
		{
			name:       "success login",
			email:      "test123@gmail.com",
			password:   "123456",
			expectCode: http.StatusOK,
			err:        false,
		},
		{
			name:       "fail login",
			email:      "noemail342@gmail.com",
			password:   "123456",
			expectCode: http.StatusUnauthorized,
			err:        true,
		},
		{
			name:       "empty email",
			email:      "",
			password:   "123456",
			expectCode: http.StatusUnauthorized,
			err:        true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			data, _ := json.Marshal(&dto.LoginRequest{
				Email:    tt.email,
				Password: tt.password,
			})

			resp, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
			if err != nil {
				t.Error(err)
				return
			}
			if resp.StatusCode != tt.expectCode {
				t.Errorf("got %v, want %v", resp.StatusCode, tt.expectCode)
				return
			}

			if tt.err == false {
				var respBody map[string]interface{}
				json.NewDecoder(resp.Body).Decode(&respBody)
				if _, ok := respBody["token"]; !ok {
					t.Errorf("expected token in response, got %v", respBody)
				}
			}
		})
	}
}
