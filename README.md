# Fire-Dragon
Backend for Best Wishes iOS App

# Development
## Docker

### Run locally
docker build --build-arg mysql_user="" --build-arg mysql_pass="" --build-arg mysql_host_port="" --build-arg mysql_db="" --build-arg jwt_sec_key="" .
docker run -p 127.0.0.1:443:443 b3b66e45a5bf

