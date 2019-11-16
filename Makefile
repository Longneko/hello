mysql_var_names := MYSQL_DATABASE MYSQL MYSQL_ROOT_PASSWORD
enmysql_env_string :=$(foreach var,$(mysql_var_names),-e $(var)=$$HELLO_$(var) \
)

mysql_container_name := hello-mysql
docker-up:
ifneq ($(shell docker ps -a | grep -w ${mysql_container_name}),)
	$(info starting existing "${mysql_container_name}" MySQL container...)
	docker start $(mysql_container_name)
else
	$(info MySQL container "${mysql_container_name}" not found. Running new...)
	docker run -d -p 3306:3306 --name $(mysql_container_name) --volume=$(shell pwd)/data/mysql/:/var/lib/mysql/ ${enmysql_env_string} mysql:8.0.18
endif

docker-down:
	docker stop $(mysql_container_name)
	
dev:
	go run main.go
