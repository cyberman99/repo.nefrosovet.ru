#!/bin/bash

./cockroach sql --execute="CREATE DATABASE IF NOT EXISTS ${PATIENT_DB_DATABASE};" --url=postgresql://${PATIENT_DB_HOST}?sslmode=disable
./cockroach sql --execute="CREATE USER IF NOT EXISTS ${PATIENT_DB_LOGIN};" --url=postgresql://${PATIENT_DB_HOST}?sslmode=disable
./cockroach sql --execute="GRANT ALL ON DATABASE ${PATIENT_DB_DATABASE} TO ${PATIENT_DB_LOGIN};" --url=postgresql://${PATIENT_DB_HOST}?sslmode=disable
