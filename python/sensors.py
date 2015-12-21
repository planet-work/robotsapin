#!/usr/bin/python

import time
import sys
import os
import base64
import json
from sense_hat import SenseHat

sense = SenseHat()


sensors = {"temperature": sense.temperature, 
           "humidity": sense.humidity, 
           "pressure" : sense.pressure,
           "orientation" : sense.orientation}
print json.dumps(sensors)
