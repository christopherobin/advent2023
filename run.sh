#!/bin/sh
targets=$(ls -1 cmd)
if [ ! -z ${1} ]; then
    targets=$(find cmd -name '*'${1:-01}'*' | xargs basename)
fi

for target in $targets; do
    echo -e "\033[1;33m> Running day $target\033[0m"
    go run "./cmd/${target}" "inputs/${target}"
done