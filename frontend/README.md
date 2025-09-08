# Free Contest Frontend

A React.js web application for managing officer tests and contests.

## Features

- **Login Page**: Simple officer ID-based authentication
- **Subject Selection**: View officer information and select available subjects for testing
- **Test Taking**: Interactive test interface with timer and question navigation
- **Results Display**: Test score and performance analysis
- **Leaderboard**: Ranking of all officers based on performance

## Prerequisites

- Node.js (version 14 or higher)
- npm or yarn
- Backend API server running on http://localhost:8080

## Installation

1. Clone the repository and navigate to the frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
```

3. Start the development server:
```bash
npm start
```

The application will open in your browser at http://localhost:3000

## API Configuration

The application is configured to proxy API requests to http://localhost:8080. Make sure your backend server is running on this port.

If you need to change the API endpoint, update the `proxy` field in `package.json`:

```json
{
  "proxy": "http://your-api-server:port"
}
```

## Available Scripts

- `npm start` - Runs the app in development mode
- `npm test` - Launches the test runner
- `npm run build` - Builds the app for production
- `npm run eject` - Ejects from Create React App (one-way operation)

## Application Structure

```
src/
├── api/
│   └── api.js              # API service functions
├── components/
│   └── Navigation.js       # Navigation component
├── pages/
│   ├── LoginPage.js        # Officer ID login
│   ├── SelectSubjectPage.js # Subject selection and officer info
│   ├── TestPage.js         # Test taking interface
│   ├── ResultPage.js       # Test results display
│   └── LeaderBoardPage.js  # Officers leaderboard
├── App.js                  # Main application component
├── index.js                # Application entry point
└── index.css               # Global styles
```

## Usage

1. **Login**: Enter your Officer ID to access the system
2. **Select Subject**: View your profile and choose a subject to test
3. **Take Test**: Answer questions within the time limit
4. **View Results**: See your score and performance analysis
5. **Leaderboard**: Compare your performance with other officers

## API Endpoints Used

- `GET /api/v1/officers` - Get all officers
- `GET /api/v1/officers/{id}` - Get officer by ID
- `GET /api/v1/subjects` - Get all subjects
- `GET /api/v1/tests/officer-subject` - Get test for officer and subject
- `POST /api/v1/tests/start` - Start a test
- `POST /api/v1/tests/submit` - Submit test answers

## Browser Support

This application supports all modern browsers that support ES6+ features.

## License

This project is licensed under the Apache 2.0 License.
