package main

import (
	"fmt"
)

var (
	gyroHeading = 0.0
)

// main simulates turning using a modern robotics gryo in basic mode.
func main() {
	turn(90, false)
	turn(90, true)
	turn(90, false)
}

func getHeading(turnLeft bool) float64 {
	if turnLeft {
		if gyroHeading-1 < 0 {
			gyroHeading = 359
		} else {
			gyroHeading -= 1
		}
	} else {
		if gyroHeading+1 > 359 {
			gyroHeading = 0
		} else {
			gyroHeading += 1
		}
	}
	return gyroHeading
}

// Right now this turns both motors with equal but opposite power.
// What if we just turned one motor, and set the other motor to float (not brake), would that work better?

func turn(targetHeading float64, turnLeft bool) {
	midPower := 0.0

	targetHeading = targetHeading - gyroHeading
	fmt.Println(targetHeading, gyroHeading)
	if targetHeading > 0 && turnLeft {
		targetHeading += 180
	}

	for {
		currentHeading := getHeading(turnLeft)
		if !turnLeft && currentHeading >= targetHeading {
			break
		}
		if turnLeft && currentHeading <= targetHeading {
			break
		}

		headingError := targetHeading - currentHeading
		driveSteering := (headingError / 180)
		leftPower := 0.0
		rightPower := 0.0
		if turnLeft {
			leftPower = midPower + driveSteering
			rightPower = midPower - driveSteering
		} else {
			leftPower = midPower + driveSteering
			rightPower = midPower - driveSteering
		}
		if leftPower < .10 && leftPower > 0 {
			leftPower = .1
		}
		if leftPower > -.10 && leftPower < 0 {
			leftPower = -.1
		}
		if rightPower < .10 && rightPower > 0 {
			rightPower = .1
		}
		if rightPower > -.10 && rightPower < 0 {
			rightPower = -.1
		}

		fmt.Printf("Current Heading: %f, Heading Error: %f, drive steering: %f, left power: %f, right power %f\n", currentHeading, headingError, driveSteering, leftPower, rightPower)
	}
}
