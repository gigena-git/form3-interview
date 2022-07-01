# Form3 Go client library
Form3 take home exercise, by [Maximiliano Gigena](mailto:maximiliano@gigena.xyz). This is my first time writing a project in Go. From this assignment, I've acquired some familiarity with the language, but I have not yet grasped all of its features.

In order to separate the docker containers from my local network, I created the subnet `form3` inside `docker-compose.yml`

The approach I used to implement the solution was inspired by [How Do You Structure Your Go Apps?](https://www.youtube.com/watch?v=1rxDzs0zgcE), in which the main idea is that dependencies should  point inward.
This means that if you want to make changes to the core domain of your app, you should only change that and not how your app interacts with the outside world.

Thus, the model the app uses is the following:
```
form3-interview/
├─ accounts/        -- This is the entry point of the app.
├─ pkg/
│  ├─ create/       -- This contains structs associated with the create request.
│  ├─ delete/       -- This contains structs associated with the delete request.
│  ├─ fetch/        -- This contains structs associated with the fetch request.
│  ├─ message/      -- This contains structs associated parsing requests to the mock api.
│  ├─ rest/         -- This contains structs, error types and functions that handle http requests.
```
## Requirements
To run this, you will need to have [Docker Engine](https://docs.docker.com/engine/install/), [Docker Compose](https://docs.docker.com/compose/install/), and [Go](https://go.dev/dl/) installed on your machine.
## Usage:
First, spin up the mock server with:
```
sudo docker compose up
```
Second, create a Go project and import the library:
```go
import "github.com/gigena-git/form3-interview/accounts"
```
### Create
To create an account, instantiate a Create request:
```go
import "github.com/gigena-git/form3-interview/pkg/create"

func main() {
    attributes := create.AccountAttributes{
		Country:                 "GB",
		BaseCurrency:            "GBP",
		BankID:                  "400300",
		BankIDCode:              "GBDSC",
		Bic:                     "NWBKGB22",
		Name:                    []string{"John Doe"},
		AlternativeNames:        []string{"XRIEESSWVW"},
		AccountClassification:   "Personal",
		JointAccount:            false,
		AccountMatchingOptOut:   false,
		SecondaryIdentification: "SSQRGELQGT",
	}
    create_request := accounts.AccountDataRequest {
        Type:           "accounts",
		ID:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
        Version:        int64(0),
    }
```
and then call the Create function, with said request.
```go
accounts.Create(create_request)
```
### Fetch
To get an account from the api, instantiate a Fetch request:
```go
import "github.com/gigena-git/form3-interview/pkg/fetch"

fetch_request := &fetch.AccountDataRequest{
	ID: "13b97dcd-eff2-4e30-ae0b-54629d70e56b",
}
```
and then call the Fetch function, with said request.
```go
res, _ := accounts.Fetch(fetch_request)
```
### Delete
To delete an account, instantiate a Delete request:
```go
import "github.com/gigena-git/form3-interview/pkg/delete"

delete_request := &delete.AccountDataRequest{
	ID:      "13b97dcd-eff2-4e30-ae0b-54629d70e56b",
	Version: int64(0),
}
```
and then call the Delete function, with said request:
```go
accounts.Delete(delete_request)
```
## Testing
The tests are located in the test folder. To run the tests by yourself, do
```bash
sudo docker compose up
```
this will run the tests, after spinning up the mock server and the db.

## To do
Things that are pending to implement:
- Implement client side validation (i.e: testing that the values sent in a creation request are actually created in the server, test for other status codes, etc).
- Implement libraries for http request, mock data generation and error types.
- Running the test container outside the form3 network specified in docker-compose.yml, so that I can test if the api can connect to a remote server (via the SERVER_PATH variable).