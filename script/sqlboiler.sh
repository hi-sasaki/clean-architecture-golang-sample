#!/bin/bash

set -eux

cd $(dirname $0)/../pkg/driver/mysql
sqlboiler -o dto -p dto --no-tests --wipe mysql