#!/bin/bash

for i in `seq 1 100`
do
  echo "Please enter command"
  read cmd
  echo "Exuecting..."
  eval $cmd
done
