package services

import (
	"crypto/rsa"
	"log"

	account "godgifu/modules/account/models"
	"godgifu/modules/auth/db"
	"godgifu/modules/auth/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

/*
In the future implement the encryption of tokens being passed to hide the JWT claims.
*/

// TokenService models the token request data passed to the service layer
type jwtService struct {
	TokenRepository            db.RedisDB
	PrivateKey                 *rsa.PrivateKey
	PublicKey                  *rsa.PublicKey
	RefreshSecretKey           string
	IDTokenExpirationSecs      int64
	RefreshTokenExpirationSecs int64
}

// ConfigTokenService will hold the repositories that are injected into
// the service layer
type ConfigTokenService struct {
	TokenRepository            db.RedisDB
	PrivateKey                 *rsa.PrivateKey
	PublicKey                  *rsa.PublicKey
	RefreshSecretKey           string
	IDTokenExpirationSecs      int64
	RefreshTokenExpirationSecs int64
}

func NewJWTService(config *ConfigTokenService) JWTService {
	return &jwtService{
		TokenRepository:            config.TokenRepository,
		PrivateKey:                 config.PrivateKey,
		PublicKey:                  config.PublicKey,
		RefreshSecretKey:           config.RefreshSecretKey,
		IDTokenExpirationSecs:      config.IDTokenExpirationSecs,
		RefreshTokenExpirationSecs: config.RefreshTokenExpirationSecs,
	}
}

// NewTokenPairFromAccount creates new tokens for Account signup and signin.
// If a previous token is included it will be removed from the tokens repository.
func (service *jwtService) NewTokenPairFromAccount(ctx echo.Context, account *account.AccountEmployee, previousTokenID string) (*models.JWTTokenPair, error) {
	log.Print("Hello previousTokenID: ", previousTokenID)
	log.Print("Hello id_code ", account.ID)
	if previousTokenID != "" {
		ctxRequest := ctx.Request().Context()
		if err := service.TokenRepository.DeleteRefreshToken(ctxRequest, account.ID.String(), previousTokenID); err != nil {
			log.Printf("Could not delete previous refresh token for account id_code %v, token ID: %v\n", account.ID, previousTokenID)

			return nil, err
		}
	}

	// No need for a repository as idToken is unrelated to any data source
	idToken, err := generateIDToken(account, service.PrivateKey, service.IDTokenExpirationSecs)
	if err != nil {
		log.Printf("Error generating new idToken for account ID Code: %v. Error: %v\n", account.ID, err.Error())
		return nil, echo.ErrInternalServerError
	}

	refreshToken, err := generateRefreshToken(account.ID, service.RefreshSecretKey, service.RefreshTokenExpirationSecs)
	if err != nil {
		log.Printf("Error generating refreshToken for account: %v. Error: %v\n", account.ID, err.Error())
		return nil, echo.ErrInternalServerError
	}

	// On account signup or signin, generate a new token for the session
	ctxRequest := ctx.Request().Context()
	if err := service.TokenRepository.SetRefreshToken(ctxRequest, account.ID.String(), refreshToken.ID.String(), refreshToken.ExpirationTime); err != nil {
		log.Printf("Error storing tokenID for account id_code: %v. Error: %v\n", account.ID, err.Error())
		return nil, echo.ErrInternalServerError
	}

	return &models.JWTTokenPair{
		JWTIDToken: models.JWTIDToken{SignedString: idToken},
		JWTRefreshToken: models.JWTRefreshToken{
			// ID:  refreshToken.ID,
			// UID: account.IDCode,
			JWT_ID:       refreshToken.ID,
			AccountID:    account.ID,
			SignedString: refreshToken.SignedString,
		},
	}, nil
}

func (service *jwtService) ValidateIDToken(tokenString string) (*models.JWTToken, error) {
	claims, err := validateIDToken(tokenString, service.PublicKey) // Signs the token using the public RSA Key
	if err != nil {
		log.Printf("Unable to validate or parse ID on Token - Error %v\n: ", err)
		return nil, echo.ErrUnauthorized
	}
	return claims, err
	// return claims.IDCode, err
}

func (service *jwtService) ValidateRefreshToken(tokenString string) (*models.JWTRefreshToken, error) {
	claims, err := validateRefreshToken(tokenString, service.RefreshSecretKey)
	if err != nil {
		log.Printf("unable to validate or parse Refresh Token for token string: %s\n%v\n", tokenString, err)
		return nil, echo.ErrUnauthorized
	}

	// JWT standard claims are a string. The model uses UUID for ID so we need to parse it.
	tokenUUID, err := uuid.Parse(claims.ID)
	if err != nil {
		log.Printf("unable to parse claims.IDCode as UUID: %s\n%v\n", claims.ID, err)
		return nil, echo.ErrUnauthorized
	}

	return &models.JWTRefreshToken{
		// ID:  tokenUUID,
		// UID: claims.IDCode,
		JWT_ID:       tokenUUID,
		AccountID:    claims.TokenID,
		SignedString: tokenString,
	}, nil
}

func (service *jwtService) Signout(ctx echo.Context, accountID uuid.UUID) error {
	ctxRequest := ctx.Request().Context()
	return service.TokenRepository.DeleteAccountRefreshTokens(ctxRequest, accountID.String())
}
