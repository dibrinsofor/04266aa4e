# UrlPlaylists

This tool allows you to share a burst of urls as one link (bunch of supporting urls) as one link to teammates, friends, etc.


`docker run --name mongodb -e MONGO_INITDB_ROOT_USERNAME=myuser -e MONGO_INITDB_ROOT_PASSWORD=mypassword -e MONGO_INITDB_DATABASE=tasks -p 27017:27017 -d mongo:latest`
`docker build --tag urlplaylists .`
`docker run urlplaylists`