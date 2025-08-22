package sqlite

import (
	"database/sql"
	"time"

	core "workshop-cursor/backend/internal/core/user"
)

type SQLiteUserRepository struct {
	db *sql.DB
}

func NewSQLiteUserRepository(db *sql.DB) *SQLiteUserRepository {
	return &SQLiteUserRepository{db: db}
}

func (r *SQLiteUserRepository) InitSchema() error {
	_, err := r.db.Exec(`
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    phone TEXT,
    member_code TEXT NOT NULL UNIQUE,
    membership_level TEXT NOT NULL,
    points INTEGER NOT NULL DEFAULT 0,
    joined_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);
`)
	return err
}

func (r *SQLiteUserRepository) FindByEmail(email string) (*core.User, error) {
	row := r.db.QueryRow(`SELECT id, email, password_hash, first_name, last_name, phone, member_code, membership_level, points, joined_at, created_at, updated_at FROM users WHERE email = ?`, email)
	u := &core.User{}
	err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.FirstName, &u.LastName, &u.Phone, &u.MemberCode, &u.MembershipLevel, &u.Points, &u.JoinedAt, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *SQLiteUserRepository) FindByID(id int64) (*core.User, error) {
	row := r.db.QueryRow(`SELECT id, email, password_hash, first_name, last_name, phone, member_code, membership_level, points, joined_at, created_at, updated_at FROM users WHERE id = ?`, id)
	u := &core.User{}
	err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.FirstName, &u.LastName, &u.Phone, &u.MemberCode, &u.MembershipLevel, &u.Points, &u.JoinedAt, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *SQLiteUserRepository) UpdateProfile(id int64, input core.UpdateProfileInput) (*core.User, error) {
	_, err := r.db.Exec(`UPDATE users SET first_name = ?, last_name = ?, phone = ?, updated_at = ? WHERE id = ?`, input.FirstName, input.LastName, input.Phone, time.Now(), id)
	if err != nil {
		return nil, err
	}
	return r.FindByID(id)
}

func (r *SQLiteUserRepository) SeedInitialUserIfEmpty(seed *core.User) error {
	var count int
	if err := r.db.QueryRow(`SELECT COUNT(1) FROM users`).Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	_, err := r.db.Exec(`INSERT INTO users (email, password_hash, first_name, last_name, phone, member_code, membership_level, points, joined_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		seed.Email, seed.PasswordHash, seed.FirstName, seed.LastName, seed.Phone, seed.MemberCode, seed.MembershipLevel, seed.Points, seed.JoinedAt, seed.CreatedAt, seed.UpdatedAt,
	)
	return err
}

var _ core.UserRepository = (*SQLiteUserRepository)(nil)
