# Upwork for Tutors

A platform for connecting tutors with parents or students for physical and online learning sessions.

## Project Structure

The project consists of two main components:

### Frontend (`eagle-frontend`)
- Built with Next.js
- TypeScript for type safety
- Tailwind CSS for styling
- React components for UI

### Backend (`eagle-backend`)
- Built with Go
- MongoDB for database
- REST API endpoints

## Getting Started

### Prerequisites
- Node.js (for frontend)
- Go (for backend)
- MongoDB (for database)

### Installation

1. Clone the repository
```bash
git clone https://github.com/tise-genene/upworkfortutors.git
```

2. Install frontend dependencies
```bash
cd eagle-frontend
npm install
```

3. Set up environment variables
- Copy `.env.example` to `.env` in the backend directory
- Update the environment variables as needed

4. Start the development servers
- Start the backend server (in eagle-backend directory):
```bash
go run cmd/server/main.go
```
- Start the frontend development server (in eagle-frontend directory):
```bash
npm run dev
```

## Features

- User authentication
- Job listings for tutors
- Session booking system
- Application management
- Profile management

## Tech Stack

### Frontend
- Next.js
- TypeScript
- Tailwind CSS
- React Query
- Clerk Auth

### Backend
- Go
- MongoDB
- Gin Web Framework
- JWT Authentication

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
