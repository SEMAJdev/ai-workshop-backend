package auth

import (
    "errors"
    "time"

    "golang.org/x/crypto/bcrypt"

    core "workshop-cursor/backend/internal/core/user"
)

type TokenSigner interface {
    Sign(userID int64, expiresAt time.Time) (string, error)
    Verify(token string) (int64, error)
}

type UseCase struct {
    users       core.UserRepository
    tokenSigner TokenSigner
}

func NewUseCase(users core.UserRepository, signer TokenSigner) *UseCase {
    return &UseCase{users: users, tokenSigner: signer}
}

func (u *UseCase) Login(email, password string) (string, *core.User, error) {
    usr, err := u.users.FindByEmail(email)
    if err != nil {
        return "", nil, errors.New("invalid credentials")
    }
    if bcrypt.CompareHashAndPassword([]byte(usr.PasswordHash), []byte(password)) != nil {
        return "", nil, errors.New("invalid credentials")
    }
    token, err := u.tokenSigner.Sign(usr.ID, time.Now().Add(24*time.Hour))
    if err != nil {
        return "", nil, err
    }
    return token, usr, nil
}

func (u *UseCase) GetProfile(userID int64) (*core.User, error) {
    return u.users.FindByID(userID)
}

func (u *UseCase) UpdateProfile(userID int64, input core.UpdateProfileInput) (*core.User, error) {
    return u.users.UpdateProfile(userID, input)
}


