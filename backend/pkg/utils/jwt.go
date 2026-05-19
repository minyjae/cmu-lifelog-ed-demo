package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email string
	jwt.RegisteredClaims
}

// ฟังก์ชันสำหรับสร้าง JWT TOken
var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

type JWTClaims struct {
	Name               string `json:"cmuitaccount_name"`
	Email              string `json:"cmuitaccount"`
	PreNameID          string `json:"prename_id"`
	PreNameTH          string `json:"prename_th"`
	PreNameEN          string `json:"prename_en"`
	FirstNameTH        string `json:"firstname_th"`
	FirstNameEN        string `json:"firstname_en"`
	LastNameTH         string `json:"lastname_th"`
	LastNameEN         string `json:"lastname_en"`
	OrganizationCode   string `json:"organization_code"`
	OrganizationNameTH string `json:"organization_name_th"`
	OrganizationNameEN string `json:"organization_name_en"`
	ITAccountTypeID    string `json:"itaccounttype_id"`
	ITAccountTypeTH    string `json:"itaccounttype_th"`
	ITAccountTypeEN    string `json:"itaccounttype_en"`
	jwt.RegisteredClaims
}

func GenerateJWT(name, email, prename_id, prename_th, prename_en, firstname_th,
	firstname_en, lastname_th, lastname_en, organizationcode, organizationname_th, organizationname_en, itaccounttype_id, itaccounttype_th, itaccounttype_en string) (string, error) {
	// fmt.Printf("In [jwt_service] : ITAccountTypeID = %v\n", itaccounttype_id)
	// fmt.Printf("In [jwt_service] : ITAccountTypeTH = %v\n", itaccounttype_th)
	// fmt.Printf("In [jwt_service] : ITAccountTypeEN = %v\n", itaccounttype_en)
	claims := JWTClaims{
		Name:               name,
		Email:              email,
		PreNameID:          prename_id,
		PreNameTH:          prename_th,
		PreNameEN:          prename_en,
		FirstNameTH:        firstname_th,
		FirstNameEN:        firstname_en,
		LastNameTH:         lastname_th,
		LastNameEN:         lastname_en,
		OrganizationCode:   organizationcode,
		OrganizationNameTH: organizationname_th,
		OrganizationNameEN: organizationname_en,
		ITAccountTypeID:    itaccounttype_id,
		ITAccountTypeTH:    itaccounttype_th,
		ITAccountTypeEN:    itaccounttype_en,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
