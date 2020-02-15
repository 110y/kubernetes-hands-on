.PHONY: app-container
app-container:
	cd ./app && docker build -t 110y/kubernetes-hands-on:$(shell git rev-parse --short HEAD) .

.PHONY: app-registry
app-registry: app-container
	docker push 110y/kubernetes-hands-on:$(shell git rev-parse --short HEAD)
