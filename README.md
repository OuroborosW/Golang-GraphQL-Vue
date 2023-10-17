## Golang GraphQL Backend with MongoDB
This repository, Golang-GraphQL-Vue, contains a Golang-based backend that provides a GraphQL API for fetching state names. The data for states is stored in MongoDB.

# Prerequisites
Golang (latest version recommended)<br />
Docker and Docker Compose<br />
Git (for cloning the repository)<br />
# Getting Started
Clone the repository:
```bash
# git clone https://github.com/OuroborosW/Golang-GraphQL-Vue.git
# cd Golang-GraphQL-Vue
Start MongoDB with Docker:
Make sure Docker and Docker Compose are properly installed on your system. Run the following command to start a MongoDB container:

bash
Copy code
docker-compose up -d
Initialize the database:
This step involves populating the MongoDB with the provided states.json data.

bash
Copy code
go run initdata.go
Run the GraphQL server:
Navigate to the main directory:

bash
Copy code
cd main
Then run:

bash
Copy code
go run *.go
The server will start, and the GraphQL endpoint will be available at http://localhost:8080/graphql.

Use the GraphQL API:
To fetch state names, send a POST request to the GraphQL endpoint with a query like:

graphql
Copy code
{
  search(input: "A") {
    name
  }
}
This will return all state names starting with the letter 'A'.

Development
The main logic for the application is found within the main directory. You can expand the functionality by adding more resolvers or modifying the existing GraphQL schema. Ensure any new dependencies are added to the go.mod file.

Cleanup
To stop the MongoDB container and remove it, run:

bash
Copy code
docker-compose down
Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

License
This project is licensed under the MIT License. See the LICENSE file for details.

