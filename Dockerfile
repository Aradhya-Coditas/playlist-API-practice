# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.22.3


# Parameterise at the microservice level and env variables
ARG MICROSERVICE
ARG POSTGRES_HOST
ARG POSTGRES_PORT
ARG POSTGRES_USER
ARG POSTGRES_PASSWORD
ARG POSTGRES_DB_NAME
ARG COCKROACHDB_HOST
ARG COCKROACHDB_PORT
ARG COCKROACHDB_USER
ARG COCKROACHDB_PASSWORD
ARG COCKROACHDB_NAME
ARG BANNER_URL_PREFIX
ARG OLTP_HTTP_ENDPOINT_URL
ARG APP_ENVIRONMENT
ARG REST_BASE_URL
ARG SCANNER_BASE_URL
ARG GLOBAL_SEARCH_BASE_URL
ARG IPO_BASE_URL
ARG USE_MOCKS
ARG USE_DB_MOCKS
ARG IS_MONOLITH
ARG ENABLE_UIBFF_ENCRYPT_DECRYPT
ARG ENABLE_OPEN_TELEMETRY
ARG ENABLE_RATE_LIMIT
ARG CMOTS_TOKEN
ARG DEFAULT_MW_SCRIPTS_COUNT
ARG OTP_EXPIRY_IN_MINUTES
ARG OTP_LENGTH 
ARG ALLOWED_MAXIMUM_LOGIN_COUNT
ARG PASSWORD_VALIDATION_REGEX
ARG RESEND_OTP_TIMEOUT
ARG ADMIN_USER_NAME
ARG ADMIN_PASSWORD
ARG BASKET_LIMIT
ARG BASKETNAMEVALIDATIONREGEX
ARG MAXIMUM_BASKET_ORDERS_LIMIT
ARG WATCHLIST_LIMIT
ARG WATCHLIST_SCRIP_LIMIT_PER_WATCHLIST
ARG OPENSEARCH_HOST
ARG OPENSEARCH_PASSWORD
ARG OPENSEARCH_PORT
ARG OPENSEARCH_USER
ARG OPENSEARCH_SSL_MODE
ARG OPENSEARCH_CACERT_PATH
ARG TIME_ZONE
ARG SCRIP_MASTER_TIME_INTERVAL_IN_HOURS
ARG SCRIP_MASTER_BATCH_SIZE
ARG CMOTS_YEAR_MONTH_REGEX
ARG CMOTS_SCHEDULER_INTERVAL_IN_DAYS
ARG CMOTS_SCHEDULER_RUN_TIME
ARG CMOTS_MAX_RETRY_COUNT
# Set Nest related env variables
ARG MML_LOC_BROK_ADDR
ARG MML_DMN_SRVR_ADDR
ARG MML_DS_FO_ADDR
ARG MML_LIC_SRVR_ADDR
ARG ADMIN_NAME
ARG INT_DD_NAME
ARG BCAST_DD_NAME
ARG INT_REQ_DD_NAME
ARG RMS_GET_PRSNT_DD_NAME
ARG TOUCHLINE_DD_NAME
ARG RMS_DD_NAME
ARG CLIENT_LEDGER_URL
ARG GLOBAL_PL_URL
ARG ENABLE_IPO_CACHE
ARG HOLDINGS_URL
ARG CLIENT_FA_SUMMARY_URL
ARG APPLICATION_NAME

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy SSL certificate and key to the /app directory
COPY src/utils/sslCertificate/webSocket.crt /app/
COPY src/utils/sslCertificate/webSocket.key /app/

# Read the content of SSL certificate and key into environment variables
ENV SSL_CERTIFICATE_CRT=/app/webSocket.crt
ENV SSL_CERTIFICATE_KEY=/app/webSocket.key
ENV POSTGRES_HOST=$POSTGRES_HOST 
ENV POSTGRES_PORT=$POSTGRES_PORT
ENV POSTGRES_USER=$POSTGRES_USER
ENV POSTGRES_PASSWORD=$POSTGRES_PASSWORD
ENV POSTGRES_DB_NAME=$POSTGRES_DB_NAME
ENV COCKROACHDB_HOST=$COCKROACHDB_HOST 
ENV COCKROACHDB_PORT=$COCKROACHDB_PORT
ENV COCKROACHDB_USER=$COCKROACHDB_USER
ENV COCKROACHDB_PASSWORD=$COCKROACHDB_PASSWORD
ENV COCKROACHDB_NAME=$COCKROACHDB_NAME
ENV BANNER_URL_PREFIX=$BANNER_URL_PREFIX
ENV OLTP_HTTP_ENDPOINT_URL=$OLTP_HTTP_ENDPOINT_URL
ENV APP_ENVIRONMENT=$APP_ENVIRONMENT
ENV REST_BASE_URL=$REST_BASE_URL
ENV SCANNER_BASE_URL=$SCANNER_BASE_URL
ENV GLOBAL_SEARCH_BASE_URL=$GLOBAL_SEARCH_BASE_URL
ENV IPO_BASE_URL=$IPO_BASE_URL
ENV USE_MOCKS=$USE_MOCKS
ENV USE_DB_MOCKS=$USE_DB_MOCKS
ENV IS_MONOLITH=$IS_MONOLITH
ENV ENABLE_UIBFF_ENCRYPT_DECRYPT=$ENABLE_UIBFF_ENCRYPT_DECRYPT
ENV ENABLE_OPEN_TELEMETRY=$ENABLE_OPEN_TELEMETRY
ENV ENABLE_RATE_LIMIT=$ENABLE_RATE_LIMIT
ENV DEFAULT_MW_SCRIPTS_COUNT=$DEFAULT_MW_SCRIPTS_COUNT
ENV OTP_EXPIRY_IN_MINUTES=$OTP_EXPIRY_IN_MINUTES
ENV CMOTS_TOKEN=$CMOTS_TOKEN
ENV OTP_LENGTH=$OTP_LENGTH
ENV ALLOWED_MAXIMUM_LOGIN_COUNT=$ALLOWED_MAXIMUM_LOGIN_COUNT
ENV PASSWORD_VALIDATION_REGEX=$PASSWORD_VALIDATION_REGEX
ENV RESEND_OTP_TIMEOUT=$RESEND_OTP_TIMEOUT
ENV ADMIN_USER_NAME=$ADMIN_USER_NAME
ENV ADMIN_PASSWORD=$ADMIN_PASSWORD
ENV BASKET_LIMIT=$BASKET_LIMIT
ENV BASKETNAMEVALIDATIONREGEX=$BASKETNAMEVALIDATIONREGEX
ENV MAXIMUM_BASKET_ORDERS_LIMIT=$MAXIMUM_BASKET_ORDERS_LIMIT
ENV WATCHLIST_LIMIT=$WATCHLIST_LIMIT
ENV WATCHLIST_SCRIP_LIMIT_PER_WATCHLIST=$WATCHLIST_SCRIP_LIMIT_PER_WATCHLIST
ENV CLIENT_LEDGER_URL=$CLIENT_LEDGER_URL
ENV GLOBAL_PL_URL=$GLOBAL_PL_URL
ENV ENABLE_IPO_CACHE=$ENABLE_IPO_CACHE
ENV HOLDINGS_URL=$HOLDINGS_URL
ENV CLIENT_FA_SUMMARY_URL=$CLIENT_FA_SUMMARY_URL
ENV OPENSEARCH_HOST=$OPENSEARCH_HOST
ENV OPENSEARCH_USER=$OPENSEARCH_USER
ENV OPENSEARCH_PORT=$OPENSEARCH_PORT
ENV OPENSEARCH_PASSWORD=$OPENSEARCH_PASSWORD
ENV OPENSEARCH_CACERT_PATH=$OPENSEARCH_CACERT_PATH
ENV OPENSEARCH_SSL_MODE=$OPENSEARCH_SSL_MODE
ENV TIME_ZONE=$TIME_ZONE
ENV SCRIP_MASTER_TIME_INTERVAL_IN_HOURS=$SCRIP_MASTER_TIME_INTERVAL_IN_HOURS
ENV SCRIP_MASTER_BATCH_SIZE=$SCRIP_MASTER_BATCH_SIZE
ENV CMOTS_YEAR_MONTH_REGEX=$CMOTS_YEAR_MONTH_REGEX
ENV CMOTS_SCHEDULER_INTERVAL_IN_DAYS=$CMOTS_SCHEDULER_INTERVAL_IN_DAYS
ENV CMOTS_SCHEDULER_RUN_TIME=$CMOTS_SCHEDULER_RUN_TIME
ENV CMOTS_MAX_RETRY_COUNT=$CMOTS_MAX_RETRY_COUNT

# Set Nest related env variables
ENV MML_LOC_BROK_ADDR=$MML_LOC_BROK_ADDR
ENV MML_DMN_SRVR_ADDR=$MML_DMN_SRVR_ADDR
ENV MML_DS_FO_ADDR=$MML_DS_FO_ADDR
ENV MML_LIC_SRVR_ADDR=$MML_LIC_SRVR_ADDR
ENV ADMIN_NAME=$ADMIN_NAME
ENV INT_DD_NAME=$INT_DD_NAME
ENV BCAST_DD_NAME=$BCAST_DD_NAME
ENV INT_REQ_DD_NAME=$INT_REQ_DD_NAME
ENV RMS_GET_PRSNT_DD_NAME=$RMS_GET_PRSNT_DD_NAME
ENV TOUCHLINE_DD_NAME=$TOUCHLINE_DD_NAME
ENV RMS_DD_NAME=$RMS_DD_NAME
ENV APPLICATION_NAME=$APPLICATION_NAME
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
WORKDIR /app/src
RUN mkdir encryptionKeys


WORKDIR /app

# Remove Workspace file
RUN rm -f *.work*

# Automigrate DB
RUN go run main.go

# Remove other microservices' folders
RUN find /app/src/app/ -mindepth 1 -maxdepth 1 -type d ! -name $MICROSERVICE -exec rm -r {} \;

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
WORKDIR /app/src/app/$MICROSERVICE
RUN go install .

# SWAG installation
RUN curl -o /usr/local/bin/Swag https://github.com/linuxserver/docker-Swag/releases/latest/download/Swag && chmod +x /usr/local/bin/Swag

# Build the Go app
RUN go build -o main .

# Swagger update
RUN /usr/local/bin/Swag init -g main.go

# Run test cases
#WORKDIR /app
#RUN make $MICROSERVICE

# Copy Dockerfile file to store it in the pod
ADD ./Dockerfile /tmp/

# Changing working directory
#WORKDIR /app/src/app/$MICROSERVICE

# Command to run the executable
CMD ["./main"]