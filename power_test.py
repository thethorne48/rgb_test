#!/usr/bin/python
import RPi.GPIO as GPIO
import time

redLED = 35
GPIO.setmode(GPIO.BCM)
GPIO.setup(redLED, GPIO.IN)

powerlow = 0
while True:
    if(GPIO.input(redLED) == 0):
        print("POWER dipped below 4.63V")
        powerlow += 1
    else:
        powerlow = 0
        if (powerlow > 3):
            print(f"Low power for {str(powerlow)} seconds")
            time.sleep(1)
