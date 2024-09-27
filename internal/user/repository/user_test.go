package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/prajnasatryass/tic-be/internal/user/domain"
	"github.com/prajnasatryass/tic-be/pkg/constants"
	"reflect"
	"testing"
)

var (
	inputUser = domain.User{
		Email:    "user@ticindo.com",
		Password: "123",
	}
	newUserID   = uuid.New()
	queryUserID = uuid.New()
	newRoleID   = constants.UserRoleRoot
)

func TestNewUserRepository(t *testing.T) {
	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want domain.UserRepository
	}{
		{
			name: "success",
			args: args{
				db: nil,
			},
			want: &userRepository{
				db: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_Create(t *testing.T) {
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
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		pre     func()
		want    uuid.UUID
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				db: sqlxDB,
			},
			args: args{
				email:    inputUser.Email,
				password: inputUser.Password,
			},
			pre: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(newUserID)
				dbMock.ExpectQuery("INSERT INTO users").WillReturnRows(rows)
			},
			want:    newUserID,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &userRepository{
				db: tt.fields.db,
			}
			tt.pre()
			got, err := ur.Create(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_DeleteByID(t *testing.T) {
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
		id uuid.UUID
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
				id: queryUserID,
			},
			pre: func() {
				dbMock.ExpectExec("UPDATE users").WithArgs(queryUserID).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &userRepository{
				db: tt.fields.db,
			}
			tt.pre()
			if err := ur.DeleteByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userRepository_GetByEmail(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal("db mock error: " + err.Error())
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	queryEmail := "user@ticindo.com"

	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		pre     func()
		want    domain.User
		wantErr bool
	}{
		{
			name:   "success",
			fields: fields{db: sqlxDB},
			args: args{
				email: queryEmail,
			},
			pre: func() {
				rows := sqlmock.NewRows([]string{"email"}).AddRow(queryEmail)
				dbMock.ExpectQuery("SELECT \\* FROM users").WithArgs(queryEmail).WillReturnRows(rows)
			},
			want: domain.User{
				Email: queryEmail,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &userRepository{
				db: tt.fields.db,
			}
			tt.pre()
			got, err := ur.GetByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_GetByID(t *testing.T) {
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
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		pre     func()
		want    domain.User
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				db: sqlxDB,
			},
			args: args{
				id: queryUserID,
			},
			pre: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(queryUserID)
				dbMock.ExpectQuery("SELECT \\* FROM users").WithArgs(queryUserID).WillReturnRows(rows)
			},
			want: domain.User{
				ID: queryUserID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &userRepository{
				db: tt.fields.db,
			}
			tt.pre()
			got, err := ur.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_PermaDeleteByID(t *testing.T) {
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
		id uuid.UUID
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
				id: queryUserID,
			},
			pre: func() {
				dbMock.ExpectExec("DELETE FROM users").WithArgs(queryUserID).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &userRepository{
				db: tt.fields.db,
			}
			tt.pre()
			if err := ur.PermaDeleteByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("PermaDeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userRepository_UpdateRoleByID(t *testing.T) {
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
		id     uuid.UUID
		roleID constants.UserRole
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
				id:     queryUserID,
				roleID: newRoleID,
			},
			pre: func() {
				dbMock.ExpectExec("UPDATE users").WithArgs(newRoleID, queryUserID).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &userRepository{
				db: tt.fields.db,
			}
			tt.pre()
			if err := ur.UpdateRoleByID(tt.args.id, tt.args.roleID); (err != nil) != tt.wantErr {
				t.Errorf("UpdateRoleByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
