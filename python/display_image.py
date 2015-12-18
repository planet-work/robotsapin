#!/usr/bin/python

import time
import sys
import os
import base64
from sense_hat import SenseHat

sense = SenseHat()
sense.low_light = True

if os.path.exists(sys.argv[1]):
    sense.clear()
    sense.load_image(sys.argv[1])
else:
    try:
       img = base64.b64decode(sys.argv[1])
       sense.clear()
       sense.load_image(img)
    except TypeError:
       sys.exit(1)
