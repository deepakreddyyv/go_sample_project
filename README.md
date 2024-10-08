## Go CSV Tax Calculator

### Overview

- This Go application reads a CSV file containing prices, computes the tax for each price at different rates (ranging from 0% to  40% with a step size of 10%), and stores the results in a JSON file.

### Features

1) Reads prices from a CSV file.
2) Computes tax for each price at rates: 0%, 10%, 20%, 30%, and 40%.
3) Stores the computed results in a JSON file.

### SetUp

1) Clone the Repo

- git clone https://github.com/deepakreddyyv/go_sample_project.git

- cd go_sample_project

2) Build

- go build -o tax_calculator

3) Execute

- ./tax_calculator

