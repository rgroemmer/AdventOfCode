#!/usr/bin/env bash

YEAR=$(date +"%Y")
DAY=$(date +"%-d")
COOKIE="session=$(cat .cookie)"
FILE_SOURCE="https://adventofcode.com/$YEAR/day/$DAY/input"

function setup {
	cargo new --vcs none $DEST_PATH
	curl --cookie $COOKIE $FILE_SOURCE -o ./$DEST_PATH/src/input
}

mkdir -p "./rust/$YEAR"

for d in $(seq 1 "$DAY"); do
	DEST_PATH="./rust/$YEAR/day$d"
	[[ ! -f $DEST_PATH/src/input ]] && setup
done
