# go-full-stack

Student Management System
This project is a Student Management System with a Go backend and a React frontend using Vite.
Project Structure
Copystudent-management-system/
├── backend/
│ ├── config/
│ ├── routes/
│ ├── controllers/
│ ├── main.go
│ ├── go.mod
│ └── .env
├── frontend/
│ ├── src/
│ ├── public/
│ ├── package.json
│ └── vite.config.js
└── README.md

Backend (Go)
Prerequisites

Go (version 1.16 or later)
MongoDB

Dependencies

github.com/gorilla/mux
github.com/rs/cors
go.mongodb.org/mongo-driver
github.com/joho/godotenv

Setup

Navigate to the backend directory:
Copycd backend

Install dependencies:
Copygo mod tidy

Create a .env file in the backend directory with the following content:
CopyMONGODB_URI=your_mongodb_connection_string
PORT=5555

Run the server:
Copygo run main.go

The server should start on http://localhost:5555.
Frontend (React with Vite)
Prerequisites

Node.js (version 14 or later)
npm (usually comes with Node.js)

Dependencies

react
react-dom
axios (for API calls)

Setup

Navigate to the frontend directory:
Copycd frontend

Install dependencies:
Copynpm install

Start the development server:
Copynpm run dev

The frontend should start on http://localhost:5173.
Running the Application

Start the backend server (from the backend directory):
Copygo run main.go

In a new terminal, start the frontend development server (from the frontend directory):
Copynpm run dev

Open your browser and go to http://localhost:5173 to use the application.

API Endpoints

GET /students - Fetch all students
POST /students - Create a new student
PUT /students/:id - Update a student
DELETE /students/:id - Delete a student

Contributing
Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests.
License
This project is licensed under the MIT License - see the LICENSE.md file for details.
