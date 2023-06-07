#!/bin/bash

go build -o bookings cmd/web/*.go && ./bookings

# make the run file executable by running the following program
# chmod +x run.sh