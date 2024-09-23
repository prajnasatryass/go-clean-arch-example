package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/prajnasatryass/tic-be/internal/user/domain"
	"reflect"
	"testing"
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
		user *domain.User
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
				user: &domain.User{},
			},
			pre: func() {
				dbMock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(1, 1))
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
			if err := ur.Create(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
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

	queryUserID := uuid.New()

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

	queryUserID := uuid.New()

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

	queryUserID := uuid.New()

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
