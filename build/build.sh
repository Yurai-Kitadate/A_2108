#!/usr/bin/bash
source ./build/.envs

# Build binary
cd $PROJECT_HOME
go build -o $BUILD_DIR "$PROJECT_HOME/$MAIN_SOURCE.go"
cp -f "$BUILD_DIR/$MAIN_SOURCE" $INSTALL_DIR