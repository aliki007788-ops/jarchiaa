#!/bin/bash

echo "?? Seeding database..."
for file in database/seeds/*.sql; do
echo "Seeding: $(basename $file)"
psql -h postgres -U jarchia -d jarchia -f $file
done
echo "? Seeding completed"


