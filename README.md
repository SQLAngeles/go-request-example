# go-request-example



### Requirements
* Docker

### Install
Run below command on your "github.com" directory </br>
ex) /Users/SQLAngeles/Desktop/go/src/github.com/SQLAngeles
```
$git clone https://github.com/SQLAngeles/go-request-example.git
```

### Branches
```
 master
 api/get
 api/post
 api/delete
 api/put
```

### How to test
Run below commands on "go-request-example" directory </br>
```
$docker build -t sqlangeles .
$docker run -it -p 8080:8080 -v [put your path here]:/go/src/github.com/[your-github-account]/go-request-example/ sqlangeles
```
You can simply test each apis on any api testing tools

* Example [POSTMAN]
<img width="1160" alt="screen shot 2018-04-24 at 12 20 18 am" src="https://user-images.githubusercontent.com/31301769/39172030-60eb70ee-4755-11e8-9f9f-fc3c6095093c.png">

