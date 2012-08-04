#!/bin/bash

echo Simple solution
`dirname $0`/../robosolver 3 "19 01 03 08 00 02 0c 04 06" to 0 8

echo Unsolvable:
`dirname $0`/../robosolver 3 "19 01 03 08 00 02 0c 04 06" to 0 4
