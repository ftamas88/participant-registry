# GRAIL - New Beginnings - Participant Registry
This is a participant registry microservice which supports adding, updating, removing and retrieving personal information about participants in the study.

## First steps
Make sure you have `Makefile` support on your system. (Can be installed on Windows, Linux, Mac)  

`make help` - List all available commands

If you unable or don't want to compile the `Go` files, you need to have `Docker` installed.  
For more information please visit: https://docs.docker.com/get-docker/


### Setup
Run the `make setup` command to have the linting and go tools available on your system.  
It is used to run the go fmt, imports and linting checker prior to every build and run.

In order to create the **environment** files, you should either make a copy of the `.env.dist` files or run the `make init-env` command.

### Test
There are 3 different commands available in the makefile to trigger the tests.  
`make test` runs everything, but it can take some time.  
`make test-ci` excludes the race condition check since it is NOT available on Alpine distributions.  
`make test-quick` as the name suggest only running the short tests and exits on the first failure.

### Build & Run 
Simply type `make run` command.  
(Make sure you have your `.env` files ready.)

If you don't have `Go` installed on your system, you can use `Docker` too.  
With the `make docker-up` command you can run the service without any dependency.

## Application

### Adding a new participant
Send a **POST** request to the `/api/v1/participants/create` endpoint with a sample payload like this:
```json
{
	"reference": "aa-11-bb",
	"name": "John Doe",
	"date_of_birth": "2020-09-05T18:21:25.455Z",
	"phone": "123-456-789",
	"address": {
		"address_type": "flat",
		"street_name": "Test Street",
		"post_code": "P0S C0D",
		"town_name": "Town",
		"address_line": [
			"Address line 1",
			"Address line 2"
		]
	}
}
```


### Updating an existing participant
Send a **POST** request to the `/api/v1/participants/{{reference}}/update` endpoint with a sample payload like this:
```json
{
	"reference": "aa-11-bb",
	"name": "John Doe",
	"date_of_birth": "2020-09-05T18:21:25.455Z",
	"phone": "123-456-789",
	"address": {
		"address_type": "flat",
		"street_name": "Test Street",
		"post_code": "P0S C0D",
		"town_name": "Town",
		"address_line": [
			"Address line 1",
			"Address line 2"
		]
	}
}
```