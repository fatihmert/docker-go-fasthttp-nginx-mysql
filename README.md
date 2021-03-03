# Go FastHTTP API Development with MySQL on NGINX

My first docker project, so it is not stable 🥺 I'm created repo for development continues..

## Intro

* NGINX
* Go v1.14
* MySQL 8.0
* PhpMyAdmin

### Host

| Container   |      URL & Host      |  Port |
|----------|:-------------:|------:|
| Go Api |  api.dev.go-api | 80 |
| Static HTML |    dev.go-api   |   80 |
| PhpMyAdmin | dev.go-api |    8080 |

## Application Structure

```
server              -- nginx configs, ssl's
  cert              -- ssl cert files
  nginx
    conf.d          -- your nginx applications configs
      default.conf  --
    nginx.conf      -- main nginx conf

api                 -- go API codes
client              -- html/css
```

# Todo

 - [x] NGINX proxy pass correctly
 - [x] Fix NGINX issues
 - [ ] Add Letscrypt SSL Cert
