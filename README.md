menu
====

A tool for showing menu under linux command line.

## Usage
### Install

```bash
go get github.com/daviddengcn/menu
go install github.com/daviddengcn/menu
cp $GOPATH/src/github.com/daviddengcn/menu/menu.sh $GOPATH/bin
```

### Run

```bash
menu.sh "Item 0" "Item 1" "Item 2"
```

When the user make the choice by pressing `Enter`, the command returns and exit code is set to the zero-based index of the selected item. If the user press `Ctrl+C`, exit code is set to `255`.

### Example

The following script shows an menu of all non current branches in the current git repository, and the user can choose a branch to checkout. A `cancel` button is provided as well.
```base
#!/bin/bash

branches=`git branch | sed -e "s/^.*[*].*$//g"`
branches=($branches)

branchline=$(echo "$branches" | tr "\\n", " "| sed -e "s/  */ /g")

menu.sh "`echo -e "\x1b[31mcancel\x1b[m"`" ${branches[*]}

sel=$?

[[ $sel -eq 0 ]] && exit 0
[[ $sel -eq 255 ]] && exit 1

selbranch=${branches[`expr $sel - 1`]}

git checkout $selbranch
```
