#!/bin/bash
CURRENT_DIR=$1
for x in $(find ${CURRENT_DIR}/genproto/* -type d); do
  cd ${x} && ls *.pb.go | sudo xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
done