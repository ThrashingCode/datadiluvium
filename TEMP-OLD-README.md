# Data Diluvium

<p><a href="https://app.codeship.com/projects/192753"><img src="https://camo.githubusercontent.com/c1db7e8dfc1f282478b94828bd1ced0a07ce50f2/68747470733a2f2f696d672e736869656c64732e696f2f636f6465736869702f61363063303130302d616564642d303133342d343863652d3661623531303239343865382f6d61737465722e7376673f7374796c653d666c6174" alt="Codeship Status for Adron/datadiluvium" data-canonical-src="https://img.shields.io/codeship/a60c0100-aedd-0134-48ce-6ab5102948e8/master.svg?style=flat" style="max-width:100%;"></a> <a href="https://goreportcard.com/report/Adron/datadiluvium"><img src="https://camo.githubusercontent.com/92e33f81c97ab1fd45f730a35a018cb5000e8f29/68747470733a2f2f676f7265706f7274636172642e636f6d2f62616467652f4164726f6e2f6461746164696c757669756d" alt="Go Report Card" data-canonical-src="https://goreportcard.com/badge/Adron/datadiluvium" style="max-width:100%;"></a> <a href="https://hub.docker.com/r/adron/datadiluvium/"><img src="https://camo.githubusercontent.com/15b899950b587a1ee7f3037b1b790b0203d02e2b/68747470733a2f2f696d672e736869656c64732e696f2f646f636b65722f6175746f6d617465642f6164726f6e2f6461746164696c757669756d2e7376673f7374796c653d666c6174" alt="Docker Automated buil" data-canonical-src="https://img.shields.io/docker/automated/adron/datadiluvium.svg?style=flat" style="max-width:100%;"></a>  <a href="https://waffle.io/Adron/datadiluvium"><img src="https://camo.githubusercontent.com/de99eeaa245b28544413aba863c5af7ff7040b5b/68747470733a2f2f696d672e736869656c64732e696f2f776166666c652f6c6162656c2f4164726f6e2f6461746164696c757669756d2f72656164792e7376673f7374796c653d666c617426636f6c6f72423d677265656e" alt="Waffle.io" data-canonical-src="https://img.shields.io/waffle/label/Adron/datadiluvium/ready.svg?style=flat&amp;colorB=green" style="max-width:100%;"></a> <a href="https://waffle.io/Adron/datadiluvium"><img src="https://camo.githubusercontent.com/03e0bc8e91b5eb41603fdea798773122dfd3ac85/68747470733a2f2f696d672e736869656c64732e696f2f776166666c652f6c6162656c2f4164726f6e2f6461746164696c757669756d2f696e25323070726f67726573732e7376673f7374796c653d666c617426636f6c6f72423d677265656e" alt="Waffle.io" data-canonical-src="https://img.shields.io/waffle/label/Adron/datadiluvium/in%20progress.svg?style=flat&amp;colorB=green" style="max-width:100%;"></a> <a href="https://github.com/Adron/datadiluvium/blob/master/LICENSE"><img src="https://camo.githubusercontent.com/5f1f0b9009dedf14412f1233759d062d080cbc9e/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f6c6963656e73652f4164726f6e2f6461746164696c757669756d2e7376673f7374796c653d666c6174" alt="Data Dilvium License" data-canonical-src="https://img.shields.io/github/license/Adron/datadiluvium.svg?style=flat" style="max-width:100%;"></a> <a href="https://github.com/Adron/datadiluvium/tree/master/.github"><img src="https://camo.githubusercontent.com/3e9472f0943771d602a3aabc3fc580c171bd4514/68747470733a2f2f696d672e736869656c64732e696f2f6d61696e74656e616e63652f7965732f323031372e7376673f7374796c653d666c6174" alt="Maintenance" data-canonical-src="https://img.shields.io/maintenance/yes/2017.svg?style=flat" style="max-width:100%;"></a></p>

# Installation of the Service

 1. If you don't have Go installed, that's the first step. Install Go.
 2. Fork the project (if you intend to contribute, if you're just going to try it out, skip to step 3)
 3. git clone git@github.com:YourUsername/datadiluvium.git (or git@github.com:Adron/datadiluvium.git if you're just trying it out)
 4. If you don't have *glide* or *hugo* installed, run the `setup.sh` file.
 
This should enable you to run the project with a simple `go run *.go`. As for deployment, that's a different matter and an exercies that I'll leave up to you.

This project is a data generation library for varying databases and other storage mechanisms. For more information on the library ping me [@adron](https://twitter.com/Adron) to ask questions, dive in with some code of your own, or other suggestions or comments.

For more information check out the project site @ http://datadiluvium.com

...and yes, more docs are on their way.  ;-)

# Services

* `curl -X POST -H "Content-Type: application/json" -d '{"Schema":"blagh request","Database":"yup"}' http://localhost:8080/schemagen/` will submit a POST against the end point and return the generated data results.
* `curl localhost:8080/schemagen/` will submit a GET against the end point and return instructions for the schemagen POST verb command at the same end point.
* `curl localhost:8080/` will submit a GET against the end point and return directions on the API.

Sample JSON Schemas to submit to the service. You can submit one or multiple to the service within a JSON array.

```javascript
[
  {
    "datastore": "postgresql",
    "connection": "the-postgresql-connection-string",
    "generate": [
      {
        "name": "users",
        "genrows": "320",
        "data": [
          {
            "name": "id",
            "key": "pk",
            "data": "uuid"
          },{
            "name": "username",
            "data": "user"
          },{
            "name": "firstname",
            "data": "first"
          },{
            "name": "lastname",
            "data": "last"
          },{
            "name": "email",
            "data": "email"
          },{
            "name": "password",
            "data": "password"
          }
        ]
      },{
        "name": "purchases",
        "genrows": "400",
        "data": [
          {
            "name": "fk_id",
            "key": "fk",
            "primary": "users"
          },{
            "name": "product",
            "data": "name"
          },{
            "name": "price",
            "data": "money"
          },{
            "name": "paid",
            "data": "money"
          },{
            "name": "notes",
            "data": "lorumipsum"
          }
        ]
      },{
        
      }
    ]
  },
  {
    "datastore": "redis",
    "genrows": "225",
    "connection": "the-redis-connection-string",
    "generate": [
      {
        "name": "key",
        "data": "uuid"
      },{
        "name": "value",
        "data": "string"
      }
    ]
  },
  {
    "datastore": "console",
    "genrows": "25",
    "generate": [
      {
        "name": "username",
        "data": "user"
      },
      { 
        "name": "email",
        "data": "email"
      }
    ]
  }
]
```

Simple curl statement with just the console generation request.

```
curl -X POST -H "Content-Type: application/json" -d '[{"datastore": "console","genrows": "25","generate": [{"name": "username","data": "user"},{"name": "email","data": "email"}]}]' http://localhost:8080/schemagen/
```


For more information about Hugo the static site generator written in Go for the [Data Diluvium](http://datadiluvium.com) site, checkout my [NOTES.md](docs/hugo/NOTES.md) file.
