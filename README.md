# test-mafin
Test task for mafin company

**How to use**

Send a POST request with the following headers on localhost:8080:

`POST /api/v1/image/upload HTTP/1.1
Content-Length: 21744
Content-Type: image/png
Content-Disposition: image; filename="image.png"
...bytes...`

In response, you should receive something like the following json:

`{
     "files": {
         "filename": "image.png",
         "id": 1,
         "size": 21744,
         "url": "/image/2020/5p/17nvwu/image.png"
     },
     "status": "ok"
 }`
 
 By "url" you can get the downloaded file
 
 If you want to run the application in docker, copy the docker-compose.override.yml.dist file and remove the .dist extension.
 Then run `docker-compose up -d --build`
