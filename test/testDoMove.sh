#!/bin/bash
echo Simple move on three by three:
./robosolver 3 100000000000000000 move 0 2

echo Illegal move:
./robosolver 3 100000000000000000 move 0 1
