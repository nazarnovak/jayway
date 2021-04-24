# CLI mode

<img width="1277" alt="Screen Shot 2021-04-24 at 11 31 31 PM" src="https://user-images.githubusercontent.com/20066923/115973492-ab51cc80-a555-11eb-8cb0-ab906632a71a.png">

## Starting CLI mode

`cd backend`

You can run the CLI directly by running
`go run *.go -cli`

Or you can build the executable first and run it
`go build -o robot_cli`
`./robot_cli -cli`

# Server mode

<img width="1139" alt="Screen Shot 2021-04-24 at 7 11 50 PM" src="https://user-images.githubusercontent.com/20066923/115973491-a8ef7280-a555-11eb-848d-0a601acf0eff.png">

## Starting server mode

1. Start the backend by doing
`cd backend`
`go run *.go -server`

Or you can build the executable first and run it
`go build -o robot_server`
`./robot_server -server

2. Start the frontend by doing
`cd frontend` (from the root of the repo)
`npm run start`
