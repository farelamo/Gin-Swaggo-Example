# Gin Golang + Swagger Example

This is my learning about using Swagger with **the Gin Framework and Go 1.19 version, as well as my Linux Fedora system** to automatically generate API documentation. According to the official Swagger website, it's only necessary to use **annotation** on directive comments and structs of Golang type for live changes. I was trying to learn this because one of the most popular API documentation tools, **POSTMAN** requires enterprise users to pay between $14 and $20 per month. For any further information, we can visit the documentation of Go Swagger at https://github.com/go-swagger/go-swagger

# Overview
 ## 1. First Page
 ![First Page](image.png)

 ## 2. With Params Example
 ![With Params](image-1.png)

 ## 3. With Request Body
 ![With Request Body](image-2.png)

# Pre-requisite
1. Install the swag-CLI
    ```
    go install github.com/swaggo/swag/cmd/swag@latest
    ```

2. Then, You can init the swagger docs, use it from environment variables or use alias
 - For Linux
    - Use Env
        ```
        $HOME/go/bin/swag init
        ```
    - Use Alias (After run this, you can simply call like : swag init)
        ```
        alias swag="$HOME/go/bin/swag"
        ```
 - For Windows 
    For windows you can simply add to environment variables (search on internet if you have any problem) :D

3. get all dependencies required
    ```
    go mod tidy
    ```
4. 🚀 Run the project !
    ```
    go run .
    ```
5. Go to the swagger page
    ```
    localhost:3000/swagger/index.html
    ```
6. If you have problems after running ***go run .***, Well you're not alone. Its because of the openapi library I use is compatible with go 1.19 version as well as in description. If you are higher, then you need to follow next step. (6 and 7 step is version that compatible with my go version, to search what compatible with yours, go to the github.com/go-openapi/swag and search which the tags and release date you want)

7. Go-OpenAPI/Swag install
    ```
    go get -u github.com/go-openapi/swag@v0.20.0
    ```

8. Go-OpenAPI/Spec Install
    ```
    go get github.com/go-openapi/spec@v0.20.11
    ```
9. 🚀 Try Run Again !
    ```
    go run .
    ```