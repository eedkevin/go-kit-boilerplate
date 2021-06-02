define genAllProto
	for dir in * ; do \
		if [ -d $$dir ]; then \
		  if [[ -n $$(find . -type f -name *.proto ) ]]; then \
				cd $$dir && truss --svcout . *.proto && cd - ; \
			fi \
		fi \
	done
endef

.PHONY: gen-all
gen-all:
	$(call genAllProto)

.PHONY: up
up:
	docker-compose up

.PHONY: build-up
build-up:
	docker-compose up --build

.PHONY: gen-all-build-up
gen-all-build-up: gen-all build-up
	$(call gen-all, build-up)
