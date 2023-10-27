# Hospital Patient Queue System

Hospital Patient Queue System is a simple prototype of a patient queue management system designed for hospital receptionists. The application provides a command-line interface for managing patient admissions and dispatching based on different modes.

## Description

- This application allows receptionists to add patients to the queue with MRNumber and gender.
- Dispatch patients in a first-in, first-out (FIFO) order.
- Switch between FIFO and round-robin dispatch modes based on gender.
- Command-line interface for receptionists.
- Error handling for invalid inputs.

## Requirements

- Node.js or Go installed on your computer

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/yogapermana76/hospital-queue-system-test.git
   ```

2. Navigate to the project directory:
   ```shell
   cd hospital-queue-system-test
   ```

## Usage

1. To start the application, run the following command:

   ```shell
   node queue-system.js
   ```

   or

   ```shell
   go run queue-system.go
   ```

   Follow the prompts and use the commands described below to manage the patient queue.

## Command

- IN <MRNumber> <Gender>: Add a patient to the queue.
  Example: IN MR2015 F
- OUT: Dispatch a patient from the queue based on the current mode.
- ROUNDROBIN: Switch to round-robin dispatch mode.
- DEFAULT: Switch back to the default FIFO dispatch mode.
- EXIT: Quit the application.
