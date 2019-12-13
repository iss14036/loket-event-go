# Loket Event API

## Problem Statement
Loket.com is a white-label event ticketing company, their main business is selling event ticket directly from promotor to customer or via affiliated partner. Loket will have many event at one time, each event will be held on specific location at specific date, also each event will have one or more ticket that can be purchased by customer. Your job here is to create database schedule to accomodate storing event data and export functionalities through HTTP API.

For further requirement, you can check [here](docs/REQUIREMENT.md)

## Dependencies

#### 1. Golang
In this project, I use Mac as OS which is using `brew` as installation. Follow this step to install Go in mac :
```
brew update
brew install golang
```
Add environment variable

We'll add some environment variables into shell config. One of does files located at your home directory bash_profile, bashrc or .zshrc (for Oh My Zsh Army)
```
export GOPATH=/Users/$USER/go
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```
Basically, I use Go Module for dependencies manager. So you don't have to worry about project placement

#### 2. Mysql Database
You can install mysql with this command :
```
brew install mysql
```

Start mysql :
```
brew services start mysql
```

#### 3. DBMate as Migration
I use DBMate as migration, for installing you can type this following command : 
```
brew install dbmate
```

#### 3. Redis
`To be added`

## How to use
#### 1. Environment variable
Copy `.env.example` and paste into `.env` and fill each variables as you need

#### 2. Run the migration
Follow this command for running the migration :
```
dbmate up
``` 
For further command, you can open [HERE](https://github.com/amacneil/dbmate).

#### 3. Run the application
```
go run main.go
```