#!/bin/bash

# build and run the go program
echo "Running quiz app..."
go build && ./quiz-app -csv=questions.csv
