
# Go component

This is the backend part of application. This contains:


## CICD 

```Github Actions ``` as CICD which will trigger when any push occur in this repository and its yaml located inside the ```.github/workflows/```.

## Coding Standards
In the CICD, I have:
```bash
-> Implemented GolangCI-Lint for the Go component.
-> Configure GolangCI-Lint to enforce Go coding standards.
-> If the tests fails then pipeline will fail
```


## Installation

Install the project with docker compose.

Step 1. Clone the main repository
```bash
git clone https://github.com/chndel-abhishek/go-component.git
```
Step2. Initialize the go modules \
``` go mod init ```

Step3: Start the server \
``` go run  main.go ``` \
Once the compose file is up you can access the application:

``` backend api: http://localhost:8080/api ``` 


