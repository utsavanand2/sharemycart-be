# ShareMyCart GoLang backend

## Building the server binary
```sh
make apigen
```

## Build the Dockerfile

```sh
make docker
```

## Run the docker image

```sh
# Export the contents of the firebase service account JSON file to an environment variable
export SVC=$(cat path/to/your/firebaseServiceAccount.json)
docker run --rm -d -p 8080:8080 utsavanand2/sharemycart:latest --svc=$SVC
```