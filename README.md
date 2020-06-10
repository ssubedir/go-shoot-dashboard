# About Go Shoot Dashboard
Dashboard for [go-shoot](https://github.com/ssubedir/go-shoot)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine. See deployment for notes on how to deploy the project on a live system.

### Prerequisites
Required prerequisites
| dependence|  version | 
|---|---|
| go  |1.14.4|
| node | 14.4.0 |
| npm| 6.14.5  |


### Download

Use git clone to get your local copy 
```
git clone https://github.com/ssubedir/go-shoot-dashboard
```

## Deployment

Instructions for building and serving dashboard

### Building Dashboard

Install Node modules
```
npm install
```
Build Dashboard Angular App 
```
ng build --prod
```
Serve Dashboard through go

### Serving Dashboard

Create a .env file with the following Environmental Variables
```
# MongoDB info
DB_USER=username
DB_PASSWORD=password
DB_CLUSTER=address/cluster
```

Build / Run go-shoot (API / Dashboard)
```
go build -v main.go 
```
or
```
go run main.go 
```
Services ports
```
  Api :9000
  Dashboard :9001
```
### Docker

```
// Build
docker build -t go-shoot .
// Run
docker run -d -p 9000:9000 -p 9001:9001 go-shoot
```

## Built With

* [GO](https://golang.org/) - Programming language
* [Angular](https://angular.io/) - Front-end JS Framework

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/ssubedir/go-shoot-dashboard/blob/master/LICENSE) file for details


