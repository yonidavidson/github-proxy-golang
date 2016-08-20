help:
	@echo "\
help     this message\n\
shell    shell promt in container\n\
setup    setup envrioment - not in use\n\
test     run all tests\n\
run      start app\n\
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