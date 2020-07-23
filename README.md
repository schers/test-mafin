# test-mafin
Test task for mafin company

**how to use**

Send a POST request with the following headers:

`POST /files HTTP/1.1
Content-Length: 21744
Content-Disposition: attachment; filename="image.jpg"
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