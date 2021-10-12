# Introduction
---
> This project has been designed based on the hexagonal architecture, to read more https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749?gi=fb67606a28ab


### Directory layout


```
.
├── build                   # Build files for instance Dockerfile.
├── cmd                     # Application entry point
├── config                  # Application configuration
├── internal                # Golang's standard internal package
│   ├── core     
│   │   ├── domain          # The models' structure and rules for the application.
│   │   ├── port            # Act like interactors to connect to the outside of the core.
│   │   └── service         # The business logic of the application.
│   │
│   ├── handler             # The handlers, to handle the incoming requests.
│   └── repostiory          # The repositories such as Database, a microservice API exposed via gRPC or REST, or just a simple CSV file. 
│
├── pkg                     # utility packages.
└── protocol                # The protocols to serve for client.
```

---
### How to run ?

Here is the example to run the application on your local machine

```make serve-rest```

the command above is an example to make the application to serve REST