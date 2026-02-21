#!/bin/bash

echo "?? Obfuscating frontend code..."
cd frontend
npm install -g javascript-obfuscator
javascript-obfuscator ./dist --output ./dist-obfuscated
echo "? Obfuscation completed"


