faas-cli version

docker login registry.gitlab.com -u ${UserName} -p ${Password}
faas-cli build -f app.yml
faas-cli push -f app.yml

