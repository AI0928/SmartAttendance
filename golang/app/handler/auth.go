package handler

import (
	"app/model"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"golang.org/x/crypto/bcrypt"
)

func (c *jwtCustomClaims) Valid() error {
    if time.Now().Unix() > c.ExpiresAt.Unix() {
        return fmt.Errorf("Token is expired")
    }
    return nil
}

type jwtCustomClaims struct {
    UID  string    `json:"uid"`
    Name string `json:"name"`
    jwt.RegisteredClaims
}

var signingKey = []byte("secret")

var Config = middleware.JWTConfig{
    Claims: &jwtCustomClaims{},
    SigningKey: signingKey,
}

func encryptPassword(password string) (string, error) {
	// パスワードの文字列をハッシュ化する
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func compareHashPassword(hashedPassword, requestPassword string) error {
	// パスワードの文字列をハッシュ化して、既に登録されているハッシュ化したパスワードと比較します
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(requestPassword)); err != nil {
		return err
	}
	return nil
}

func Signup(c echo.Context) error {
    user := new(model.User)
    if err := c.Bind(user); err != nil {
        return err
    }

    fmt.Println(user.Id, user.Name, user.Email, user.Password)
    if user.Id == "" || user.Password == "" || user.Email == ""{
        return &echo.HTTPError{
            Code:    http.StatusBadRequest,
            Message: "invalid Id or password, Email",
        }
    }

	u := model.FindUser(&model.User{Id: user.Id})

    if  u.Name != "" {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "id already exists",
        }
    }

    u = model.FindUser(&model.User{Email: user.Email})
    if  u.Name != "" {
        return &echo.HTTPError{
            Code:    http.StatusConflict,
            Message: "email already exists",
        }
    }

	var err error
	user.Password, err = encryptPassword(user.Password)
	if  err != nil {
        return &echo.HTTPError{
            Code:  http.StatusInternalServerError,
            Message: "Error while hashing password",
        }
    }
    model.CreateUser(user)
    user.Password = ""

    return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
    u := new(model.User)
    if err := c.Bind(u); err != nil {
        return err
    }

    fmt.Println("input id", u.Id)
    fmt.Println("input email", u.Email)
    fmt.Println("input pass", u.Password)
    user := model.FindUser(&model.User{Email: u.Email})
    fmt.Println("search", user.Email, user.Password)
	password_err := compareHashPassword(user.Password, u.Password)
    if user.Name == "" || password_err != nil {
        return &echo.HTTPError{
            Code:    http.StatusUnauthorized,
            Message: "invalid id or password",
        }
    }

    claims := &jwtCustomClaims{
        user.Id,
        user.Name,
        jwt.RegisteredClaims{
            ExpiresAt:  jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString(signingKey)
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, map[string]string{
        "token": t,
    })
}

func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
    if claims, ok := user.Claims.(*jwtCustomClaims); ok {
		return claims.UID
	} else {
		// handle error here
		return ""
	}
}