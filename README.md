# E-Ticket Terminal

Simple e-ticket-terminal app, built with this [template](https://github.com/wildanfaz/go-template)

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Documentations](#documentations)

## Installation

1. Make sure you have Golang installed. If not, you can download it from [Golang Official Website](https://go.dev/doc/install).

2. Install 'make' if not already installed. 

    * On Debian/Ubuntu, you can use:

    ```bash
    sudo apt-get update
    sudo apt-get install make
    ```

   * On macOS, you can use [Homebrew](https://brew.sh/):

    ```bash
    brew install make
    ```

   * On Windows, you can use [Chocolatey](https://chocolatey.org/):

    ```bash
    choco install make
    ```

3. Clone the repository:

    ```bash
    git clone https://github.com/wildanfaz/e-ticket-terminal.git
    ```

4. Change to the project directory:

    ```bash
    cd e-ticket-terminal
    ```

## Usage

1. Start the application using docker:

    ```bash
    docker-compose up
    ```

## Commands

1. Install all dependencies
    ```bash
    make install
    ```

2. Start the application without docker
    ```bash
    make start
    ```

3. Migrate tables
    ```
    make migrate
    ```

4. Rollback tables
    ```
    make rollback
    ```

## Documentations

1. [System Design or Flowchart](https://drive.google.com/file/d/1W2v4QLf-vAwTAnOga2z-2Bww7ZEEhJ4l/view)

2. [Database Design or Entity Relationship Diagram](https://dbdiagram.io/d/Terminals-6573204b56d8064ca0a7d54a)

3. [Postman](https://documenter.getpostman.com/view/22978251/2s9YkgE62o)

4. [Current Database Seed](https://github.com/wildanfaz/e-ticket-terminal/tree/main/migrations)