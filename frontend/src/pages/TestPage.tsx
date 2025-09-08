import React, { useState, useEffect, useCallback } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { getOfficerSubjectTest, startTest, submitTest } from '../api/api';
import { Test, TestAnswers, TestPageProps } from '../types/api';

const TestPage: React.FC<TestPageProps> = ({ officerId }) => {
  const { subjectId } = useParams<{ subjectId: string }>();
  const navigate = useNavigate();
  
  const [test, setTest] = useState<Test | null>(null);
  const [answers, setAnswers] = useState<TestAnswers>({});
  const [remainingTime, setRemainingTime] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(true);
  const [submitting, setSubmitting] = useState<boolean>(false);
  const [error, setError] = useState<string>('');
  const [testStarted, setTestStarted] = useState<boolean>(false);

  const fetchTest = useCallback(async (): Promise<void> => {
    if (!subjectId) return;
    
    try {
      setLoading(true);
      const testData = await getOfficerSubjectTest(officerId, subjectId);
      setTest(testData);
      
      // Check if test has already started
      if (testData.start_time && testData.start_time > 0) {
        setTestStarted(true);
        setRemainingTime(testData.remaining_time || 0);
      }
    } catch (err) {
      setError('Không thể tải bài thi. Vui lòng thử lại.');
      console.error('Error fetching test:', err);
    } finally {
      setLoading(false);
    }
  }, [officerId, subjectId]);

  useEffect(() => {
    fetchTest();
  }, [fetchTest]);

  const handleSubmit = useCallback(async (): Promise<void> => {
    if (submitting || !test) return;
    
    const unansweredQuestions = test.questions.filter(q => !answers[q.id]);
    if (unansweredQuestions.length > 0) {
      const confirm = window.confirm(
        `Bạn còn ${unansweredQuestions.length} câu hỏi chưa trả lời. Bạn có muốn nộp bài không?`
      );
      if (!confirm) return;
    }

    try {
      setSubmitting(true);
      const submission = await submitTest(officerId, test.id, answers);
      
      // Store submission data for result page
      localStorage.setItem('lastSubmission', JSON.stringify(submission));
      navigate('/result');
    } catch (err) {
      setError('Không thể nộp bài thi. Vui lòng thử lại.');
      console.error('Error submitting test:', err);
    } finally {
      setSubmitting(false);
    }
  }, [submitting, test, answers, officerId, navigate]);

  // Timer countdown
  useEffect(() => {
    if (!testStarted || remainingTime <= 0) return;

    const timer = setInterval(() => {
      setRemainingTime(prev => {
        if (prev <= 1) {
          // Auto submit when time runs out
          handleSubmit();
          return 0;
        }
        return prev - 1;
      });
    }, 1000);

    return () => clearInterval(timer);
  }, [testStarted, remainingTime, handleSubmit]);

  const handleStartTest = async (): Promise<void> => {
    if (!test) return;
    
    try {
      setLoading(true);
      const startedTest = await startTest(officerId, test.id);
      setTest(startedTest);
      setTestStarted(true);
      setRemainingTime(startedTest.remaining_time || 0);
    } catch (err) {
      setError('Không thể bắt đầu bài thi. Vui lòng thử lại.');
      console.error('Error starting test:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleAnswerChange = (questionId: number, answer: string): void => {
    setAnswers(prev => ({
      ...prev,
      [questionId]: answer
    }));
  };

  const formatTime = (seconds: number): string => {
    const hours = Math.floor(seconds / 3600);
    const minutes = Math.floor((seconds % 3600) / 60);
    const secs = seconds % 60;
    
    if (hours > 0) {
      return `${hours}:${minutes.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`;
    }
    return `${minutes}:${secs.toString().padStart(2, '0')}`;
  };

  if (loading) {
    return (
      <div className="container">
        <div className="loading">Đang tải bài thi...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="container">
        <div className="error-message">{error}</div>
        <button onClick={() => navigate('/select-subject')} className="btn">
          Quay lại chọn môn thi
        </button>
      </div>
    );
  }

  if (!test) {
    return (
      <div className="container">
        <div className="error-message">Không tìm thấy bài thi.</div>
        <button onClick={() => navigate('/select-subject')} className="btn">
          Quay lại chọn môn thi
        </button>
      </div>
    );
  }

  if (!testStarted) {
    return (
      <div className="container">
        <div className="card" style={{ textAlign: 'center' }}>
          <h2>Bài thi {test.subject?.name}</h2>
          <p>Thời gian: {test.duration ? Math.floor(test.duration / 60) : 0} phút</p>
          <p>Số câu hỏi: {test.questions?.length || 0}</p>
          <p>Cán bộ: {test.officer?.name}</p>
          
          <div style={{ margin: '30px 0' }}>
            <p><strong>Hướng dẫn:</strong></p>
            <ul style={{ textAlign: 'left', display: 'inline-block' }}>
              <li>Trả lời tất cả các câu hỏi một cách tốt nhất</li>
              <li>Bạn có thể thay đổi câu trả lời trước khi nộp bài</li>
              <li>Bài thi sẽ tự động nộp khi hết thời gian</li>
              <li>Đảm bảo bạn có kết nối internet ổn định</li>
            </ul>
          </div>
          
          <button onClick={handleStartTest} className="btn btn-success">
            Bắt đầu thi
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="container">
      {/* Timer */}
      {remainingTime > 0 && (
        <div className="timer">
          Thời gian còn lại: {formatTime(remainingTime)}
        </div>
      )}

      <div className="card">
        <h2>Bài thi {test.subject?.name}</h2>
        <p>Cán bộ: {test.officer?.name} | Số câu hỏi: {test.questions?.length || 0}</p>
      </div>

      {/* Questions */}
      {test.questions && test.questions.map((question, index) => (
        <div key={question.id} className="question-card">
          <div className="question-title">
            Câu {index + 1}: {question.content}
          </div>
          
          <div>
            {(['A', 'B', 'C', 'D'] as const).map(option => {
              const answerKey = `answer_${option.toLowerCase()}` as keyof typeof question;
              const answerText = question[answerKey] as string;
              
              if (!answerText) return null;
              
              return (
                <label key={option} className="answer-option">
                  <input
                    type="radio"
                    name={`question_${question.id}`}
                    value={option}
                    checked={answers[question.id] === option}
                    onChange={(e) => handleAnswerChange(question.id, e.target.value)}
                  />
                  <span>{option}. {answerText}</span>
                </label>
              );
            })}
          </div>
        </div>
      ))}

      {/* Submit Button */}
      <div className="card" style={{ textAlign: 'center' }}>
        <div style={{ marginBottom: '20px' }}>
          <p>
            Đã trả lời: {Object.keys(answers).length} / {test.questions?.length || 0} câu hỏi
          </p>
        </div>
        
        <button 
          onClick={handleSubmit}
          className="btn btn-success"
          disabled={submitting}
          style={{ fontSize: '18px', padding: '16px 32px' }}
        >
          {submitting ? 'Đang nộp bài...' : 'Nộp bài thi'}
        </button>
        
        <div style={{ marginTop: '20px', color: '#666' }}>
          <p>Hãy kiểm tra lại câu trả lời trước khi nộp bài.</p>
        </div>
      </div>
    </div>
  );
};

export default TestPage;
