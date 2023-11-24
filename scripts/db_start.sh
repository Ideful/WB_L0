docker run -d \
  --name orders \
  -p 5432:5432 \
  -v /Users/mirianfu/golang/wb/l0/scheme/:/docker-entrypoint-initdb.d/ \
  -v /Users/mirianfu/golang/wb/l0/scheme/:/scheme \
  -e POSTGRES_PASSWORD=0 \
  postgres:latest

until docker exec -it orders pg_isready -h localhost -p 5432 -U postgres; do
  echo "waiting for DB container start"
  sleep 1
done
