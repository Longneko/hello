mysql_var_names := MYSQL_DATABASE MYSQL_ROOT_PASSWORD
mysql_env_string :=$(foreach var,$(mysql_var_names),-e $(var)=$$HELLO_$(var))

mysql_container_name := hello-mysql
docker-up:
ifneq ($(shell docker ps -a | grep -w ${mysql_container_name}),)
	$(info starting existing "${mysql_container_name}" MySQL container...)
	docker start $(mysql_container_name)
else
	$(info MySQL container "${mysql_container_name}" not found. Running new...)
	# ensure necessary variables are not empty
	$(foreach var, $(mysql_var_names), $(if ${HELLO_$(var)}, , $(error var HELLO_$(var) is not set)))

	docker run -d -p 3306:3306 \
	--name $(mysql_container_name) \
	--volume=$(shell pwd)/data/mysql/:/var/lib/mysql/ ${mysql_env_string} \
	mysql:8.0.18
endif

docker-down:
	docker stop $(mysql_container_name)
	
dev: export HELLO_APP_MODE=debug
dev:
	go run main.go

prod: export HELLO_APP_MODE=release
prod:
	go run main.go

test:
	$(if ${HELLO_MYSQL_DATABASE}, $(info ${HELLO_MYSQL_DATABASE} != empty), $(error var HELLO_MYSQL_DATABASE == empty))
