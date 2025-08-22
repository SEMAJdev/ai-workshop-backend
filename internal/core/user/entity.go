package user

import "time"

// User represents an application member profile.
type User struct {
    ID               int64
    Email            string
    PasswordHash     string
    FirstName        string
    LastName         string
    Phone            string
    MemberCode       string
    MembershipLevel  string
    Points           int64
    JoinedAt         time.Time
    CreatedAt        time.Time
    UpdatedAt        time.Time
}

// PublicProfile is a safe view of user data returned to clients.
type PublicProfile struct {
    ID              int64     `json:"id"`
    Email           string    `json:"email"`
    FirstName       string    `json:"firstName"`
    LastName        string    `json:"lastName"`
    Phone           string    `json:"phone"`
    MemberCode      string    `json:"memberCode"`
    MembershipLevel string    `json:"membershipLevel"`
    Points          int64     `json:"points"`
    JoinedAt        time.Time `json:"joinedAt"`
}

// UpdateProfileInput captures editable fields.
type UpdateProfileInput struct {
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
    Phone     string `json:"phone"`
}

// UserRepository defines persistence operations for users.
type UserRepository interface {
    FindByEmail(email string) (*User, error)
    FindByID(id int64) (*User, error)
    UpdateProfile(id int64, input UpdateProfileInput) (*User, error)
    SeedInitialUserIfEmpty(seed *User) error
}


