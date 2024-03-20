#Get the base image
FROM golang:1.12.0-alpine3.9

#execute linux command using RUN
RUN mkdir /app

#ADD command to place main file in the above mentioned directory
ADD . /app

#go to the directory
WORKDIR /app

#RUN command to create the executable file
RUN go build -o main .

#CMD command to execute the exe file and do the application work
CMD ["/app/main"]  