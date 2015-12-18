#!/usr/bin/python

import time
import sys
import os
import base64
import StringIO
from PIL import Image  # pillow
from sense_hat import SenseHat

sense = SenseHat()
sense.low_light = True

try:
    data = sys.argv[1]
except IndexError:
    display_pixels =  sense.get_pixels()
    img = Image.new( 'RGB', (8,8), "black")
    pixels = img.load()
    ## Set pixels ... 
    idx = 0
    for i in range(8):
        for j in range(8):
            pixels[i,j] = tuple(display_pixels[idx])
            idx += 1
    img = img.rotate(-90)
    output = StringIO.StringIO()
    img.save(output,format="PNG")
    print base64.b64encode(output.getvalue())
    sys.exit(0)

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
