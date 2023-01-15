#!/bin/bash

while [ 1 ]; do
    ulimit -n 999999
    ./cnc
done