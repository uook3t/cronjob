#!/bin/bash

mkdir -p "output"

go build -o "./output/bootstrap.bin"
chmod +x ./output/bootstrap.bin