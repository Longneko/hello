# Hello (lamp)

This is a web app that says "hello" to the world and can be said "hello" back to. It's sole purpose is a demonstration of me setting up a LAMP (LEMG actually).

## Requirements
- Docker / MySQL 8.0
- Nginx
- GoLang 1.13

## Configuration
Hello comes with a config templates for the app and the nginx server. Both contain widely used default values ("8080" port for the application, "root:root" credentials and localhost address for the MySQL, etc.). All (or most) of the values are overriden in the application if corresponding environment variables are set.

### Initializing config files


### Nginx config

### Environment Variables
As mentioned above, the env variables are mostly required to override values not set in the config, which is especially relevant for the credentials. Some of them, however, are also relevant during config initialization and docker MySQL container creation.


### First launch
Before the first launch, it is required to create config files from templates. If using docker contaier for MySQL, the first run also initiates a new database with set credentials, unless there already is one provied in the volume

### Docker MySQL container

