package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/radityacandra/go-cms/internal/application/auth/service"
	"github.com/radityacandra/go-cms/internal/application/auth/types"
	"github.com/radityacandra/go-cms/internal/application/user/model"
	"github.com/radityacandra/go-cms/internal/application/user/repository"
	userTypes "github.com/radityacandra/go-cms/internal/application/user/types"
	mockRepository "github.com/radityacandra/go-cms/mocks/internal_/application/user/repository"
	"github.com/radityacandra/go-cms/pkg/hash"
	"github.com/radityacandra/go-cms/pkg/jwt"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	privateKey := `
-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCjbZARaqvbqTIb
aJvrSqBoeQD4xQreM0Phyh3x03QITwF06cNJ1Nz/S1/mrbzWovl0LGe2vwYKggkt
Q61IH3y11HTE+VdYt3eODBTQV+1Lb3vB2NoGX9qZsjTMSGsvh35HqBo71ePcfPOl
FNyrw940EU+iLC++gTyqBOOwds4+DLmSxFb9D+H1OVO3/+8Wl56MAfhTZMowH3JD
z/NoJxU0Op5XyIYYtT/jtVBJwv92pq6XkKE9p+vb0NJuuQrwU8QtZzIHOZVWnYZE
FtIgDLAY+O9Do8++zBQOb7XS4XPYu0einGYo7tPViUaNOnQ65ekh5nlI+rBBrXSJ
fOrET0VjAgMBAAECggEAAh/in3VvheT3yiWqTgAcMU341BhrGMY/+wVJm3bjSFfk
frGGOqnjohKpXxh0G+NBl3U0OofKM8RvbLl6e4SqCK9jNYK1KXwEt/upahVjTVAC
+ELZmESGNqCsVFNvcp2EVaCqOkLDmJbixx1nh2BljL8ExTfr4h2TRmNpQ8ZSqXBw
ja1FktikmRuM7H/pwZvPm+hGB9aoWqsqG4CL7wpY0mnJcQhfEtCYQs5B5WF9CROl
LrTDiegNya76zGSURDVX10qYhEDpr3QhcSiqAD5J43/fi8tF34n+jLYsA5mDBHzZ
o5xD196AglaajlJgE3Xg5NquWzx4P4GsbQVAPe6amQKBgQDmhei9CogkVoJtcjaK
6HH2xmyn8IW6sPUOrVA4NRxYObmsOs/POIIr8dWfSx0fOhQkexhsQOXorR1BMyg2
9W4G7O7UnO6dusp+dExBa/H4YUssvid3kIA/nmDCTpayLp9tGmgDqln8Ydj/clx8
CHgrB1cS8lirVzegLWG759AsqQKBgQC1fVxRD3Jsamh6yV+ae0VgUGkkFQbaCtxY
q6hmn93UWhgxtELZtDF+2hLDGKjiZ/6Ru3IWjViD0qzBN/muxwoKEcH5rFcpfgvl
JfX1LxlMFqThzoU8rAzcN4YAXrl524qZT8Pa+0ayWuh94sbtharl1wmS4H1xvv5E
TRCfjMe9KwKBgGyQ6mCFokKC36BN5vQarvmKz8d0FncrOe50n1ApazT90r0TYaV1
NpEdrv77cDaxsqTPuFvbYKvpQ9reDfV8NxpZ4c0OL67nNtDBUtyIywewQqhK0emU
i2Hq5CT+wCggnwLSKeR4CTM8necIZBgiIP4a0d5hdnMTe2YbmWjWrwspAoGAX709
0dUgO2j2rK9GK6wTsPc6P7qH3sYT7wK+10RGNRtB4BaDnWydH5nSg/CiRq0tcZs8
WAFATGn6kAMDR9vfw+gSN69eW5kOlVctJKYv5h+b4zKavqLUNedkXRWbKllSCAY3
/3DGnpeuRZo37lyxBoYlmsGp6zMh1s4AkuolA/kCgYBd4pAi4EvGRLVGCt7lcsLB
ou+8oj9iddh6JJ0e//22SS+/j81tIvKxk5TTz++fUUSG9ZhIGRKLAboMAmSW+PgB
Y9zcozoszYaT4GqXQw6BthHz6fwFicLOWYMDCqXYN6BYMQEUMvQR5sQ6F2fkXwnc
szLQ5UyjT6hr1KxVnGRfeQ==
-----END PRIVATE KEY-----
`

	var userId = uuid.NewString()

	token, exp, _ := jwt.BuildToken(map[string]interface{}{
		"sub": userId,
		"scopes": []string{
			"get-profile",
		},
	}, privateKey)

	type fields struct {
		Repository repository.IRepository
		PrivateKey string
	}
	type args struct {
		ctx   context.Context
		input types.LoginInput
	}
	type test struct {
		name    string
		fields  fields
		args    args
		mock    func(tt test) test
		want    types.LoginOutput
		wantErr error
	}

	tests := []test{
		{
			name: "should return error if user is not found",
			args: args{
				ctx: context.Background(),
				input: types.LoginInput{
					Username: "someuser",
					Password: "somepassword",
				},
			},
			wantErr: types.ErrUserNotFound,
			fields: fields{
				PrivateKey: "someprivatekey",
			},
			mock: func(tt test) test {
				mockRepository := mockRepository.NewMockIRepository(t)
				tt.fields.Repository = mockRepository

				mockRepository.EXPECT().FindUserByUsername(tt.args.ctx, userTypes.FindUserByUsernameInput{
					Username: tt.args.input.Username,
				}).Return(nil, errors.New("some error")).Times(1)

				return tt
			},
		},
		{
			name: "should return error if password doesn't match",
			args: args{
				ctx: context.Background(),
				input: types.LoginInput{
					Username: "someuser",
					Password: "somepassword",
				},
			},
			wantErr: types.ErrPasswordMissmatch,
			fields: fields{
				PrivateKey: "someprivatekey",
			},
			mock: func(tt test) test {
				mockRepository := mockRepository.NewMockIRepository(t)
				tt.fields.Repository = mockRepository

				hashedPassw, _ := hash.GenerateHash("wrongpassword")

				mockRepository.EXPECT().FindUserByUsername(tt.args.ctx, userTypes.FindUserByUsernameInput{
					Username: tt.args.input.Username,
				}).Return(&model.User{
					Username: tt.args.input.Username,
					Password: hashedPassw,
				}, nil).Times(1)

				return tt
			},
		},
		{
			name: "should return error if failed to generate jwt token",
			args: args{
				ctx: context.Background(),
				input: types.LoginInput{
					Username: "someuser",
					Password: "somepassword",
				},
			},
			wantErr: types.ErrFailedGenerateToken,
			fields: fields{
				PrivateKey: "someprivatekey",
			},
			mock: func(tt test) test {
				mockRepository := mockRepository.NewMockIRepository(t)
				tt.fields.Repository = mockRepository

				hashedPassw, _ := hash.GenerateHash("somepassword")

				mockRepository.EXPECT().FindUserByUsername(tt.args.ctx, userTypes.FindUserByUsernameInput{
					Username: tt.args.input.Username,
				}).Return(&model.User{
					Username: tt.args.input.Username,
					Password: hashedPassw,
				}, nil).Times(1)

				return tt
			},
		},
		{
			name: "should successfully executed and return correct data",
			args: args{
				ctx: context.Background(),
				input: types.LoginInput{
					Username: "someuser",
					Password: "somepassword",
				},
			},
			want: types.LoginOutput{
				ExpiredAt: exp,
				Token:     token,
			},
			fields: fields{
				PrivateKey: privateKey,
			},
			mock: func(tt test) test {
				mockRepository := mockRepository.NewMockIRepository(t)
				tt.fields.Repository = mockRepository

				hashedPassw, _ := hash.GenerateHash("somepassword")

				mockRepository.EXPECT().FindUserByUsername(tt.args.ctx, userTypes.FindUserByUsernameInput{
					Username: tt.args.input.Username,
				}).Return(&model.User{
					Username: tt.args.input.Username,
					Password: hashedPassw,
					Id:       userId,
					UserRoles: []model.UserRole{
						{
							RoleAcls: []model.RoleAcl{
								{
									Access: "get-profile",
								},
							},
						},
					},
				}, nil).Times(1)

				return tt
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt = tt.mock(tt)

			s := &service.Service{
				Repository: tt.fields.Repository,
				PrivateKey: tt.fields.PrivateKey,
			}
			got, err := s.Login(tt.args.ctx, tt.args.input)

			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
