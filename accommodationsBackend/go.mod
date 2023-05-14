module accommodationsBackend

go 1.20

replace accommodationsBackend/common/proto/user_service => ./common/proto/user_service

replace accommodationsBackend/common/proto/accommodation_service => ./common/proto/accommodation_service

replace accommodationsBackend/common/proto/auth_service => ./common/proto/auth_service

require github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
