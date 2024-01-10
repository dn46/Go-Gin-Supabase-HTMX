# Go + Supabase + HTMX Integration Example

This project is an example of how to integrate Supabase with Go and HTMX. It demonstrates how to perform CRUD operations on a book database. 

**Note:** This project does not include any design or styling besides a few spacing and positioning options.

## Prerequisites

- Go 1.16 or later
- Supabase account

## Setup

1. Clone the repository:

`
git clone https://github.com/dn46/Go-Gin-Supabase-HTMX.git
`

2. Navigate to the project directory:

`
cd yourrepository
`

3. Copy the `.env.example` file to a new file named `.env`:

`
cp .env.example .env
`

4. Open the `.env` file and replace `SUPABASE_URL` and `SUPABASE_KEY` with your Supabase URL and key.

## Running the Application

To run the application, use the following command:

`
go run main.go
`

The application will start on `localhost:8080`.

## Project Structure

- `main.go`: This file contains the main function and routing.
- `handlers.go`: This file contains the HTTP handlers.
- `models.go`: This file contains the data models.
- `db.go`: This file contains the database operations.

## Features

- List all books: `GET /books`
- Add a new book: `POST /books`
- Update a book: `POST /update/:isbn`
- Delete a book: `POST /delete/:isbn`

## Contributing

Contributions are welcome. Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.