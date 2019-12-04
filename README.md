# Hello (lamp)

This is a web app that says "hello" to the world and can be said "hello" back to. It's sole purpose is a demonstration of me setting up a LAMP (LEMG actually).

## Requirements
- Docker / MySQL 8.0
- Nginx
- GoLang 1.13

## Configuration
Hello comes with a config templates for the app and the nginx server. Both initially provide widely used default values ("8080" port for the application, "root:root" credentials and localhost address for the MySQL, etc.). All (or most) of the values can be overriden in the application if corresponding environment variables are set.

### Environment Variables
As mentioned above, the env variables are mostly required to override values not set in the app config, which is especially relevant for the credentials. They can also be used in the config files creation and running docker MySQL container.


Var Name | Description
--- | ---
HELLO_APP_MODE | ["debug", "release"] mode set for the app's Gin-Gonic server affecting like log output
HELLO_MYSQL_ADDRESS | MySQL server host address
HELLO_MYSQL_DATABASE | MySQL database name
HELLO_MYSQL_USER | MySQL user
HELLO_MYSQL_PASSWORD | MySQL password
HELLO_MYSQL_ROOT_PASSWORD | used to set root password by docker image on when creating a new container 
HELLO_SERVER_HOST | host on which the app server should be listening
HELLO_SERVER_PORT | port on which the app server should be listening        
HELLO_SERVER_READ_TO | "00h00m00s" format (e.g. "15s") app server's read timeout
HELLO_SERVER_WRITE_TO | app server's write timeout
*NOTE: values should not be quoted*

### Initializing config files
You can create needed config files for both the application and the nginx server from provided tempalates by running while in the app's root dir:
```bash
make configs
```
This will create (but not overwrite if already exist) both files:
* `./conf/app.conf` - app config
* `./conf/external/hello.nginx.conf` - Nginx config

If any of the above listed env variables are set, they will be used to populate the config instead of the defaults. With the exception of the credentials, for which the defaults are always used.

### Nginx config
The Nginx config can be found in the `./conf/external/` directory after its initialization. It proxy passes requests to the app server and serves the static files from the project directory (currently empty). To work its contents can either be copied into an existing Nginx configuration file, or (which is more conventional) the file itself can be included in an already existing Nginx config. The default Nginx configuration will include files added to the `/etc/nginx/sites-enabled/` directory if they fit into the "http" directive (which our file does).

### Docker MySQL container
If you have already installed docker and wish to use a MySQL container, you can do so by running:
```bash 
make docker-up
```
Which will either:
* create a new container if non exists yet. This requires the **HELLO_MYSQL_DATABASE** and **HELLO_MYSQL_ROOT_PASSWORD** env variables to be set.
* start the already existing container unless already running

To stop the container run:
```bash 
make docker-down
```
*Note: the container uses applciation's `./data/mysql/` directory as a volume to store the database values.*

## First launch
On first launch, if all the settings are valid and the connection to the MySQL database is established correctly, the necessary table will be created automatically (unless already exists) and the app is in working condition.
