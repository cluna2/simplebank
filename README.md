<h1 align="center"> SimpleBank Project Design Walkthrough And Learnings   </h1>
 This is a simple Golang-based web service to create and manage money transfers between available accounts. 
 
# About
This document provides an overview of the technologies used and concepts applied through exploration of four main phases of development: Database Design, Building RESTful HTTP APIs, Deploying To Production, Applying Advanced Backend Concepts.
Key features of the project include account creation and management, inter-account transfers and 

The project spans several critical areas of backend development, including database design, API development, authorization/authentication, unit and integration testing, and session management. Additionally, the repository demonstrates proficiency in deploying applications using Docker and Kubernetes on AWS, complete with domain registration and traffic routing.

## Key Technologies and Concepts
This project incorporates a diverse set of technologies and concepts essential for backend development:

- **Programming Language**: Golang
- **Protocols**: HTTP and GRPC
- **Web Framework**: Gin
- **API Documentation**: Swagger
- **Database Design and Interaction**: DBML, PostgreSQL, SQLC
- **Security and Authentication**: JWT, PASETO, HTTPS, TLS
- **Containerization and Orchestration**: Docker, Kubernetes
- **Cloud Computing**: AWS - ECR, EKS, Secrets Manager, IAM
- **Continuous Integration/Deployment**: GitHub Actions
- **Testing**: Gomock, Testify


## Details 

### Database Design, Connection, and Communication

<details> 
 <summary> Creating Schema </summary>
 <p> The database schema was created with dbdiagram.io. The initial schema consists of three tables: 
  <ul>
    <li> Accounts: Each account has a unique owner, a set currency type, and a balance. </li>
    <li> Transfers: Each row of this table references a sender account and receiver account, and the amount transferred.</li>
    <li> Entries: Each entry records changes in an account's balance. </li>
  </ul>
  <img width="1214" height="337" alt="SimpleBank_Schema_1 (1)" src="https://github.com/user-attachments/assets/561693fa-ed80-4ad0-8891-597ae4c307d1" />

 </p>
</details>


<details> 
  <summary> Docker + Postgres images </summary>
  <p> Docker is used to pull a Postgres image that will serve as the database. A separate image is created for the service that is used for deployment later.
  I use the TablePlus GUI to visually inspect changes to the database and keep track of migrations. </p>
  <img width="1346" height="786" alt="image" src="https://github.com/user-attachments/assets/be3d7c14-3feb-4d72-9927-5f325236ce27" />

</details>

<details>
  <summary> Makefile </summary>
  <p> A Makefile is used to automate several tasks related to the project, such as pulling and starting Docker images, create and migrate the database, and compiling the service. This also has the benefit of making collaboration in a team setting easier by reducing setup time to running a few .PHONY targets. </p>
 
</details>

<details>
  <summary> Database Migrations </summary>
  <p> Typically in enterprise applications, business requirements change and induce updates to the database schema. To keep track of migrations, I use Golang's migrate library. Updates to the schema in SQL can be found in db/migration. </p>
</details>


<details>
 <summary> CRUD code from SQL </summary>
 <p> Various libraries exist to communicate with the database in Go such as database/sql, sqlx, and Gorm. For this project, I use sqlc which allows me to take SQL queries and generate type-safe Golang code\ from them. Then my application code can call those methods that sqlc generated. This not only simplifies database communication like ORMs typically do,
 but also helps catch incorrect SQL queries early during compliation. </p>
</details>

<details> 
  <summary> Unit Tests for CRUD operations. </summary>
  <p> Writing unit tests for the database operations was done using Go's testing package and Testify's require package (for assertions). Later unit tests for APIs will make use of mocking. </p>
</details>

<details> 
  <summary> CI/CD setup with Github Actions. </summary>
  <p> Github Actions allowed me to create a CI/CD pipeline by automating some tasks to test and eventually deploy the service. To start, with I created some .yaml files to pull in a Postgres service, Go, and the migrate library before running the unit tests I created. Later, I will expand this pipeline using AWS. The .yaml files are found under .github/workflows </p>
</details>

<details> 
  <summary> DB Transaction Locks and Handling Deadlocks </summary>
  <p> Deadlocks were occuring when running certain queries to update account balances or when performing a transfer. To combat this, I updated some SQL queries with the 'FOR NO KEY UPDATE' clause to inform Postgres to not modify foreign keys. In addition, I force a consistent ordering on queries to first lock accounts with a smaller ID before locking accounts with larger IDs. This reduces the risk of different transactions attempting to simultaneously lock mutually referenced accounts which casuses deadlocks. Enforcing a 'Repeatable Read' isolation level also helps maintain data consistency, which is critical in banking applications like this service. </p>
</details>

### 2. Building RESTful HTTP APIs


<details>
    <summary>Web framework</summary>
    <p> For this project, I used the Gin web framework for its ease of use for HTTP requests. I created APIs for creating accounts, updating an account's balance, and making transfers. These APIs called on the database store generated by sqlc. Additional APIs were later added along with GRPC equivalents. </p>
</details>

<details>
    <summary> Loading env vars with Viper </summary>
    <p> 
</details>
