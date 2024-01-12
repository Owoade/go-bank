#!/bin/bash

docker run --name pg-bank -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=bank-app -p 5432:5432 -d postgres
