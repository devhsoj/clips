# Project Clips

Project Clips is a simple video hosting platform that enables people to share and view clips.

## Installation

**Requirements:** [Go](https://go.dev/dl/) - [npm](https://nodejs.org/en/download/) - [PostgreSQL](https://www.postgresql.org/download/)

```bash
git clone https://github.com/devhsoj/clips
cd clips/
npm install
npm run build
```

## Postgres Setup
```sql
CREATE DATABASE clips;
```

## Starting

**Note:** Make sure you have run `npm run build` before-hand, and have a proper `.env`

#### Run with npm
```bash
npm start # or with pm2: pm2 start npm --name "clips" -- start
```
or

#### Run from built executable
```bash
./build/clips # or ./build/clips.exe
```

## Development

#### Starting
```bash
# concurrently starts rollup in watch mode for web/app/, and nodemon to wait for changes in internal/
npm run dev 
```

## Docker
**Coming Soon**