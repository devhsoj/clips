# Project Clips

Project Clips is a simple video hosting platform that enables people to share and view clips.

## Getting Started

### Using Docker

**Note:** Make sure you have a valid .env that works with docker (See [.env.docker.example](https://github.com/devhsoj/clips/blob/master/.env.docker.example))
```bash
git clone https://github.com/devhsoj/clips
cd clips/
docker-compose -f .\docker\docker-compose.yml up -d --build
```

---

### Without Docker

**Requirements:** [Go](https://go.dev/dl/) - [npm](https://nodejs.org/en/download/) - [PostgreSQL](https://www.postgresql.org/download/)

```bash
git clone https://github.com/devhsoj/clips
cd clips/
npm install
npm run build
```

**Postgres Setup:**
```postgresql
CREATE DATABASE clips;
\c clips;
CREATE EXTENSION pgcrypto; -- Used for generating UUIDs
```

### Starting

**Note:** Make sure you have run the command `npm run build` before-hand, and have a valid `.env`

#### Run with npm
```bash
npm start # or with pm2: pm2 start npm --name clips -- start
```
or

#### Run from built executable
```bash
./build/clips # or ./build/clips.exe
```
---
## Development

#### Starting
```bash
# uses concurrently to start rollup & nodemon in watch mode for the web/app/ & internal/ directories
npm run dev 
```