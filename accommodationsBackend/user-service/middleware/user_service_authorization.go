package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
	"sync"
	"time"
)

type AuthInterceptor struct {
	//generatedToken string
}

var GlobalTokenStore sync.Map

type TokenClaimsContextKey struct{}

func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
}

const (
	jWTPrivateToken = "SecrteTokenSecrteToken" // privatni kljuc
)
const (
	userServiceKey = "userServiceKey" // privatni kljuc
)

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		err := interceptor.authenticate(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) authenticate(ctx context.Context, fullMethod string) error {

	if fullMethod == "/UserService/Register" {
		/*claims := &JwtClaimsNoAuthInterservice{
			Id:          primitive.NewObjectID(),
			ServiceRole: "user-service",
		}
		expirationTime := time.Now().Add(4 * time.Hour)
		token, err := GenerateNoAuthToken(claims, expirationTime)
		if err != nil {

			return status.Errorf(codes.Internal, "Failed to generate no-auth token: %v", err)
		}
		fmt.Println("token tirng u autenticat emetodi ", token)
		GlobalTokenStore.Store("t", token)
		return nil
		*/
		expirationTime := time.Now().Add(4 * time.Hour)
		token, err := GenerateNoAuthToken(expirationTime)
		if err != nil {
			return status.Errorf(codes.Internal, "Failed to generate no-auth token: %v", err)
		}
		fmt.Println("Token string in authenticate method: ", token)
		GlobalTokenStore.Store("t", token)
		return nil
	}
	/* md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is missing")
	}

	authorizationHeaders, ok := md["authorization"]
	if !ok || len(authorizationHeaders) != 1 {
		return status.Errorf(codes.Unauthenticated, "authorization token is missing")
	}

	tokenString := strings.TrimPrefix(authorizationHeaders[0], "Bearer ")
	if tokenString == "" {
		return status.Errorf(codes.Unauthenticated, "authorization token is missing")
	}
	*/
	tokenString, err := interceptor.extractTokenFromMetadata(ctx)
	if err != nil {
		return err
	}

	_, err = interceptor.parseClaims(tokenString, ctx)
	if err != nil {
		return err
	}

	return nil
}

func (interceptor *AuthInterceptor) extractTokenFromMetadata(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata is missing")
	}

	authorizationHeaders, ok := md["authorization"]
	if !ok || len(authorizationHeaders) != 1 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is missing")
	}

	tokenString := strings.TrimPrefix(authorizationHeaders[0], "Bearer ")
	if tokenString == "" {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is missing")
	}

	return tokenString, nil
}

func (interceptor *AuthInterceptor) parseClaims(tokenString string, ctx context.Context) (context.Context, error) {
	claims := &JwtClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jWTPrivateToken), nil
	})

	if err != nil {
		return ctx, err
	}
	fmt.Printf("ID: %s\n", claims.Id.Hex())
	fmt.Printf("Name: %s\n", claims.Name)
	fmt.Printf("Username: %s\n", claims.Username)
	fmt.Printf("Role: %s\n", claims.Role)

	expirationTime := time.Now().Add(4 * time.Hour)
	newToken, tokenErr := GenerateToken(claims, expirationTime)

	if tokenErr != nil {
		return ctx, tokenErr
	}
	GlobalTokenStore.Store("t", newToken)
	return context.WithValue(ctx, TokenClaimsContextKey{}, claims), nil
}

func GenerateToken(claims *JwtClaims, expirationTime time.Time) (string, error) {
	interserviceClaims := &JwtClaimsInterservice{
		Id:   claims.Id,
		Role: claims.Role,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
			Issuer:    "user-service",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, interserviceClaims)
	tokenString, err := token.SignedString([]byte(userServiceKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (interceptor *AuthInterceptor) AddTokenToMetadata(ctx context.Context) context.Context {
	md, _ := metadata.FromOutgoingContext(ctx)
	newMD := md.Copy()
	token, ok := GlobalTokenStore.Load("t")
	if !ok {
		return ctx
	}
	tokenString, valid := token.(string)
	if !valid {
		return ctx
	}
	newMD.Set("authorization", "Bearer "+tokenString)
	return metadata.NewOutgoingContext(ctx, newMD)
}

/*
func GenerateNoAuthToken(claims *JwtClaimsNoAuthInterservice, expirationTime time.Time) (string, error) {
	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jWTPrivateToken))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}*/

func GenerateNoAuthToken(expirationTime time.Time) (string, error) {
	claims := &jwt.StandardClaims{
		Id:        primitive.NewObjectID().Hex(),
		Issuer:    "user-service",
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().UTC().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jWTPrivateToken))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

type JwtClaims struct {
	Id       primitive.ObjectID `bson:"id,omitempty" json:"id"`
	Name     string             `bson:"name"     json:"name"`
	Username string             `bson:"username" json:"username"`
	Role     string             `bson:"role" json:"role"`

	jwt.StandardClaims
}

type JwtClaimsInterservice struct {
	Id primitive.ObjectID `bson:"id,omitempty" json:"id"`

	Role string `bson:"role" json:"role"`

	jwt.StandardClaims
}

type JwtClaimsNoAuthInterservice struct {
	Id          primitive.ObjectID `bson:"id,omitempty" json:"id"`
	ServiceRole string             `bson:"serviceRole" json:"serviceRole"`

	jwt.StandardClaims
}

func (claims JwtClaims) Valid() error {
	var now = time.Now().UTC().Unix()
	if claims.VerifyExpiresAt(now, true) {
		return nil
	}
	return fmt.Errorf("Token is invalid")
}
