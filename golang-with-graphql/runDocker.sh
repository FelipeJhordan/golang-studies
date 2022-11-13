
docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=hackernews -d mysql:latest


docker exec -it mysql bash



mysql -u root -p



CREATE DATABASE hackernews;