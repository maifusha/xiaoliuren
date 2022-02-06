#!/bin/sh

set -eo pipefail

if [ "$(ls -A ./scripts)" ]
then
    for script in ./scripts/*; do
        echo "running $script"; "$script";
    done
fi

exec "$@"