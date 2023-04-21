# KubePack

This is the repository for our submisssion for a course project. (Software Engineering)

## About

KubePack will be a tool that automates the deployment process of applications to a production environment by providing a streamlined and efficient way to manage and deploy the source code.

## Features:
- Resource Management: KubePack helps manage and allocate computing resources such as CPU, memory, and storage for containerized applications running on Kubernetes clusters. It ensures that resources are used efficiently and effectively, and can scale up or down as needed.

- Configuration Management: KubePack makes it easy to manage application configurations and environment variables. This helps ensure that applications are deployed with the correct settings, reducing the risk of errors and inconsistencies.

- Service Discovery and Load Balancing: KubePack provides a built-in service discovery and load balancing mechanism, which enables applications to communicate with each other and share data seamlessly. It helps ensure that traffic is distributed evenly across multiple instances of an application.

- Rolling Updates and Rollbacks: KubePack enables rolling updates and rollbacks of applications, which means that new versions of an application can be deployed incrementally while the old version is still running. If any issues arise, the deployment can be easily rolled back to the previous version.

- Monitoring and Logging: KubePack provides built-in monitoring and logging capabilities to help track application performance and identify issues. It can integrate with popular logging and monitoring tools, such as Prometheus and Grafana, to provide deeper insights into application health and performance.

## Why is KubePack relevant?

KubePack is relevant because it simplifies the complex process of deploying and managing applications on Kubernetes clusters, while ensuring consistency and reliability. By providing a comprehensive set of features and tools such as resource management, configuration management, service discovery, rolling updates and rollbacks, and monitoring and logging, KubePack helps organizations save time and reduce the risk of errors in their deployment process.

## How to start?

- Clone the repository to your local system using `git clone -b` along with branch name and SSH Key
- Run the terminal in that folder and run the following command `go install ./...`
- This will install all the packages from go.mod file
- Now, run the command `go build`
- This will create a binary named `CSN-254`
- Test the binary by executing the command `./CSN-254`

