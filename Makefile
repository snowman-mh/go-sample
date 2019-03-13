###############
## Variables ##
###############

DOCKER_COMPOSE_FILE := docker-compose.yml
DOCKER_COMPOSE_FILE_TEST := docker-compose.test.yml
CONTAINER_NAME_PREFIX := go_sample
COMPOSE_COMMAND :=
DOCKER_COMMAND :=
TEST_SUFFIX :=
TEST_OPTION :=

##############################################################
## Usage: make (local|test) (start|restart|stop|logs|mysql) ##
##############################################################

## Operate local development environment
local:
	@$(eval COMPOSE_COMMAND = cd docker; docker-compose -f ${DOCKER_COMPOSE_FILE})
	@$(eval DOCKER_COMMAND = docker exec -it ${CONTAINER_NAME_PREFIX})

## Operate test environment
test:
	@$(eval COMPOSE_COMMAND = cd docker; docker-compose -f ${DOCKER_COMPOSE_FILE_TEST})
	@$(eval DOCKER_COMMAND = docker exec -it ${CONTAINER_NAME_PREFIX})
	@$(eval TEST_SUFFIX = _test)
	@$(eval TEST_OPTION = --project-name test)

### Start selected environment
start:
	${COMPOSE_COMMAND} ${TEST_OPTION} up -d

### Build and Start selected environment
restart:
	${COMPOSE_COMMAND} ${TEST_OPTION} up -d --build

### Shut down selected environment
stop:
	${COMPOSE_COMMAND} ${TEST_OPTION} down -v

### Show logs from selected environment
logs:
	${COMPOSE_COMMAND} ${TEST_OPTION} logs -f

### Connect MySQL on selected environment
mysql:
	${DOCKER_COMMAND}_mysql${TEST_SUFFIX} mysql -uroot -ppassword go_sample

###############################################################
## Usage: make test (run|clean)                              ##
##        make test run target=path/to/test/files options=-v ##
###############################################################

### Run test on test environment
run:
ifdef target
	${DOCKER_COMMAND}_api${TEST_SUFFIX} go test ${target} -cover ${options}
else
	${DOCKER_COMMAND}_api${TEST_SUFFIX} go test ./... -cover ${options}
endif

### Clean test cache on test environment
clean:
	${DOCKER_COMMAND}_api${TEST_SUFFIX} go clean -testcache

#################################################
## Usage: make migrate new name=migration_name ##
##        make migrate (up|down)               ##
#################################################

## Execute migration
migrate:
	@$(eval DOCKER_COMMAND = docker exec -it ${CONTAINER_NAME_PREFIX}_dbmate dbmate)

### Create new migration
new:
ifdef name
	${DOCKER_COMMAND} new ${name}
else
	@echo "Usage: make migrate new name=migration_name"
endif

### Migrate up
up:
	${DOCKER_COMMAND} up

### Migrate down
down:
	${DOCKER_COMMAND} down

###########################################################
## Usage: make addmod target=github.com/org/repo@version ##
###########################################################

## Add new module to go.mod
addmod:
ifdef target
	docker exec -it ${CONTAINER_NAME_PREFIX}_api go mod edit -require=${target}
else
	@echo "Usage: make addmod target=github.com/org/repo@version"
endif
