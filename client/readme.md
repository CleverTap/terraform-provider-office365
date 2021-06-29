# Client tesing

This folder contains two files client.go and client_test.go

### Client.go

    This is client library file of office365 api which contains all the functions to call all the four methods ie. CRUD

### Client_test.go

    This is a testing file that perform unit testing of client.go file containing functions to test all the four CRUD operations

## Functions for client testing

1.  ```TestClient_NewUser```  tests UserCreat function in 2 steps
    A . If user Does not exist then It will create newUser and check the respone with expected response
    B . If User exist then it wil show error  

2. ```TestClient_GetUser``` tests GetUser fucntion in 2 steps
    A.If user exist then it will compare the response with expected response
    B.If User exist then it wil show error  

3. ```TestClient_UpdateUser``` tests userUpdate in 2 steps
   A. If user exist then it will update the user and compare the response with expected respone
   B. If User exist then it wil show error  
4.  ```TestClient_DeleteItem``` tests deletedUser in 2 steps
  A. If user exist then it will delete the user
  B. If User exist then it wil show error  

### Steps to perform testing of client.go

1. Open client folder in terminal 
2. Run  
  ```
  OFFICE365_CLIENT_ID=<clientID> OFFICE365_CLIENT_SECRET=<clientSceret> OFFICE365_TENANT_ID=<tenanatID> go test -v
  ```
3. To check coverage run go test -cover