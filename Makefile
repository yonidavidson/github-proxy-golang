help:
	@echo "\
help     this message\n\
shell    shell promt in container\n\
setup    build all packages, run all tests\n\
test     generate code-coverage report\n\
run      setup local environment, and drop into shell\n\
teardown tearing down local env\n\
"

define pinfo
@printf "\033[0;33m-- $1\033[0m\n"
endef

setup:
	$(call pinfo,setting up local env - not used)
	docker-compose build

test:
	$(call pinfo,testing)
	docker-compose up tester

run:
	$(call pinfo,running)
	docker-compose up web

teardown:
	$(call pinfo,tearing down local env)
	docker-compose kill
	docker-compose rm -v

shell:
	$(call pinfo,entring shell promt)
	docker-compose run --rm tester  /bin/bash