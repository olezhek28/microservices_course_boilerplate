package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/neracastle/auth/internal/domain/user"
)

func TestNewUser(t *testing.T) {
	type testData struct {
		Name     string
		Email    string
		Password string
	}

	var (
		userData = testData{
			Name:     gofakeit.Name(),
			Email:    gofakeit.Email(),
			Password: gofakeit.Password(true, true, true, false, false, 8),
		}

		wrongEmailErr = errors.New("email не может быть пустым")
	)

	tests := []struct {
		name string
		args testData
		want *user.User
		err  error
	}{
		{
			name: "Success",
			args: userData,
			want: &user.User{
				Name:     userData.Name,
				Email:    userData.Email,
				Password: userData.Password,
				IsAdmin:  false,
				RegDate:  time.Now(),
			},
			err: nil,
		},
		{
			name: "Empty email",
			args: func(dto testData) testData {
				dto.Email = ""
				return dto
			}(userData),
			want: nil,
			err:  wrongEmailErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usr, err := user.NewUser(tt.args.Email, tt.args.Password, tt.args.Name)

			if usr != nil {
				usr.RegDate = tt.want.RegDate
			}

			require.Equal(t, tt.want, usr)
			require.Equal(t, tt.err, err)
		})
	}
}
