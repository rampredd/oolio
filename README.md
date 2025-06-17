# Introduction
This is golang backend application to serve PoS APIs as per openapi specification given in this repo.

# List of APIs implemented
GET api/product to get list of products available
GET api/product/{productId} to get details of one particular product
POST api/order to place an order
GET api/order to get list of orders placed

# Begin
Clone repo into GOPATH and run below commands.
cd oolio
mkdir coupons

Unzip coupon files and place in below folder (replace GOPATH your local value)
<GOPATH>/oolio/coupons/

# Run app
Run below commands to setup application and run
go mod tidy
go run main.go

Wait for below message to appear
Listening localhost:8020

# Unit tests
Run below commands in other tab to test model functions
go test model/product/*
go test model/order/*

# API tests
Run below commands in other tab to test REST APIs
go test api/order/*
go test api/product/*
