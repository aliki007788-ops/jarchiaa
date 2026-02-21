#!/bin/bash

if [ -z "$1" ]; then
echo "Usage: $0 backup-file.sql.gz"
exit 1
fi

gunzip -c $1 | psql -h postgres -U jarchia jarchia
echo "? Database restored"


