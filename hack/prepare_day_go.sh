!#/usr/bin/env bash

YEAR=$(date +"%Y")
DAY=$(date +"%-d")

AOC_SESSION_COOKIE=$(cat ~/.aoc-session-cookie)
AOC_INPUT_FILE="https://adventofcode.com/$YEAR/day/$DAY/input"

CURRENTDAY=$(YEAR)/day$(DAY)

#all: $(CURRENTDAY)
#
#
#function setup {
#	mkdir -p ./go/$YEAR/
#	cp -r _template ./$(CURRENTDAY)
#	@sed -i 's/Day0/Day'$(DAY)'/g' ./$(CURRENTDAY)/_template_test.go
#	@sed -i 's/Day0/Day'$(DAY)'/g' ./$(CURRENTDAY)/_template.go
#	@sed -i 's/day0/day'$(DAY)'/g' ./$(CURRENTDAY)/_template_test.go
#	@sed -i 's/day0/day'$(DAY)'/g' ./$(CURRENTDAY)/_template.go
#	@mv ./$(CURRENTDAY)/_template.go ./$(CURRENTDAY)/day$(DAY).go
#	@mv ./$(CURRENTDAY)/_template_test.go ./$(CURRENTDAY)/'day'$(DAY)'_test.go'
#	@echo "folder ./$(CURRENTDAY) from template created"
#	@echo $(AOC_SESSION_COOKIE)
#	@curl --cookie $(AOC_SESSION_COOKIE) $(AOC_INPUT_FILE) -o ./$(CURRENTDAY)/input.txt
#}
#
#$(CURRENTDAY):
