# gojailms

## Installation:
- Download Latest Release from Go here: https://go.dev/dl/ .
- Install it and open cmd and type ``go version ``. It should print the current go version installed
- Download this repository by Git. Open the Git Bash and type this `` git clone https://github.com/MakMoinee/gojailms.git ``
- After downloading, go to the project directory by typing in the same Git Bash Location `` cd gojailms ``
- After changing the directory, type ``go mod tidy`` to install the dependencies defined in the go.mod.
- After the dependencies are all downloaded, type `` cd cmd/webapp ``
- Review the settings.yaml especially the mysql configurations specifically the password of your mysql. Change the current set in settings.yaml and replace it with your own set up password for mysql.
- Create or Import existing Database with the database name ``jaildb``
- Run the command in the same Git Bash ``go run main.go``
- By Default, The server will start on port 8443. Try hitting this. ``http://localhost:8443/health``

## HTTP Request
### Retrieve Users
- GET http://localhost:8443/get/jms/users
### Save User
- POST http://localhost:8443/create/jms/user
- Request Body: 
`` {
    "userName": "sampleUser2",
    "userPassword": "sampleUser1234"
}``
### Delete User
- DELETE http://localhost:8443/delete/jms/user?id=2
- where `2` is the target id you want to delete
