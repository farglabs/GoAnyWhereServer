## GoAnywhere-Server

Run the following commands to get up and going
As of right now this is designed to run on `:443` and `:80` which are protected ports on most systems. That being said, you'll need to run this with elevated privliges. 

```sh
go run fileserver
```

### Routes
1. https://fileserver/files/
2. https://fileserver/files/test.jpg
3. https://fileserver/files/test.pdf
4. https://fileserver/api/

All other routes should yeild a `404`


