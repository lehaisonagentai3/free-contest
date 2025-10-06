import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Submission } from '../types/api';

const ResultPage: React.FC = () => {
  const [submission, setSubmission] = useState<Submission | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const navigate = useNavigate();

  useEffect(() => {
    // Get submission data from localStorage
    const submissionData = localStorage.getItem('lastSubmission');
    
    if (submissionData) {
      try {
        const parsedSubmission: Submission = JSON.parse(submissionData);
        setSubmission(parsedSubmission);
      } catch (err) {
        console.error('Error parsing submission data:', err);
      }
    }
    
    setLoading(false);
  }, []);

  const getScoreColor = (score: number): string => {
    if (score >= 8) return '#28a745'; // Green
    if (score >= 6) return '#ffc107'; // Yellow
    return '#dc3545'; // Red
  };

  const getScoreMessage = (score: number): string => {
    if (score >= 9) return 'Tuyệt vời! Thành tích xuất sắc!';
    if (score >= 8) return 'Làm tốt lắm! Xuất sắc!';
    if (score >= 7) return 'Làm tốt! Tiếp tục phát huy!';
    if (score >= 6) return 'Không tệ! Còn chỗ để cải thiện.';
    return 'Hãy học tập thêm và thử lại!';
  };

  if (loading) {
    return (
      <div className="container">
        <div className="loading">Đang tải kết quả...</div>
      </div>
    );
  }

  if (!submission) {
    return (
      <div className="container">
        <div className="card" style={{ textAlign: 'center' }}>
          <h2>Không tìm thấy kết quả</h2>
          <p>Không tìm thấy kết quả thi. Vui lòng làm bài thi trước.</p>
          <button 
            onClick={() => navigate('/select-subject')} 
            className="btn"
          >
            Làm bài thi
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="container">
      <div className="card">
        <div className="score-display">
          <h2>Kết quả thi</h2>
          
          <div className="score-number" style={{ color: getScoreColor(submission.score) }}>
            {submission.score}
          </div>
          
          <div style={{ fontSize: '18px', marginBottom: '20px', color: '#666' }}>
            {getScoreMessage(submission.score)}
          </div>
          
          <div style={{ 
            background: '#f8f9fa', 
            padding: '20px', 
            borderRadius: '8px',
            marginBottom: '30px'
          }}>
            <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit, minmax(200px, 1fr))', gap: '16px' }}>
              <div>
                <div style={{ fontWeight: 'bold', color: '#333' }}>Môn thi</div>
                <div style={{ color: '#666' }}>{submission.subject_name}</div>
              </div>
              
              <div>
                <div style={{ fontWeight: 'bold', color: '#333' }}>Điểm số</div>
                <div style={{ color: getScoreColor(submission.score), fontWeight: 'bold' }}>
                  {submission.score}/10
                </div>
              </div>
              
              <div>
                <div style={{ fontWeight: 'bold', color: '#333' }}>Thời gian nộp bài</div>
                <div style={{ color: '#666' }}>
                  {new Date(submission.submitted_at * 1000).toLocaleString('vi-VN')}
                </div>
              </div>
              
              <div>
                <div style={{ fontWeight: 'bold', color: '#333' }}>Câu đã trả lời</div>
                <div style={{ color: '#666' }}>
                  {Object.keys(submission.answers || {}).length}
                </div>
              </div>
            </div>
          </div>
          
          <div style={{ display: 'flex', gap: '16px', justifyContent: 'center', flexWrap: 'wrap' }}>
            <button 
              onClick={() => navigate('/select-subject')} 
              className="btn"
            >
              Làm bài thi khác
            </button>
            
            <button 
              onClick={() => navigate('/leaderboard')} 
              className="btn btn-success"
            >
              Xem bảng xếp hạng
            </button>
          </div>
          
          {/* Performance Analysis */}
          <div style={{ 
            marginTop: '40px', 
            padding: '20px', 
            background: '#f8f9fa', 
            borderRadius: '8px',
            textAlign: 'left'
          }}>
            <h4 style={{ marginBottom: '16px', textAlign: 'center' }}>Phân tích kết quả</h4>
            
            <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit, minmax(250px, 1fr))', gap: '16px' }}>
              <div>
                <strong>Mức điểm:</strong>
                <div style={{ marginTop: '8px' }}>
                  {submission.score >= 9 && (
                    <span style={{ color: '#28a745' }}>⭐ Xuất sắc (9-10)</span>
                  )}
                  {submission.score >= 8 && submission.score < 9 && (
                    <span style={{ color: '#28a745' }}>✅ Rất tốt (8-8.9)</span>
                  )}
                  {submission.score >= 7 && submission.score < 8 && (
                    <span style={{ color: '#ffc107' }}>👍 Tốt (7-7.9)</span>
                  )}
                  {submission.score >= 6 && submission.score < 7 && (
                    <span style={{ color: '#fd7e14' }}>⚠️ Đạt yêu cầu (6-6.9)</span>
                  )}
                  {submission.score < 6 && (
                    <span style={{ color: '#dc3545' }}>❌ Cần cải thiện (&lt;6)</span>
                  )}
                </div>
              </div>
              
              <div>
                <strong>Đề xuất:</strong>
                <div style={{ marginTop: '8px', fontSize: '14px', color: '#666' }}>
                  {submission.score >= 8 ? (
                    "Làm tốt lắm! Hãy thử các môn thi nâng cao."
                  ) : submission.score >= 6 ? (
                    "Nỗ lực tốt! Hãy ôn tập lại và luyện tập thêm."
                  ) : (
                    "Hãy tập trung học các kiến thức cơ bản và thi lại."
                  )}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ResultPage;
