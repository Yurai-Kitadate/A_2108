#!/usr/bin/bash
mkdir db
cat << __EOF__ >> ../.gitignore
docker_compose/db
docker_compose/docker-compose.yml
__EOF__
cp docker-compose.yml docker-compose.yml.1
sed "s/# REPLACE MARKER FOR PERSISTENCE/- .\/db:\/var\/lib\/mysql/g" docker-compose.yml > docker_compose.yml
mv -f docker_compose.yml docker-compose.yml