# Clean Architecture Sample
This is a sample implementation of clean architecture principles.
The overall architecture of the system uses Clean Architecture with reference from [here](https://threedots.tech/post/introducing-clean-architecture/)

## Getting started
#### Local deployment(in Docker)
```bash
 make run
```
The app should be available at `http://127.0.0.1:3000`.  
Postman Collection is attached for easy testing **Wallet.postman_collection.json**


A full list of available `make` commands can be found running `make help` or just `make`:
```bash
$ make help
Please use `make <target>' where <target> is one of:

    help       - Display this help
    deps       - Install dependencies
    build      - build executable
    run        - run the service with docker compose
    test       - run all test for service
    load_test  - run load test for the service
```

## Project structure
### `/cmd`
Main applications for this project.

### `/internal`
Private application and library code. This is the code you don't want others importing in their applications or libraries. Note that this layout pattern is enforced by the Go compiler itself. See the Go 1.4 [`release notes`](https://golang.org/doc/go1.4#internalpackages) for more details. Note that you are not limited to the top level `internal` directory. You can have more than one `internal` directory at any level of your project tree.

### `/pkg`
Library code that's ok to use by external applications (e.g., `/pkg/mypubliclib`). Other projects will import these libraries expecting them to work, so think twice before you put something here :-) Note that the `internal` directory is a better way to ensure your private packages are not importable because it's enforced by Go. The `/pkg` directory is still a good way to explicitly communicate that the code in that directory is safe for use by others.

### `/scripts`
Scripts to perform various build, install, analysis, etc operations.

## Package structure
All packages are located under `/internal` and `/pkg` folders. They follow the [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) principles.

A package consists of four layers:
- `adapters` a.k.a. "frameworks and drivers" layer contains the tools required to connect to external services such as caching servers, databases, SQS, etc.
- `ports` a.k.a. "controllers/presenters/gateways" contains the code responsible to to convert the received data in format that the use cases accept, send it to them, and return the response in a correct format.
- `app` layer contains the application business rules logic. Here are located all services of the package.
- `domain` holds the Entities, which are the business objects of the application. They could encapsulate the most general and high-level rules. They are the least likely to change when something external changes.
