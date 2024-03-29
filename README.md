# WebLinterProject

Online Code Formatter.


## Language support

For now, four languages are available with their formatters.

- Python3
  - black
  - yapf
  - autopep8
- Golang
  - gofmt
- Javascript
  - prettier
  - js-beautify
- Java
  - google-java-format

## How it works

A frontend and a backend are working on Kubernetes.

### Frontend

The frontend is made by `create-react-app` . 

### Backend

Pure Golang backend, with gRPC and gin-go.

### Deployment

You can deploy the app via minikube. If you use Mac, you may need to pass `--vm=true` option for using Ingress along with Minikube. Also, don't forget to enable ingress addon.

```bash
minikube start --vm=true
minikube addons enable ingress
```

Now, apply k8s definitions.

```bash
# WebLinterProject-main
bash deploy_k8s.sh
```

Check your Minikube ip and you are good to go!

```bash
minikube ip
```



