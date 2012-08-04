#!/bin/bash
echo Simple move on three by three:
./robosolver 3 "19 01 03 08 00 02 0c 04 06" move 0 2

echo Illegal move:
./robosolver 3 "19 01 03 08 00 02 0c 04 06" move 0 1
