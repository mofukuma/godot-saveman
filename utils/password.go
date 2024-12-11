package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// ユーザ認証用。
// HashPassword は、与えられたパスワードをハッシュ化し、ハッシュ化されたパスワードを文字列として返す。
// bcrypt アルゴリズムを使用してハッシュ化を行います。
// パスワードのハッシュ化に失敗した場合、エラーを返します。
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}

// VerifyPassword は、ハッシュ化されたパスワードと候補のパスワードを比較し、エラーを返します。
func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
