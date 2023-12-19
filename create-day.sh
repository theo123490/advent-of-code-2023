#!/bin/bash

NEW_FILE=$1

cp -a dayx $NEW_FILE
mv ./$NEW_FILE/dayxTools "./$NEW_FILE/${NEW_FILE}Tools"
sed -i '' -e "s/dayx/$NEW_FILE/g" ./$NEW_FILE/main.go
sed -i '' -e "s/dayx/$NEW_FILE/g" ./$NEW_FILE/${NEW_FILE}Tools/tools.go
touch ./$NEW_FILE/example.txt
touch ./$NEW_FILE/input.txt
