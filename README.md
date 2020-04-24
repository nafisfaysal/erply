# erply

## Development: Getting Started

### Prerequisites
* Go
* Docker
* Docker Compose

### Prepare Workspace

Clone the repository:

	$ git clone git@github.com:nafisfaysal/erply.git
	$ cd erply

Create .env file from example-env.txt:

    $ cp example-env.txt .env

Populate .env with necessary environment variable values:

    $ nano .env

Build everything with one command:

    $ make
