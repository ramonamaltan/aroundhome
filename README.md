# Aroundhome Partner API

This API gives you information about partners that are proficient in a service.

As of now the serviceName should be `flooring` but other services can be added later.

The location of all our partners is somewhere within or in the proximity to berlin.
When making requests search for a lat/long within berlin to find partners closeby.

### API Documentation
Find detailed information on how to use this api in the [openapi spec](https://github.com/ramonamaltan/aroundhome/blob/main/openapi.yaml).
Make it look pretty pasting it into https://editor.swagger.io/.

### Running the Application

#### Prerequisites
Make sure you have Docker and Docker Compose installed on your system.
- Docker: [Installation Guide](https://www.docker.com/get-started)
- Docker Compose: [Installation Guide](https://docs.docker.com/compose/install/)

#### Clone the Repository
````
git clone https://github.com/ramonamaltan/aroundhome.git
cd your-repo
````

#### Build and Run Containers
Use Docker Compose to build the images and run the containers

````
docker-compose up --build
````

This command will:
- Build the Docker image for the Go application using the provided Dockerfile.
- Pull the PostgreSQL image from Docker Hub.
- Run the PostgreSQL container and wait until it's ready.
- Run the migration service to set up the database schema.
- Run the Go application container.

#### Verify the Setup
Access the application at http://localhost:8080.

Example Request
````
http://localhost:8080/flooring/partners?material=carpet&long=13.000&lat=53.000
````

### Notes and Assumptions
- For creating queries and db models I use [sqlc](https://docs.sqlc.dev/en/stable/)
- I made the Assumption that the Square meters of the floor and Phone number (for the partner to contact the customer)
are not relevant for GET requests but for creating customers data e.g. POST `/customer`
- If I had more time I would try to make the nearby condition part of the SQL Query
as now it's not perfect for looping through all partners to check distance (will be relevant for bigger dataset)
