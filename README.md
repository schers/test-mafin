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
         "filename": "image.jpg",
         "id": 1,
         "size": 21744,
         "url": "/image/2020/5p/17nvwu/image.jpg"
     },
     "status": "ok"
 }`
 
 By "url" you can get the downloaded file
