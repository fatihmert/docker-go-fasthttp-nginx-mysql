# Go FastHTTP API Development with MySQL on NGINX

My first docker project, so it is not stable 🥺 I'm created repo for development continues..

## Intro

* NGINX
* Go v1.14
* MySQL 8.0
* PhpMyAdmin

HTML/CSS using port is: **3310**

Go API using port is: **3000**

PhpMyAdmin using port is: **8081**

Don't forgot add `local.docker` to `etc/hosts`

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

 - [ ] NGINX proxy pass correctly
 - [ ] Fix NGINX issues
 - [ ] Add Letscrypt SSL Cert
