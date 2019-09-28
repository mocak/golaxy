# golaxy

A note on docker, psql and ./docker/docker-entrypoint-initdb.d folder:
- You have to -for now- manually create folder under: ./docker/db
- If your container has already been started and there has been a change in the contents of entrypoint folder, you'll have to: 
1. Run _docker-compose down_
2. Remove everything under ./docker/db
3. Run docker-compose up -d --build

For more information on subject see: [comment in docker/psql's issues](https://github.com/docker-library/postgres/issues/203#issuecomment-255200501)
