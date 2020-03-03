#!/usr/bin/env bash
set -x
id=$(pgrep goApiFrame)
kill $id
