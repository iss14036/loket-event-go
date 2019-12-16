# List of End Point and Assumption of Loket Event


## List of End Point

Don't forget in the header put :

```
Content-Type : application/json
Authorization : {{Your_Token}}
```

### 1. Customer

#### 1.1 Create Customer

In create customer, you can use this url with method `POST` to create a customer :
```
http://localhost:8000/customer
```

And this is the body request :
```
{
    "name": "dani",
    "email": "dani@dani.com",
    "phone": "084153"
}
```

#### 1.2 Get List Customer

In Get list customer, you can use this url with method `GET` :
```
http://localhost:8000/customer
```

#### 1.3 Get Customer

In Get list customer, you can use this url with method `GET` :
```
http://localhost:8000/customer/{id}
```
