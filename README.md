# galc - Google Accounts List Control

1. Build binary:

```go build```

2. Export vars `GCP_API_KEY` and `PROJECT_ID`

```export GCP_API_KEY=$(gcloud auth print-access-token)```

```export PROJECT_ID="project-id"```

# How to use
```./galc --type Role```

```./galc --type ServiceAccount EMAIL``` 

```./galc --type ServiceAccounts```

Have fun!
![coffee](https://github.com/tonnytg/galc/img/pix.jpg)