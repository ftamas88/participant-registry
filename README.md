# Participant Registry
This is a participant registry microservice which supports adding, updating, removing and retrieving personal information about participants in the study.

## First steps
Make sure you have `Makefile` support on your system. (Can be installed on Windows, Linux, Mac)  

`make help` - List all available commands  
`make api-docs` - Run the api docs server locally

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

### Fetching an existing participant
Send a **GET** request to the `/api/v1/participants/{{reference}}` endpoint
Sample response:  
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

### Deleting an existing participant
Send a **POST** request to the `/api/v1/participants/{{reference}}/delete` endpoint


## Next Steps, Comments
- Ideally I would properly fill out the api-docs template, so it would be more useful.
- I would add a participant reference verifier middleware, so it can check the value before it gets passed to the service/repository
- Due to the nature of the sensitive data, I would AT LEAST encrypt it with application level keys and put it inside a crypto box with a
 unique salt / participants
- I would make the changes traceable and immutable. Every address or other data change will create a new entry in the system which allows
 us to enhance security, etc.
- Designing a short reference number has its own limitations and number of unique combinations. Need more thinking on how to prevent
 customer fraud or phishing.
- Normalising the model would also be highly recommended here

For the development I've used my Windows PC with the new WSL 2 (with Ubuntu base), but it should work on any system.

### Tools
- I have used mockery to easily create mocks of existing interfaces
- testify also has useful features to compare test datasets 
- I have used Viper pkg here to load the `.env` and parse values from env, however usually we use a plugin for Goland which makes that
 package not necessary at all.
- I will attach the `Insomnia` (a better Postman alternative) API collection/export in JSON 
