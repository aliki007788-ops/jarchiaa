#!/bin/bash

BACKUP_DIR="/backups/jarchia"
DATE=$(date +%Y%m%d_%H%M%S)

mkdir -p $BACKUP_DIR
pg_dump -h postgres -U jarchia jarchia | gzip > $BACKUP_DIR/db_$DATE.sql.gz
echo "? Backup completed: $BACKUP_DIR/db_$DATE.sql.gz"


