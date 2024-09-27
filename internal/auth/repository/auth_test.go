package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/prajnasatryass/go-clean-arch-example/internal/auth/domain"
	userDomain "github.com/prajnasatryass/go-clean-arch-example/internal/user/domain"
	mockTokenUtil "github.com/prajnasatryass/go-clean-arch-example/mocks/pkg/tokenutil"
	"github.com/prajnasatryass/go-clean-arch-example/pkg/tokenutil"
	"reflect"
	"testing"
)

var (
	inputUser = userDomain.User{
		Email:    "user@example.com",
		Password: "123",
	}

	queryRefreshToken = "refresh_token"
	queryUserID       = uuid.New()

	refreshTokenRecord = domain.RefreshTokenRecord{
		Token:  queryRefreshToken,
		UserID: queryUserID,
	}
)

func TestNewAuthRepository(t *testing.T) {
	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want domain.AuthRepository
	}{
		{
			name: "success",
			args: args{
				db: nil,
			},
			want: &authRepository{
				db:        nil,
				tokenUtil: tokenutil.NewTokenUtil(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authRepository_CreateAccessToken(t *testing.T) {
	mockTestTokenUtil := mockTokenUtil.NewMockTokenUtil(t)

	type fields struct {
		db        *sqlx.DB
		tokenUtil tokenutil.TokenUtil
	}
	type args struct {
		user   *userDomain.User
		secret string
		ttl    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		pre     func()
		want    string
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				db:        nil,
				tokenUtil: mockTestTokenUtil,
			},
			args: args{
				user:   &inputUser,
				secret: "",
				ttl:    0,
			},
			pre: func() {
				mockTestTokenUtil.EXPECT().CreateAccessToken(&inputUser, "", 0).Return("", nil).Once()
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := &authRepository{
				db:        tt.fields.db,
				tokenUtil: tt.fields.tokenUtil,
			}
			tt.pre()
			got, err := ar.CreateAccessToken(tt.args.user, tt.args.secret, tt.args.ttl)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateAccessToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authRepository_CreateRefreshToken(t *testing.T) {
	mockTestTokenUtil := mockTokenUtil.NewMockTokenUtil(t)

	type fields struct {
		db        *sqlx.DB
		tokenUtil tokenutil.TokenUtil
	}
	type args struct {
		user   *userDomain.User
		secret string
		ttl    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		pre     func()
		want    string
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				db:        nil,
				tokenUtil: mockTestTokenUtil,
			},
			args: args{
				user:   &inputUser,
				secret: "",
				ttl:    0,
			},
			pre: func() {
				mockTestTokenUtil.EXPECT().CreateRefreshToken(&inputUser, "", 0).Return("", nil).Once()
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := &authRepository{
				db:        tt.fields.db,
				tokenUtil: tt.fields.tokenUtil,
			}
			tt.pre()
			got, err := ar.CreateRefreshToken(tt.args.user, tt.args.secret, tt.args.ttl)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateRefreshToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authRepository_DeleteExpiredRefreshTokens(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal("db mock error: " + err.Error())
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	type fields struct {
		db *sqlx.DB
	}
	tests := []struct {
		name    string
		fields  fields
		pre     func()
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				db: sqlxDB,
			},
			pre: func() {
				dbMock.ExpectExec("DELETE FROM jwt_refresh_tokens").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := &authRepository{
				db: tt.fields.db,
			}
			tt.pre()
			if err := ar.DeleteExpiredRefreshTokens(); (err != nil) != tt.wantErr {
				t.Errorf("DeleteExpiredRefreshTokens() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authRepository_DeleteRefreshToken(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal("db mock error: " + err.Error())
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		refreshToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		pre     func()
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				db: sqlxDB,
			},
			args: args{
				refreshToken: queryRefreshToken,
			},
			pre: func() {
				dbMock.ExpectExec("DELETE FROM jwt_refresh_tokens").WithArgs(queryRefreshToken).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := &authRepository{
				db: tt.fields.db,
			}
			tt.pre()
			if err := ar.DeleteRefreshToken(tt.args.refreshToken); (err != nil) != tt.wantErr {
				t.Errorf("DeleteRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authRepository_DeleteUserRefreshTokens(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal("db mock error: " + err.Error())
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		pre     func()
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				db: sqlxDB,
			},
			args: args{
				userID: queryUserID,
			},
			pre: func() {
				dbMock.ExpectExec("DELETE FROM jwt_refresh_tokens").WithArgs(queryUserID).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := &authRepository{
				db: tt.fields.db,
			}
			tt.pre()
			if err := ar.DeleteUserRefreshTokens(tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserRefreshTokens() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authRepository_RetrieveRefreshToken(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal("db mock error: " + err.Error())
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		refreshToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		pre     func()
		want    domain.RefreshTokenRecord
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				db: sqlxDB,
			},
			args: args{
				refreshToken: queryRefreshToken,
			},
			pre: func() {
				rows := sqlmock.NewRows([]string{"token", "user_id", "ignore_after"}).AddRow(refreshTokenRecord.Token, refreshTokenRecord.UserID, refreshTokenRecord.IgnoreAfter)
				dbMock.ExpectQuery("SELECT \\* FROM jwt_refresh_tokens").WithArgs(queryRefreshToken).WillReturnRows(rows)
			},
			want:    refreshTokenRecord,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := &authRepository{
				db: tt.fields.db,
			}
			tt.pre()
			got, err := ar.RetrieveRefreshToken(tt.args.refreshToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("RetrieveRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RetrieveRefreshToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authRepository_StoreRefreshToken(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal("db mock error: " + err.Error())
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		record *domain.RefreshTokenRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		pre     func()
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				db: sqlxDB,
			},
			args: args{
				record: &refreshTokenRecord,
			},
			pre: func() {
				dbMock.ExpectExec("INSERT INTO jwt_refresh_tokens").WithArgs(refreshTokenRecord.Token, refreshTokenRecord.UserID, refreshTokenRecord.IgnoreAfter).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := &authRepository{
				db: tt.fields.db,
			}
			tt.pre()
			if err := ar.StoreRefreshToken(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("StoreRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
