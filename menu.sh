#!/bin/bash
stty raw -echo
menu "$@"
res=$?
stty cooked echo
exit "$res"
