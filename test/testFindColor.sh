#!/bin/bash

echo Find 1 in the center
`dirname $0`/../robosolver 3 "09 01 03 08 10 02 0c 04 06" find 1

echo Search in vain
`dirname $0`/../robosolver 3 "09 01 03 08 20 02 0c 04 06" find 1
