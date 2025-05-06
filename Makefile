# Define the directories and microservices for each path
APP_MICROSERVICES := authentication watchlist orders portfolio alerts profile payments stocks search global-search scanners corporate-action mutual-funds
ADMIN_APP_MICROSERVICES := authentication orders search watchlist stocks portfolio profile global-search
DIR_PATHS := ./src/app ./src/admin-app
GO_VERSION := 1.20

# Generate Swagger for each microservice
generate-swagger:
	@for path in $(DIR_PATHS); do \
		if [ "$$path" = "./src/app" ]; then \
			for service in $(APP_MICROSERVICES); do \
				echo "Generating Swagger for $$service in $$path"; \
				(cd $$path/$$service && swag init -g main.go); \
			done \
		else \
			for service in $(ADMIN_APP_MICROSERVICES); do \
				echo "Generating Swagger for $$service in $$path"; \
				(cd $$path/$$service && swag init -g main.go); \
			done \
		fi \
	done

# Run go mod tidy for each microservice
go-mod-tidy:
	@for path in $(DIR_PATHS); do \
		if [ "$$path" = "./src/app" ]; then \
			for service in $(APP_MICROSERVICES); do \
				echo "Running GO mod tidy for $$service in $$path"; \
				(cd $$path/$$service && go mod tidy); \
			done \
		else \
			for service in $(ADMIN_APP_MICROSERVICES); do \
				echo "Running GO mod tidy for $$service in $$path"; \
				(cd $$path/$$service && go mod tidy); \
			done \
		fi \
	done

# Run test cases for each microservice
test-all:
	@for path in $(DIR_PATHS); do \
		if [ "$$path" = "./src/app" ]; then \
			for service in $(APP_MICROSERVICES); do \
				echo "Running test cases for $$service in $$path"; \
				(cd $$path/$$service && go test ./tests/... -coverpkg=./... -coverprofile $$service.out -covermode count); \
			done \
		else \
			for service in $(ADMIN_APP_MICROSERVICES); do \
				echo "Running test cases for $$service in $$path"; \
				(cd $$path/$$service && go test ./tests/... -coverpkg=./... -coverprofile $$service.out -covermode count); \
			done \
		fi \
	done

test-admin-all:
	@for path in $(DIR_PATHS); do \
		if [ "$$path" = "./src/admin-app" ]; then \
			for service in $(ADMIN_APP_MICROSERVICES); do \
				echo "Running test cases for $$service in $$path"; \
				(cd $$path/$$service && go test ./tests/... -coverpkg=./... -coverprofile $$service.out -covermode count); \
			done \
		fi \
	done

# Copy api.yml in test setup enviroment
copy-api-yml:
	cp ./src/configs/apis.yml ./src/setupTest/testConfigs/apis.yml

# Generate swagger files and run GO mod tidy
all: generate-swagger go-mod-tidy copy-api-yml test-all

APP_DIR_PATH := ./src/app
ADMIN_DIR_PATH := ./src/admin-app

# Services in ./src/app
authentication: 
	cd $(APP_DIR_PATH)/authentication && go mod tidy

watchlist: 
	cd $(APP_DIR_PATH)/watchlist && go mod tidy
	cd $(APP_DIR_PATH)/watchlist && go test ./tests/... -coverpkg=./... -coverprofile watchlist.out -covermode count

orders: 
	cd $(APP_DIR_PATH)/orders && go mod tidy
	cd $(APP_DIR_PATH)/orders && go test ./tests/... -coverpkg=./... -coverprofile orders.out -covermode count

portfolio: 
	cd $(APP_DIR_PATH)/portfolio && go mod tidy
	cd $(APP_DIR_PATH)/portfolio && go test ./tests/... -coverpkg=./... -coverprofile portfolio.out -covermode count

alerts: 
	cd $(APP_DIR_PATH)/alerts && go mod tidy
	cd $(APP_DIR_PATH)/alerts && go test ./tests/... -coverpkg=./... -coverprofile alerts.out -covermode count

profile: 
	cd $(APP_DIR_PATH)/profile && go mod tidy
	cd $(APP_DIR_PATH)/profile && go test ./tests/... -coverpkg=./... -coverprofile profile.out -covermode count

payments: 
	cd $(APP_DIR_PATH)/payments && go mod tidy
	cd $(APP_DIR_PATH)/payments && go test ./tests/... -coverpkg=./... -coverprofile payments.out -covermode count

stocks: 
	cd $(APP_DIR_PATH)/stocks && go mod tidy
	cd $(APP_DIR_PATH)/stocks && go test ./tests/... -coverpkg=./... -coverprofile stocks.out -covermode count

search: 
	cd $(APP_DIR_PATH)/search && go mod tidy
	cd $(APP_DIR_PATH)/search && go test ./tests/... -coverpkg=./... -coverprofile search.out -covermode count

global-search: 
	cd $(APP_DIR_PATH)/global-search && go mod tidy
	cd $(APP_DIR_PATH)/global-search && go test ./tests/... -coverpkg=./... -coverprofile global-search.out -covermode count

scanners: 
	cd $(APP_DIR_PATH)/scanners && go mod tidy
	cd $(APP_DIR_PATH)/scanners && go test ./tests/... -coverpkg=./... -coverprofile scanners.out -covermode count

corporate-action: 
	cd $(APP_DIR_PATH)/corporate-action && go mod tidy
	cd $(APP_DIR_PATH)/corporate-action && go test ./tests/... -coverpkg=./... -coverprofile corporate-action.out -covermode count

mutual-funds: 
	cd $(APP_DIR_PATH)/mutual-funds && go mod tidy
	cd $(APP_DIR_PATH)/mutual-funds && go test ./tests/... -coverpkg=./... -coverprofile mutual-funds.out -covermode count

# Services in ./src/admin-app
admin-authentication: 
	cd $(ADMIN_DIR_PATH)/authentication && go mod tidy

admin-orders: 
	cd $(ADMIN_DIR_PATH)/orders && go mod tidy
	cd $(ADMIN_DIR_PATH)/orders && go test ./tests/... -coverpkg=./... -coverprofile orders.out -covermode count

admin-search: 
	cd $(ADMIN_DIR_PATH)/search && go mod tidy
	cd $(ADMIN_DIR_PATH)/search && go test ./tests/... -coverpkg=./... -coverprofile search.out -covermode count

admin-watchlist: 
	cd $(ADMIN_DIR_PATH)/watchlist && go mod tidy
	cd $(ADMIN_DIR_PATH)/watchlist && go test ./tests/... -coverpkg=./... -coverprofile watchlist.out -covermode count

admin-stocks:
	cd $(ADMIN_DIR_PATH)/stocks && go mod tidy
	cd $(ADMIN_DIR_PATH)/stocks && go test ./tests/... -coverpkg=./... -coverprofile stocks.out -covermode count

admin-profile: 
	cd $(ADMIN_DIR_PATH)/profile && go mod tidy
	cd $(ADMIN_DIR_PATH)/profile && go test ./tests/... -coverpkg=./... -coverprofile profile.out -covermode count

admin-global-search:
	cd $(ADMIN_DIR_PATH)/global-search && go mod tidy
	cd $(ADMIN_DIR_PATH)/global-search && go test ./tests/... -coverpkg=./... -coverprofile global-search.out -covermode count

create_service:
	./scripts/createMicroservice.sh $(servicename) $(port) $(GO_VERSION)

.PHONY: create_service
