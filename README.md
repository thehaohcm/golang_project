
Author: Hao Nguyen

<h1>Introduction</h1>

Hi there,
this is my simple Golang assigment. And I'm going to show you how to setup it to run in your local machine:

First of all, please make sure your computer has been installed the latest version of Docker Engine as well as Docker Compose already.
Then clone this project from github url to your local machine and go to this root path, open the terminal (if you are using MacOS or Linux) or cmd (in case you are using Windows OS)

type: docker-compose up [-d] (-d is an option to specify that all project's containers will be run in the background), then hit Enter

The project has 3 containers running inside, there are:
1. Postgres container: contain the postgres applicaiton (postgres:14.1-alpine) with its data, which is attached by using Volume and defined in docker-compose file
2. Migrate container: defined in docker-compose by using the image "migrate/migrate:v4.15.2" from dockerhub, which is responsible for creating the latest data of project when it get started
 - NOTICE: unfortunately, it is not running properly so far, I need to look into and check it more to find a new way to run it correctly
3. Main Application (named: main): created based on golang:alpine image and its also implement in Dockerfile. You need to build it before running in your local machine. Actually, when we type "docker-compose up" for the 1st time, it will be built and run without any another manual way. But in the future, if you want to build it again because of the new change (for example), please run a command "docker-compose build".

After built and run successfully, these containers can be run (except for the Migrate container), now you can access the project's Swagger URL by web browser: http://localhost:8080/swagger/index.html

<img src="https://github.com/thehaohcm/golang_project/blob/master/asserts/swagger-screenshot.png">

Now you can check them out and test these APIs without using another API request application (such as Postman or curl)

to stop all project's containers, press Ctrl + C (if it's running in the frontground - without "-d" parameter when you started) or docker-compose stop (if it's running in the background - with "-d" parameter when you started)
