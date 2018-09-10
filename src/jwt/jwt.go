import (
	"context"
	"crypto/rsa"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

// AppClaims provides custom claim for JWT
type AppClaims struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// using asymmetric crypto/RSA keys
// location of private/public key files
const (
	// openssl genrsa -out app.rsa 1024
	privKeyPath = "keys/app.rsa"
	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "keys/app.rsa.pub"
)

// Private key for signing and public key for verification
var (
	//verifyKey, signKey []byte
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// InitRSAKeys reads the key files to be used for JWT
func InitRSAKeys() {

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
}

// GenerateJWT generates a new JWT token for authenticated user.
func GenerateJWT(name, role string) (string, error) {
	// Create the Claims
	claims := AppClaims{
		UserName: name,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			Issuer:    "admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func main(){
	
}