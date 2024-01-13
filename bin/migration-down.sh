#!/bin/bash 

migrate -path db/migration -database "postgres://root:secret@localhost:5432/bank-app?sslmode=disable" -verbose down 