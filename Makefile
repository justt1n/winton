init-dependencies:
	@echo "Installing dependencies..."
	@go get -u github.com/gin-gonic/gin
	@go get -u github.com/jinzhu/gorm
	@go get -u github.com/dgrijalva/jwt-go
	@go get -u github.com/joho/godotenv
	@go get -u golang.org/x/crypto
	@echo "Dependencies installed."