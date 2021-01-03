# What is Lollipop?

Lollipop is cloud native performance testing tool built for making performance testing a productive and enjoyable experience.
Using Lollipop, you'll be able to catch performance regression and problems earlier, allowing you to build resilient systems and robust applications.

# Key features

Lollipop is packed with features, which includes:

* Developed in Go for high performance
* Powered by Kubernetes cloud platform
* Very user-friendly performance testing management interface
* Comprehensive data analysis reports

# Installation

## Prerequisite

### Installing Docker Desktop

**⚠ HINT: Ignore this step if you already have a K8s cluster or don't plan to use it on your own Desktop.**  

First make sure you have the desktop version of Docker installed. 
If you haven't installed it yet, please open [this page](https://www.docker.com/products/docker-desktop) and follow the instructions.

Once installed you can follow the instructions on [this page](https://docs.docker.com/docker-for-mac/#kubernetes) to enable Kubernetes on Mac, 
or follow the instructions on [this page](https://docs.docker.com/docker-for-windows/#kubernetes) to enable Kubernetes on Windows.

Once enabled, you can use ```kubectl version --short``` to check the version of Kubernetes.

### Installing Helm

**⚠ HINT: Ignore this step if you are already using Helm.**

The Helm tool is required for the installation of Lollipop, and the components it depends on. 
If you have not installed Helm, please open [this page](https://helm.sh/docs/intro/install/) and follow the instructions.

Once installed, you can use ```helm version --short``` to check the version of Helm.

### Installing K8s Ingress Controller

**⚠ HINT: Ignore this step if your K8s cluster already has an Ingress controller or if you are using another type of gateway.**

Lollipop provides a web-based management platform. In order to use this feature, the Ingress controller must be installed on k8s.

First create a K8s namespace for NGINX Ingress controller:

```kubectl create ns ingress-nginx```

NGINX Ingress controller can be installed via Helm:

```
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx/
helm install ingress-nginx ingress-nginx/ingress-nginx -n ingress-nginx
```

To check if the ingress controller pods have started, run the following command:

```
kubectl get pods -n ingress-nginx -l app.kubernetes.io/name=ingress-nginx --watch
```

### Creating Namespaces

Creates **l6p-system** and **l6p-space** K8s namespaces for the system and runtime components of Lollipop:

```
kubectl create ns l6p-system 
kubectl create ns l6p-space
```

### Installing MongoDB

**⚠ HINT: Ignore this step if you already have MongoDB installed in your K8s cluster.**

Download or **git clone** this project and go to the "utils/mongodb" directory.
Please change the **values.yaml** file according to your requirements, and run the command:

```
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install mongodb -n l6p-system -f ./values.yaml bitnami/mongodb
```

### Installing Kafka

Lollipop uses Kafka as a message queue to store the logs generated during testing.

Download or **git clone** this project and go to the "utils/kafka" directory.
Please change the **values.yaml** file according to your requirements, and run the command: 

```
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install kafka -n l6p-system -f ./values.yaml bitnami/kafka
```

## Installing Lollipop

Download or **git clone** this project and go to the "charts/lollipop" directory. 
Open and edit the **values.yaml** file and modify the configuration as required.

| Parameter | Description | Default |
| --------- | ----------- | ------- |
| global.mongodb.host | MongoDB address and port in the K8s cluster | mongodb.l6p-system.svc.cluster.local:27017 |
| global.mongodb.user | MongoDB username | root |
| global.mongodb.pass | MongoDB password | rootpassword |
| global.kafka.endpoint | Kafka endpoint | kafka.l6p-system.svc.cluster.local:9092 |
| global.kafka.topic | Kafka topic | l6p.log |
| server.storageClass | Storage class of backing PVC | hostpath |
| server.service.port | Port of the API server | 80 |
| server.ingress.hosts[0].host | Hostname of the API server | local.l6p.io |
| server.ingress.hosts[0].path | Path within the API url structure | "/api/v1" |
| web.apiBaseUrl | Base URL of API | "http://local.l6p.io/api/v1" |
| web.service.port | Port of the Web server | 80 |
| web.ingress.hosts[0].host | Hostname of the Web server | local.l6p.io |
| web.ingress.hosts[0].path | Root path of the Web server | "/" |

**⚠ HINT: Please change "local.l6p.io" in the configuration to your actual domain name, such as "l6p.mycompany.com", and resolve the domain name to the external IP address of Ingress Controller. The external IP address of Ingress Controller can be retrieved by using the command "kubectl get service -n ingress-nginx". If you are only running on desktop, you can add a record in "/etc/hosts" to point "local.l6p.io" to the external IP address of Ingress Controller.**

After modifying the configuration, please go back to the **root directory** and run:

```
helm install l6p ./charts/lollipop -n l6p-system
```

To check if the Lollipop pods have started, run the following command:

```
kubectl get pods -n l6p-system --watch
```

When all pods are running and ready it means that the system has installed successfully.

# Initialize and configure Lollipop

Open the Chrome browser and access the home page of the Lollipop web console.
For example, according to the default configuration, the home page address is "http://local.l6p.io".

First, you will be prompted to enter the Activation Code. 
We provide you with a free Activation Code for up to **100 virtual users**, enough for general performance testing.
If you need more virtual users, please email us at <l6p.io@hotmail.com>, and we will generate a new activation code for you.

```text
-----BEGIN ACTIVATION CODE-----
eyJSIjoiZXlKTllYaFdWWE5sY25NaU9qRXdNSDA9IiwiUlNpZyI6Ik14Tm9wWXRF
MW9CTGV1NXMyS29PbjJDUDhQeUNsWEUwOWNmNU9VMmV2RXp2WGlkN2c1WG01YTl2
QmlQQi9ySGRxT1JHcUZDVCs0cmV5cll6YnJhZFJoLzFlS1RjYjhRc1ArbzZhcFF6
K0NkRVpjTXplTitvK1N4RGQ5VU1LMG8xMW5oQVZ0NlVlM3JWemp2M1RXc3EwTThL
SXprUWhPbk1zSERFS3lmQWF4dz0iLCJDUCI6Ii0tLS0tQkVHSU4gQ0VSVElGSUNB
VEUtLS0tLVxuTUlJQjd6Q0NBVmlnQXdJQkFnSUJZekFOQmdrcWhraUc5dzBCQVFz
RkFEQWtNUTh3RFFZRFZRUUtFd1pzTm5BdVxuYVc4eEVUQVBCZ05WQkFNVENHeHZi
R3hwY0c5d01CNFhEVEl3TVRJeE5UQXdNREF3TUZvWERUTXdNVEl4TlRBd1xuTURB
d01Gb3dKREVQTUEwR0ExVUVDaE1HYkRad0xtbHZNUkV3RHdZRFZRUURFd2hzYjJ4
c2FYQnZjRENCbnpBTlxuQmdrcWhraUc5dzBCQVFFRkFBT0JqUUF3Z1lrQ2dZRUFt
OXBqU0JwYVhRZlNjWERTWDd3d0VhOFJ6SjY1RVlVcVxuWWFpc0JUZ1dUK1hOTE4r
RldtSmxFSUdLNkU1a05FVXRnU0dBR0JpVFlDYVVkck5xekpJelpmRTQ5TW0rcFp3
TVxuQ3BaZHEwTisxdkhZeEYvZ3pYNkZ2cThQNlVQTVg2d3M5QU5LdFRpUkFEV0dq
dWEwTVdjVWp5eFdnTk0wTkxvRlxud2xzd2RDa25zcEVDQXdFQUFhTXhNQzh3RGdZ
RFZSMFBBUUgvQkFRREFnZUFNQjBHQTFVZEpRUVdNQlFHQ0NzR1xuQVFVRkJ3TUNC
Z2dyQmdFRkJRY0RBVEFOQmdrcWhraUc5dzBCQVFzRkFBT0JnUUNrRVdjd2pQbVlB
dWFjK0x4QVxuSXQ5bTQzWUJkTU5JdlQ2azhKL1RiQXBLajV1MWRPSWNCS2pSeFNt
YjdsbXU0c0grSmFZdk0vUEtDcjFDc0pmNVxuWXBQZXNKcjdaVGlPVExqdXZmQVEw
dkRtY0FUWkpWVnZyN0U2U2drNGk2U1VuVmZHYldLK21QYWhLUkVXdlpONlxuZ2l0
ZUMxcmtPdXlxY1pHZFBlMVpHNGJJR1E9PVxuLS0tLS1FTkQgQ0VSVElGSUNBVEUt
LS0tLVxuIn0=
-----END ACTIVATION CODE-----
```

Next you will be asked to create an administrator account, after which you will be able to log in and use Lollipop.

After entering the system, please first enter the **Settings** menu to make the necessary system settings.

# Test Script Examples

In the **examples** folder we provide examples of test scripts ranging from API tests to browser based end-to-end tests.

We also provide two examples to demonstrate how to run and debug locally while creating the test cases.
