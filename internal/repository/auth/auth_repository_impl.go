package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"ganasa18/attandance-be-mobile/internal/helper"
	"ganasa18/attandance-be-mobile/internal/model/domain"
	"log"
	"path/filepath"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (repository *AuthRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, register domain.UserRegisterModel) domain.UserRegisterModel {

	SQL := "INSERT INTO ms_users (user_unique_id, username, email, password, user_role, created_at, updated_at) values($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	var lastInsertID int
	values := []interface{}{register.UserUniqueId, register.Username, register.Email, register.Password, register.UserRole, register.CreatedAt, register.UpdatedAt}
	err := tx.QueryRowContext(ctx, SQL, values...).Scan(&lastInsertID)
	helper.PanicIfError(err)
	register.Id = lastInsertID

	// LOGGING
	log.Printf("Executing SQL: %s with values: %v\n", SQL, helper.FormatArraySplitByComma(values))

	return register
}

func (repository *AuthRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, login domain.UserModel) (domain.UserModel, error) {

	// Check if the user exists in the database
	checkUserSQL := "SELECT email, password FROM ms_users WHERE email = $1"
	row := tx.QueryRowContext(ctx, checkUserSQL, login.Email)

	// Logging
	log.Printf("Executing SQL: %s\n", checkUserSQL)

	var storedPasswordHash string
	err := row.Scan(&login.Email, &storedPasswordHash)
	if err == sql.ErrNoRows {
		// Handle the case where the email is not registered
		return domain.UserModel{}, errors.New("email not registered")
	}
	helper.PanicIfError(err)

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(login.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		// Handle the case where aunthenticated failed
		return domain.UserModel{}, errors.New("authentication failed")
	}
	helper.PanicIfError(err)

	getUserSQL := "SELECT id, user_unique_id, username, email, token, is_active, user_role, created_at, updated_at FROM ms_users WHERE email = $1"
	row = tx.QueryRowContext(ctx, getUserSQL, login.Email)

	// Logging
	log.Printf("Executing SQL: %s\n", getUserSQL)

	var user domain.UserModel

	err = row.Scan(&user.Id, &user.UserUniqueId, &user.Username, &user.Email, &user.Token, &user.IsActive, &user.UserRole, &user.CreatedAt, &user.UpdatedAt)
	helper.PanicIfError(err)

	// CREATE JWT TOKEN

	tokenString, err := helper.CreateToken(user.Id, user.Username, user.Email)
	helper.PanicIfError(err)
	user.Token = &tokenString

	// INSERT OTP
	otpNumber := helper.GenerateOtpNumber()
	OTPSQL := "INSERT INTO otp (otp, created_at) VALUES ($1, $2) RETURNING id"
	row = tx.QueryRowContext(ctx, OTPSQL, otpNumber, time.Now())

	// LOGGING
	log.Printf("Executing SQL: %s\n", OTPSQL)
	var otpID int
	err = row.Scan(&otpID)
	if err != nil {
		helper.PanicIfError(err)
	}

	templateData := struct {
		Name string
		URL  string
		OTP  string
	}{
		Name: user.Username,
		URL:  "http://geektrust.in",
		OTP:  otpNumber,
	}

	// Create a new request
	r := helper.NewRequest([]string{user.Email}, "Hello MAttandance", "Hello, World!")
	// Parse the template with template data
	templatePath := filepath.Join("..", "..", "internal", "helper", "template", "email.html")
	helper.PanicIfError(r.ParseTemplate(templatePath, templateData))
	// Send the email if template parsing is successful
	if ok, err := r.SendEmail(); err == nil {
		fmt.Println(ok)
	}

	return user, nil

}

func (repository *AuthRepositoryImpl) LoginWithOutOtp(ctx context.Context, tx *sql.Tx, login domain.UserModel) (domain.UserModel, error) {

	// Check if the user exists in the database
	checkUserSQL := "SELECT email, password FROM ms_users WHERE email = $1"
	row := tx.QueryRowContext(ctx, checkUserSQL, login.Email)

	// Logging
	log.Printf("Executing SQL: %s\n", checkUserSQL)

	var storedPasswordHash string
	err := row.Scan(&login.Email, &storedPasswordHash)
	if err == sql.ErrNoRows {
		// Handle the case where the email is not registered
		return domain.UserModel{}, errors.New("email not registered")
	}
	helper.PanicIfError(err)

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(login.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		// Handle the case where aunthenticated failed
		return domain.UserModel{}, errors.New("authentication failed")
	}
	helper.PanicIfError(err)

	getUserSQL := "SELECT id, user_unique_id, username, email, token, is_active, user_role, created_at, updated_at FROM ms_users WHERE email = $1"
	row = tx.QueryRowContext(ctx, getUserSQL, login.Email)

	// Logging
	log.Printf("Executing SQL: %s\n", getUserSQL)

	var user domain.UserModel

	err = row.Scan(&user.Id, &user.UserUniqueId, &user.Username, &user.Email, &user.Token, &user.IsActive, &user.UserRole, &user.CreatedAt, &user.UpdatedAt)
	helper.PanicIfError(err)

	// CREATE JWT TOKEN

	tokenString, err := helper.CreateToken(user.Id, user.Username, user.Email)
	helper.PanicIfError(err)
	user.Token = &tokenString

	return user, nil

}
