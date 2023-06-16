FROM --platform=linux/amd64 amd64/golang:latest
LABEL authors="shivam"

COPY . /app
WORKDIR /app
RUN apt-get install -y curl
RUN ["go", "mod", "tidy"]
#ENTRYPOINT ["tail", "-f", "/dev/null"]

#CMD ["sleep", "infinity"]