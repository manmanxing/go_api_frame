#!/usr/bin/env bash
set -x
id=$(pgrep go_api_frame)
kill $id
