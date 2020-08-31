FROM golang

WORKDIR /shr

COPY . .

RUN go install -v ./...

CMD ["shr"]