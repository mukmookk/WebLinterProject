# Gomunkong

Online Code Formatter.

[Demo on Okteto Cloud](https://gomunkong-frontend-ingress-indosaram.cloud.okteto.net/)



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

The following will be supported shortly.

- Swift
- C-like languages
- HTML
- CSS



## How it works

A frontend and a backend are working on Kubernetes.

### Frontend

[Repository](https://github.com/Indosaram/gomunkong-frontend)

The frontend is made by `create-react-app` . 



### Backend

[Repository](https://github.com/Indosaram/gomunkong-backend)

Pure Golang backend, with gRPC and gin-go.



### Deployment

You can deploy the app via minikube. If you use Mac, you may need to pass `--vm=true` option for using Ingress along with Minikube. Also, don't forget to enable ingress addon.

```bash
minikube start --vm=true
minikube addons enable ingress
```

Now, apply k8s definitions.

```bash
# gomunkong-main
bash deploy_k8s.sh
```

Check your Minikube ip and you are good to go!

```bash
minikube ip
```



