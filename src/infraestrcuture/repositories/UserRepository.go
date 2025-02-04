package repositories

import (
    "database/sql"
    "Practica/src/domain/entities"
)

type userRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(id int) (*entities.User, error) {
    user := &entities.User{}
    err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (r *userRepository) UpdateUser(user *entities.User) error {
    _, err := r.db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
    return err
}

func (r *userRepository) CheckForChanges() ([]entities.User, error) {
    rows, err := r.db.Query("SELECT id, name, email FROM users WHERE updated_at > NOW() - INTERVAL 5 SECOND")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []entities.User
    for rows.Next() {
        var user entities.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}