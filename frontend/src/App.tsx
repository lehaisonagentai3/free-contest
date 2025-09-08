import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import LoginPage from './pages/LoginPage';
import SelectSubjectPage from './pages/SelectSubjectPage';
import TestPage from './pages/TestPage';
import ResultPage from './pages/ResultPage';
import LeaderBoardPage from './pages/LeaderBoardPage';
import Navigation from './components/Navigation';

const App: React.FC = () => {
  const [officerId, setOfficerId] = React.useState<string>(
    localStorage.getItem('officerId') || ''
  );

  React.useEffect(() => {
    if (officerId) {
      localStorage.setItem('officerId', officerId);
    } else {
      localStorage.removeItem('officerId');
    }
  }, [officerId]);

  return (
    <Router>
      <div className="App">
        <Navigation officerId={officerId} setOfficerId={setOfficerId} />
        <Routes>
          <Route 
            path="/" 
            element={<LoginPage setOfficerId={setOfficerId} />} 
          />
          <Route 
            path="/login" 
            element={<LoginPage setOfficerId={setOfficerId} />} 
          />
          <Route 
            path="/select-subject" 
            element={
              officerId ? (
                <SelectSubjectPage officerId={officerId} />
              ) : (
                <Navigate to="/login" />
              )
            } 
          />
          <Route 
            path="/test/:subjectId" 
            element={
              officerId ? (
                <TestPage officerId={officerId} />
              ) : (
                <Navigate to="/login" />
              )
            } 
          />
          <Route 
            path="/result" 
            element={
              officerId ? (
                <ResultPage />
              ) : (
                <Navigate to="/login" />
              )
            } 
          />
          <Route 
            path="/leaderboard" 
            element={<LeaderBoardPage />} 
          />
        </Routes>
      </div>
    </Router>
  );
};

export default App;
