#!/bin/bash

#!/usr/bin/env bash
set -eo pipefail

# Create mount directory for service

mkdir -p /var/db/micromdm/file
ln -s $HOME /var/db/micromdm/file

echo "Mounting GCS Fuse."
gcsfuse --debug_gcs --debug_fuse bricks-micromdm-file-bucket /var/db/micromdm
echo "Mounting completed."

/usr/bin/micromdm serve -server-url  https://bricks-micromdm-u6idxxekbq-de.a.run.app -api-key 12345 -tls=false

exec "$@"
