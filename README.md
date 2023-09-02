# Initializing the Project

Before starting the project, ensure that the required ports are available. Use the following commands to check:

```bash
sudo lsof -i :5432  # Check if the default PostgreSQL database port is in use
sudo lsof -i :8080  # Check if the port for the our Golang app is in use
```

The command output will resemble:

```bash
COMMAND   PID     USER   FD  TYPE DEVICE  SIZE NODE NAME
docker    104911  root   4u  IPv4 906734  t0   TCP  *:postgresql (LISTEN)
```

If any of the ports are occupied, you'll need to terminate the corresponding process. To do this, note the PID from the output and execute:

```bash
sudo kill <PID>
```

Once the ports are available, create a `.env` file in the root directory. Populate this file with the contents of `example.env`. The only variable you might want to change is `DB_PASSWORD`.

```
DB_USER=postgres
DB_NAME=loja
DB_PASSWORD=123456
DB_HOST=postgres
DB_SSLMODE=disable
```

With the environment set up, you can start the services using:

```bash
docker-compose up
```

After executing the above command, you can access the web-shop application at:

[http://localhost:8080](http://localhost:8080)

## Simple Web Shop Application using Golang

- Homepage
 ![Homepage](images/image.png)
- Create New Product
 ![Create New Product](images/image-1.png)
- Edit Product
 ![Edit Product](images/image-2.png)
- Delete Product
 ![Delete Product](images/image-3.png)
