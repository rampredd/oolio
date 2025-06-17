# Introduction
This is golang backend application to serve PoS APIs as per openapi specification given in this repo.

# List of APIs implemented
GET api/product to get list of products available <br />
GET api/product/{productId} to get details of one particular product <br />
POST api/order to place an order <br />
GET api/order to get list of orders placed <br />

# Begin
Clone repo into GOPATH and run below commands. <br />
cd oolio <br />
mkdir coupons <br />

Unzip coupon files and place in below folder (replace GOPATH with your local value) <br />
GOPATH/oolio/coupons/ <br />

# Run app
Run below commands to setup application and run <br />
go mod tidy <br />
go run main.go <br />

Wait for below message to appear <br />
Listening localhost:8020 <br />

# Unit tests
Run below commands in other tab to test model functions <br />
go test model/product/* <br />
go test model/order/* <br />

# API tests
Run below commands in other tab to test REST APIs <br />
go test api/order/* <br />
go test api/product/* <br />

Note: backend application should be running in other tab to test APIs <br />
