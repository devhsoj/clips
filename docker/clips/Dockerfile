FROM alpine:latest

WORKDIR /opt/clips

RUN apk add --update --no-cache nodejs npm go

COPY . .

RUN npm install
RUN npm run build

EXPOSE 3000

CMD ["./build/clips"]