# Notes 
Untuk web & consumer mq tidak dapat dilakukan di file yang sama (**main.go**)

Harus dibedakan dockerfile maupun service antara web & consumer
**Example**
`127.0.0.1:8000` Web Service
`127.0.0.1:8001` Consumer Web Service

Akan tetapi dengan 1 image yang sama diatur di CI/CD Action

