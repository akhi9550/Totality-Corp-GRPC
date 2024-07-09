package repository

import (
	"errors"
	"regexp"
	"testing"

	"grpc-user-service/pkg/utils/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_GetUserByID(t *testing.T) {
	tests := []struct {
		name    string
		args    int64
		stub    func(mockSQL sqlmock.Sqlmock)
		want    models.Users
		wantErr error
	}{
		{
			name: "success",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectQuery(`SELECT id, fname, city, phone, height, married FROM users WHERE id=\$1`).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "fname", "city", "phone", "height", "married"}).
						AddRow(1, "akhil", "bangalore", "9087678564", 157.9, true))
			},
			want: models.Users{
				ID:      1,
				Fname:   "akhil",
				City:    "bangalore",
				Phone:   "9087678564",
				Height:  157.9,
				Married: true,
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectQuery(`SELECT id, fname, city, phone, height, married FROM users WHERE id=\$1`).
					WithArgs(1).
					WillReturnError(errors.New("error"))
			},
			want:    models.Users{},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()
			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})
			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			result, err := u.GetUserByID(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_GetUsersByIDs(t *testing.T) {
	tests := []struct {
		name    string
		args    []int64
		stub    func(mockSQL sqlmock.Sqlmock)
		want    []models.Users
		wantErr error
	}{
		{
			name: "success",
			args: []int64{1, 2},
			stub: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectQuery(`SELECT id, fname, city, phone, height, married FROM users WHERE id=\$1`).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "fname", "city", "phone", "height", "married"}).
						AddRow(1, "akhil", "bangalore", "9087678564", 165.9, true))
				mockSQL.ExpectQuery(`SELECT id, fname, city, phone, height, married FROM users WHERE id=\$1`).
					WithArgs(2).
					WillReturnRows(sqlmock.NewRows([]string{"id", "fname", "city", "phone", "height", "married"}).
						AddRow(2, "rahul", "mumbai", "9076543210", 157.8, false))
			},
			want: []models.Users{
				{
					ID:      1,
					Fname:   "akhil",
					City:    "bangalore",
					Phone:   "9087678564",
					Height:  165.9,
					Married: true,
				},
				{
					ID:      2,
					Fname:   "rahul",
					City:    "mumbai",
					Phone:   "9076543210",
					Height:  157.8,
					Married: false,
				},
			},
			wantErr: nil,
		},
		{
			name: "partial success",
			args: []int64{1, 3},
			stub: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectQuery(`SELECT id, fname, city, phone, height, married FROM users WHERE id=\$1`).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "fname", "city", "phone", "height", "married"}).
						AddRow(1, "akhil", "bangalore", "9087678564", 157.9, true))
				mockSQL.ExpectQuery(`SELECT id, fname, city, phone, height, married FROM users WHERE id=\$1`).
					WithArgs(3).
					WillReturnError(errors.New("error"))
			},
			want: []models.Users{
				{
					ID:      1,
					Fname:   "akhil",
					City:    "bangalore",
					Phone:   "9087678564",
					Height:  157.9,
					Married: true,
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()
			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})
			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			result, err := u.GetUsersByIDs(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_SearchCity(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		stub    func(mockSQL sqlmock.Sqlmock)
		want    []models.Users
		wantErr error
	}{
		{
			name: "success",
			args: "City",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := `SELECT id, fname, city, phone, height, married FROM users WHERE city ILIKE '%' || \$1 || '%'`
				mockSQL.ExpectQuery(expectQuery).WithArgs("City").WillReturnRows(sqlmock.NewRows([]string{"id", "fname", "city", "phone", "height", "married"}).AddRow(1, "Akhil", "City", "1234567890", 157.6, true))
			},
			want: []models.Users{
				{
					ID:      1,
					Fname:   "Akhil",
					City:    "City",
					Phone:   "1234567890",
					Height:  157.6,
					Married: true,
				},
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: "City",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := `SELECT id, fname, city, phone, height, married FROM users WHERE city ILIKE '%' || \$1 || '%'`
				mockSQL.ExpectQuery(expectQuery).WithArgs("City").WillReturnError(errors.New("error"))
			},
			want:    []models.Users{},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()
			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})
			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			result, err := u.SearchCity(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_SearchPhone(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		stub    func(mockSQL sqlmock.Sqlmock)
		want    []models.Users
		wantErr error
	}{
		{
			name: "success",
			args: "1234567890",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := `SELECT id, fname, city, phone, height, married FROM users WHERE phone=\$1`
				mockSQL.ExpectQuery(expectQuery).WithArgs("1234567890").WillReturnRows(sqlmock.NewRows([]string{"id", "fname", "city", "phone", "height", "married"}).AddRow(1, "Akhil", "City", "1234567890", 157.6, true))
			},
			want: []models.Users{
				{
					ID:      1,
					Fname:   "Akhil",
					City:    "City",
					Phone:   "1234567890",
					Height:  157.6,
					Married: true,
				},
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: "1234567890",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := `SELECT id, fname, city, phone, height, married FROM users WHERE phone=\$1`
				mockSQL.ExpectQuery(expectQuery).WithArgs("1234567890").WillReturnError(errors.New("error"))
			},
			want:    []models.Users{},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()
			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})
			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			result, err := u.SearchPhone(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_SearchMarried(t *testing.T) {
	tests := []struct {
		name    string
		args    bool
		stub    func(mockSQL sqlmock.Sqlmock)
		want    []models.Users
		wantErr error
	}{
		{
			name: "success",
			args: true,
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := `SELECT id, fname, city, phone, height, married FROM users WHERE married=\$1`
				mockSQL.ExpectQuery(expectQuery).WithArgs(true).WillReturnRows(sqlmock.NewRows([]string{"id", "fname", "city", "phone", "height", "married"}).AddRow(1, "Akhil", "City", "1234567890", 157.6, true))
			},
			want: []models.Users{
				{
					ID:      1,
					Fname:   "Akhil",
					City:    "City",
					Phone:   "1234567890",
					Height:  157.6,
					Married: true,
				},
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: true,
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := `SELECT id, fname, city, phone, height, married FROM users WHERE married=\$1`
				mockSQL.ExpectQuery(expectQuery).WithArgs(true).WillReturnError(errors.New("error"))
			},
			want:    []models.Users{},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()
			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})
			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			result, err := u.SearchMarried(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_AddUser(t *testing.T) {
	tests := []struct {
		name    string
		args    models.User
		stub    func(mockSQL sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "success",
			args: models.User{
				Fname:   "Akhil",
				City:    "City",
				Phone:   "1234567890",
				Height:  157.6,
				Married: true,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := regexp.QuoteMeta("INSERT INTO users (fname, city, phone, height, married) VALUES ($1, $2, $3, $4, $5)")
				mockSQL.ExpectExec(expectQuery).WithArgs("Akhil", "City", "1234567890", sqlmock.AnyArg(), true).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: models.User{
				Fname:   "Akhil",
				City:    "City",
				Phone:   "1234567890",
				Height:  157.6,
				Married: true,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := regexp.QuoteMeta("INSERT INTO users (fname, city, phone, height, married) VALUES ($1, $2, $3, $4, $5)")
				mockSQL.ExpectExec(expectQuery).WithArgs("Akhil", "City", "1234567890", sqlmock.AnyArg(), true).WillReturnError(errors.New("error"))
			},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()
			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})
			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			err := u.AddUser(tt.args)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_CheckUserExistsByPhone(t *testing.T) {
	tests := []struct {
		name string
		args string
		stub func(mockSQL sqlmock.Sqlmock)
		want bool
	}{
		{
			name: "exists",
			args: "1234567890",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := regexp.QuoteMeta("SELECT count(*) FROM users WHERE phone = $1")
				mockSQL.ExpectQuery(expectQuery).WithArgs("1234567890").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
			},
			want: true,
		},
		{
			name: "not exists",
			args: "1234567890",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := regexp.QuoteMeta("SELECT count(*) FROM users WHERE phone = $1")
				mockSQL.ExpectQuery(expectQuery).WithArgs("1234567890").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
			},
			want: false,
		},
		{
			name: "error",
			args: "1234567890",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := regexp.QuoteMeta("SELECT count(*) FROM users WHERE phone = ?")
				mockSQL.ExpectQuery(expectQuery).WithArgs("1234567890").WillReturnError(errors.New("error"))
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()
			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})
			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			result := u.CheckUserExistsByPhone(tt.args)

			assert.Equal(t, tt.want, result)
		})
	}
}

func Test_CheckUserAvailabilityWithUserID(t *testing.T) {
	tests := []struct {
		name string
		args int64
		stub func(mockSQL sqlmock.Sqlmock)
		want bool
	}{
		{
			name: "exists",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := `SELECT count\(\*\) FROM users WHERE id = \$1`
				mockSQL.ExpectQuery(expectQuery).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
			},
			want: true,
		},
		{
			name: "not exists",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := `SELECT count\(\*\) FROM users WHERE id = \$1`
				mockSQL.ExpectQuery(expectQuery).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
			},
			want: false,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := `SELECT count\(\*\) FROM users WHERE id = \$1`
				mockSQL.ExpectQuery(expectQuery).WithArgs(1).WillReturnError(errors.New("error"))
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()
			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})
			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			result := u.CheckUserAvailabilityWithUserID(tt.args)

			assert.Equal(t, tt.want, result)
		})
	}
}

func Test_CheckUserAvailabilityWithUserIDs(t *testing.T) {
	tests := []struct {
		name string
		args []int64
		stub func(mockSQL sqlmock.Sqlmock)
		want bool
	}{
		{
			name: "all exist",
			args: []int64{1, 2, 3},
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := regexp.QuoteMeta("SELECT count(*) FROM users WHERE id = $1")
				mockSQL.ExpectQuery(expectQuery).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mockSQL.ExpectQuery(expectQuery).WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mockSQL.ExpectQuery(expectQuery).WithArgs(3).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
			},
			want: true,
		},
		{
			name: "some exist",
			args: []int64{1, 2, 3},
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := regexp.QuoteMeta("SELECT count(*) FROM users WHERE id = $1")
				mockSQL.ExpectQuery(expectQuery).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mockSQL.ExpectQuery(expectQuery).WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mockSQL.ExpectQuery(expectQuery).WithArgs(3).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
			},
			want: false,
		},
		{
			name: "error",
			args: []int64{1, 2, 3},
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectQuery := regexp.QuoteMeta("SELECT count(*) FROM users WHERE id = $1")
				mockSQL.ExpectQuery(expectQuery).WithArgs(1).WillReturnError(errors.New("error"))
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()
			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})
			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			result := u.CheckUserAvailabilityWithUserIDs(tt.args)

			assert.Equal(t, tt.want, result)
		})
	}
}
