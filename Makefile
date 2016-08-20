help:
	@echo "\
help     this message\n\
setup    build all packages, run all tests\n\
test     generate code-coverage report\n\
run      setup local environment, and drop into shell\n\
teardown tearing down local env\n\
"

define pinfo
@printf "\033[0;33m-- $1\033[0m\n"
endef

setup:
	$(call pinfo,setting up local env)

test:
	$(call pinfo,testing)

run:
	$(call pinfo,running)
	
teardown:
	$(call pinfo,tearing down local env)
	docker-compose kill
	docker-compose rm -v