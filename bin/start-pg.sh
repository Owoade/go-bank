#!/bin/bash

# deletes existing docker pg container
sh ./bin/docker-cp.sh

# runs pg container
sh ./bin/docker-pg.sh

# runs migration
sh ./bin/migration.sh
