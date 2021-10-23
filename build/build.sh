#!/usr/bin/bash
source /home/ubuntu/build/.envs

# Build binary
cd $PROJECT_HOME
go build -o $BUILD_DIR "$PROJECT_HOME/$MAIN_SOURCE"
