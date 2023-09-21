# Cu get Ref-backend

> forked from SA-Activites @Nick

**Note**: This document is a work in progress. Changes will be made over time to refine and expand upon the instructions.

## **Prerequisites**

- Ensure you have **[Docker](https://www.docker.com/get-started)** and **[Docker Compose](https://docs.docker.com/compose/install/)** installed on your machine.

## **Getting Started**

1. **Clone the Repository**:

   If you haven't already, clone this repository to your local machine.

   ```bash
   git clone https://github.com/chanakorn-aramsak/Cu-get-Ref-backend.git
   ```

2. **Navigate to the Project Root**:

   Change directory to the root of the project:

   ```bash
   cd Cu-get-Ref-backend
   ```

3. **Start the MySQL Server with Docker**:

   With Docker and Docker Compose installed, you can easily start the MySQL server:

   ```bash
   docker-compose up
   ```

   This command will read the **`docker-compose.yml`** file in the project root and set up the services as defined. In this case, it will pull the MySQL image (if not already present) and start a container with the settings you've specified.

   If you wish to run it in detached mode (in the background), use:

   ```bash
   docker-compose up -d
   ```

4. **Connecting to the MySQL Database**:

   You can connect to the MySQL database using any MySQL client. The database will be accessible on **`localhost`** at port **`3307`**. Use the username **`user`** and the password **`password`** to connect.

   Example using the MySQL command-line tool:

   ```bash
   mysql -h 127.0.0.1 -P 3307 -u user -p
   ```

   When prompted, enter the password (**`password`**).

5. **Stopping the Services**:

   When you're done, you can stop the MySQL server and other services using:

   ```bash
   docker-compose down
   ```
