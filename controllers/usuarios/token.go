package usuarios

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/frank1995alfredo/api/config"

	database "github.com/frank1995alfredo/api/database"
	"github.com/frank1995alfredo/api/models/mantenimiento"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/twinj/uuid"
)

/***************************AUTH********************************************/

//Client ...
var Client *redis.Client

//Client ...
func init() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	Client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

// TokenDetails ... detalle del token
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuID   string
	RefreshUuID  string
	AtExpires    int64
	RtExpires    int64
}

// CreateAuth ...
func CreateAuth(userid uint64, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := Client.Set(td.AccessUuID, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := Client.Set(td.RefreshUuID, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

//CreateToken ...
func CreateToken(userid uint64) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.AccessUuID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuID = uuid.NewV4().String()

	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuID
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuID
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS512, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

//Login ...
func Login(c *gin.Context) {
	user := mantenimiento.User{}
	var input usuarios.UsuarioInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Se ha presentado un error, vuelva a ingrasar sus datos.")
		return
	}

	database.DB.Select("password").Where("usuario=?", input.Usuario).Find(&user)

	pass := user.Password
	//hash me va a devolver un true o un false
	hash := config.CheckPasswordHash(input.Password, pass)

	if hash {
		//verifica si el usuario existe y genera el token
		if err := database.DB.Where("usuario=?", input.Usuario).First(&user).Error; err == nil {
			ts, err := CreateToken(user.UsuarioID)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}

			saveErr := CreateAuth(user.UsuarioID, ts)

			if saveErr != nil {
				c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
				return
			}
			tokens := map[string]string{
				"access_token":  ts.AccessToken,
				"refresh_token": ts.RefreshToken,
			}
			c.SecureJSON(http.StatusOK, gin.H{"data": tokens})

		}
	} else {
		c.JSON(http.StatusUnauthorized, "Usuario o contraceña no validos.")
	}

}

/*******************LOGOUT*****************************/

//DeleteAuth ... elimina la autenticacion que este en ese momento
func DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := Client.Del(givenUUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

//Logout ... permite salitel del sistema
func Logout(c *gin.Context) {
	au, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "No tiene permisos para realizar esta acción.")
		return
	}
	deleted, delErr := DeleteAuth(au.AccessUuID)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "No tiene permisos para realizar esta acción.")
		return
	}
	c.JSON(http.StatusOK, "Ha salido correctamente del  sistema.")
}

//ExtractToken ... extrae el token
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

//VerifyToken ... verifica el token
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//TokenValid ...  verifica que el token es valido
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

//AccessDetails ...
type AccessDetails struct {
	AccessUuID string
	UserID     uint64
}

//ExtractTokenMetadata ... extrae los metadatos del token
func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			AccessUuID: accessUuID,
			UserID:     userID,
		}, nil
	}
	return nil, err
}

//FetchAuth ...
func FetchAuth(authD *AccessDetails) (uint64, error) {
	userid, err := Client.Get(authD.AccessUuID).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

//TokenAuthMiddleware ... seguridad en las rutas
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
