#!/bin/bash
echo
echo Find 1 in the center
`dirname $0`/../robosolver 3 "09 01 03 08 10 02 0c 04 06" find 1

echo Search in vain
`dirname $0`/../robosolver 3 "09 01 03 08 20 02 0c 04 06" find 1

echo Find 1 of 4
`dirname $0`/../robosolver 3 "19 01 03\
                              28 00 02\
                              3c 44 06" find 1

echo Find 2 of 4
`dirname $0`/../robosolver 3 "19 01 03\
                              28 00 02\
                              3c 44 06" find 2

echo Find 3 of 4
`dirname $0`/../robosolver 3 "19 01 03\
                              28 00 02\
                              3c 44 06" find 3

echo Find 4 of 4
`dirname $0`/../robosolver 3 "19 01 03\
                              28 00 02\
                              3c 44 06" find 4
