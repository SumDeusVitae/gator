# README.md

## Gator CLI

Welcome to the Gator CLI! This command-line interface is designed to streamline your workflow with a focus on simplicity and efficiency. 

### Prerequisites

Before you can run the Gator CLI, you'll need to have the following installed on your machine:

- **PostgreSQL**: Make sure you have PostgreSQL set up and running. You can download it from [PostgreSQL's official website](https://www.postgresql.org/download/).

- **Go**: Ensure that you have Go installed. You can get it from [Go's official website](https://golang.org/dl/).

### Installing Gator CLI

To install the Gator CLI, use the following command in your terminal:

```bash
go clone https://github.com/SumDeusVitae/gator.git
cd gator
```

### Setting Up the Config File

1. Create a configuration file named `config.yaml` in the root of your project.
2. Populate it with the necessary database connection settings and other configuration options. Here’s a sample structure:

```yaml
database:
  user: "your_db_user"
  password: "your_db_password"
  host: "localhost"
  port: 5432
  dbname: "your_db_name"
```

### Running the Program

To run the Gator CLI, execute the following command in your terminal:

```bash
go build
```
or 
```bash
go install
```

### Available Commands

The Gator CLI supports the following commands:

- `register`: Create a new user account.
- `login`: Authenticate an existing user.
- `reset`: Reset the password for a user.
- `users`: List all registered users.
- `agg`: Perform data aggregation tasks.
- `addfeed`: Add a new feed (requires login).
- `feeds`: List all available feeds.
- `follow`: Follow a specific feed (requires login).
- `following`: List all feeds you are currently following (requires login).
- `unfollow`: Unfollow a specific feed (requires login).
- `browse`: Browse available feeds (requires login).


Exanmple of usage:

```bash
gator register <name>
```
Make sure to use the appropriate commands based on whether you are logged in or not!

Feel free to reach out if you have any questions or need further assistance. Happy coding!
