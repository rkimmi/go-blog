#### Go Blog
Currently serves /photos-blog within rkimmi.github.io. More to come ~

### Getting Started

#### Install packages
run `go mod tidy`

#### Start up application
run `go build`
run `./blog`

### Deploy
From CLI:
run `flyctl deploy`

#### Deployment troubleshooting

##### Check fly.io status' and logs

run `flyctl status` to see the status of all processes
run `flyctl machine status <ID> to see details status' of a single process.

##### Run docker file
NOTE: WIP there is a bug with .env files not loading when running docker locally.
Run the deployment docker file locally to see any errors thrown during build:
docker build -t blog .
docker run -p 8080:8080 blog

### Project structure
./cloudinary - contains code related to interacting with the cloudinary sdk for retrieving images.
./photos-blog - contains handlers and types for getting photo data 

### Notes
run `go clean -cache` if you're seeing a stale state of the app running. 
