# Ngnix config

1. Install ngnix

`sudo apt-get update`

`sudo apt-get install nginx`

2. unlink the default file for ngnix

`sudo unlink /etc/nginx/sites-enabled/default`

3. sample file(with one method) create and change ips

`cd /etc/nginx/sites-available/`

`nano load-balancing.conf`

```
upstream backend {
    server ip:8081;
    server ip:8081;
}

server {
    listen 80;
    location / {
        proxy_pass http://backend;
    }
}
```

4. run your service on two machines & change the `ip` in above file


`sudo ln -s /etc/nginx/sites-available/load-balancing.conf /etc/nginx/sites-enabled/`

`service nginx configtest`

`service nginx restart`