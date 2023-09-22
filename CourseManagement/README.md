# node-grpc

Node JS GRPC Server with Express JS REST client using Mongo DB as database.


## Getting started

Make sure you have at least Node.js 10.16.0 installed.

You can check your Node.js version by running node -v:

``` console
$ node -v
v10.16.0
```


#### `clone`

Navigate to your work directory and clone the project, change directory to the `node-grpc` and add a new remote origin pointing to the new project repo.

``` console
$ git clone https://github.com/akinmaurice/node-grpc.git
$ cd node-grpc
```



#### `env`

You can set environment variables using `.env` file, or setting them in your `.bashrc` or `.zshrc` file.

##### Using `.env`:

``` console
$ touch .env
```

``` console
# using any editor you prefer. (vim/nano/vi).

$ vim .env
```

Paste the following configuration variables with their corresponding values.

```bash
DATABASE_URI=VALUE
```

##### Using `.bashrc` or `.zshrc`:

``` console
# open your bashrc or zshrc file,
# using any editor you prefer. (vim/nano/vi).

$ vim ~/.bashrc
```

Paste the following configuration variables with their corresponding values.

```bash
DATABASE_URI=VALUE
```

Run command to reload

```console
$ source ~/.bashrc
```


Run command to install dependencies

``` console
$ npm install
```


Run command to start server on port 50051

``` console
$ npm run start-server
```


open a new terminal and run command to start client on port 3000

``` console
$ npm run start-client
```

The above will get you a copy of the project up and running on your local machine for development and testing purposes.




## Example Requests


###  Get List of all Todos

> GET {{ base_uri }}/api/todos

### Response

```json

{
    "todos": [
        {
            "id": "215782",
            "title": "Test Todo",
            "completed": false
        },
        {
            "id": "380484",
            "title": "Running Man Todo",
            "completed": false
        },
        {
            "id": "541088",
            "title": "Watch football",
            "completed": false
        }
    ]
}

```


###  Get Single Todo

> GET {{ base_uri }}/api/todos/:id

### Response

```json

{
    "id": "609871",
    "title": "Test Todo",
    "completed": false
}

```


###  Create new todo

> POST {{ base_uri }}/api/todos

### Request Query
| parameters | Type | Description |
| ---------- |:-----:|-----------:|
| title    | string |  Todo Title |

### Response

```json

{
    "id": "215782",
    "title": "Test Todo",
    "completed": false
}

```



###  Delete Todo

> DELETE {{ base_uri }}/api/todos/:id


### Response

```json

{
    "status": "true"
}

```



# Used Technologies
* gRPC
* MongoDB with Mongoose
* Node JS
* Express JS


# License

This project is licensed under the MIT License - see the [LICENSE.md](https://opensource.org/licenses/MIT) file for details
