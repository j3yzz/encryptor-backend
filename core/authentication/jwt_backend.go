package authentication

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/dgrijalva/jwt-go"
	"github.com/j3yzz/encryptor/services/models"
	"github.com/j3yzz/encryptor/settings"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type JWTAuthentication struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

const (
	TokenDuration = 72
	expireOffset  = 3600
)

var authBackendInstance *JWTAuthentication = nil

func InitJWTAuthentication() *JWTAuthentication {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthentication{
			PrivateKey: getPrivateKey(),
			PublicKey:  GetPublicKey(),
		}
	}

	return authBackendInstance
}

func (jwtAuthentication *JWTAuthentication) GenerateToken(id int) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(settings.Get().JWTExpiration)).Unix(),
		"iat": time.Now().Unix(),
		"sub": id,
	}

	tokenString, err := token.SignedString(jwtAuthentication.PrivateKey)
	if err != nil {
		panic(err)
		return "", err
	}

	return tokenString, nil
}

func (jwtAuthentication *JWTAuthentication) Authenticate(user *models.User) bool {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testing"), 10)
	testUser := models.User{
		Id:       1,
		Username: "jeyz",
		Name:     "amir",
		Password: string(hashedPassword),
	}

	return user.Username == testUser.Username && bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)) == nil
}

func getPrivateKey() *rsa.PrivateKey {
	privateKeyFile, err := os.Open(settings.Get().PrivateKeyPath)
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return privateKeyImported
}

func GetPublicKey() *rsa.PublicKey {
	publicKeyPath, err := os.Open(settings.Get().PublicKeyPath)
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyPath.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyPath)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))
	publicKeyPath.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}
