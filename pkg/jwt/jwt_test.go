package jwt_test

import (
	"judo/pkg/jwt"
	"testing"
)

const secret = "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9." +
	"\neyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0." +
	"\ntyh-VfuzIxCyGYDlkBA7DfyjrqmSHu6pQ2hoZuFqUSLPNY2N0mpHb3nk5K17HWP_3cYHBw7AhHale5wky6-sVA"

func TestJWT_CreateAndParse(t *testing.T) {
	j := jwt.NewJWT(secret)

	testCase := []struct {
		name      string
		email     string
		wantErr   bool
		wantValid bool
	}{
		{
			name:      "success create jwt",
			email:     "madk1d2321@gmail.com",
			wantErr:   false,
			wantValid: true,
		},
		{
			name:      "empty email",
			email:     "",
			wantErr:   true,
			wantValid: false,
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			var token string
			var err error

			token, err = j.GenerateToken(jwt.JWTData{
				Email: tt.email,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken error = %v, wantErr %v", err, tt.wantErr)
			}

			data, valid := j.ParseToken(token)
			if valid != tt.wantValid {
				t.Errorf("ParseToken: got %v want %v", valid, tt.wantValid)
			}

			if tt.wantValid && data.Email != tt.email {
				t.Errorf("expected email %s, got %s", tt.email, data.Email)
			}
		})
	}
}
