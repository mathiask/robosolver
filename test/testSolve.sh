#!/bin/bash
echo
echo FullSolver
`dirname $0`/../robosolver 3 "19 01 03\
                              28 00 02\
                              3c 44 06" solve 8

echo Now solvable ':-)'
`dirname $0`/../robosolver 3 "19 01 03\
                              28 00 02\
                              3c 44 06" solve 4

echo FullSolver moving pieces out of the way
`dirname $0`/../robosolver 3 "19 01 03\
                              08 00 22\
                              3c 44 06" solve 8
