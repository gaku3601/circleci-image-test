faas-cli version

docker login registry.gitlab.com -u ${UserName} -p ${Password}
faas-cli build -f hello-faas.yml
faas-cli push -f hello-faas.yml

