package repository

import (
    "context"
    "database/sql"
    "shopping/internal/models"
    "log"
    "golang.org/x/crypto/bcrypt"
    "errors"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) ValidateUser(ctx context.Context, sessionId string) (models.User, error) {

    rows, err := r.db.QueryContext(ctx, "SELECT u.id, u.first_name, u.last_name, u.username FROM user u LEFT JOIN session s ON s.user_id=u.id WHERE s.session_id = ?", sessionId)
    user := models.User{}
    if err != nil {
        log.Fatal(err)
        return user, err
    }
    for rows.Next() {
        if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username); err != nil {
            return user, err
        }
    }
    defer rows.Close()
	return user, nil
}

func (r *UserRepository) Register(ctx context.Context, username, password, firstName, lastName, token string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        // handle error
    }
    res, err := r.db.ExecContext(ctx, "INSERT INTO user (username, password, first_name, last_name, is_verified) VALUES (?, ?, ?, ?, ?)", username, hashedPassword, firstName, lastName, 0)
    if err != nil {
        return err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return err
    }
    // Add token to user_token table and add expiry time one hour from now
    _, err = r.db.ExecContext(ctx, "INSERT INTO user_token (user_id, token, expiry_time) VALUES (?, ?, DATE_ADD(NOW(), INTERVAL 1 HOUR))", id, token)
    if err != nil {
        return err
    }
    return nil
}

func (r *UserRepository) ValidateLogin(ctx context.Context, username, password string) (models.User, error) {
    row := r.db.QueryRowContext(ctx, "SELECT id, first_name, last_name, username, password FROM user WHERE username = ?", username)

    user := models.User{}
    err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Password)
    if err != nil {
        log.Fatal(err)
        return user, err
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return user, errors.New("username or password incorrect")
    }

    return user, nil
}