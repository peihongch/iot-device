#!/bin/bash

ps -ef | grep iot-device | grep -v grep | awk -F ' ' '{print $2}' | xargs -I {} kill -9 {}
