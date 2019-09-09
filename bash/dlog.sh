#!/usr/bin/env bash

ctId=`docker ps | grep ${1} | awk '{print $1}'`

if [ !$2 ] ;
then
    docker logs $ctId > $2  2>&1
else
    docker logs $ctId
fi