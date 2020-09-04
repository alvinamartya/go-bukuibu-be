# Backend for Diary app Buku Ibu
A diary application to notes your feeling in everyday

# Run via Docker
```
docker build -t api-buku-ibu:1.0 .
docker-compose up -d
docker exec -it postgres psql -U postgres -c "CREATE DATABASE buku_ibu_db"
```

