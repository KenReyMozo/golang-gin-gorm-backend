FROM golang:1.21.4

# Create the app directory
RUN mkdir /app
WORKDIR /app

# Clone the repository and build the Go application

COPY . .
RUN go build -o /app/bin .
# RUN git clone https://github.com/KenReyMozo/golang-gin-gorm-backend.git && \
#     cd golang-gin-gorm-backend && \
#     echo "PORT = 9080" > .env && \
#     echo 'DB_STRING = "host=satao.db.elephantsql.com user=xbrhqkvq password=ZfvpqzBSmiSQKiTSFNiMCIfzbf3VSjbh dbname=xbrhqkvq port=5432 sslmode=disable"' >> .env && \
#     echo 'JWT_SECRET = "qweasdzxc"' >> .env && \
#     go build -o /app/bin .

# Expose the port
EXPOSE 8080

# Set the entry point to run the application
ENTRYPOINT ["/app/bin"]
