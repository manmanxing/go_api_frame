#!/usr/bin/env bash
git pull
bash shutdown.sh
make clean
bash start.sh
