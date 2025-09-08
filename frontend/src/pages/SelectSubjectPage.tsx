import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { getOfficerById, getSubjects } from '../api/api';
import { Officer, Subject, SelectSubjectPageProps } from '../types/api';

const SelectSubjectPage: React.FC<SelectSubjectPageProps> = ({ officerId }) => {
  const [officer, setOfficer] = useState<Officer | null>(null);
  const [subjects, setSubjects] = useState<Subject[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>('');
  const navigate = useNavigate();

  useEffect(() => {
    const fetchData = async (): Promise<void> => {
      try {
        setLoading(true);
        const [officerData, subjectsData] = await Promise.all([
          getOfficerById(officerId),
          getSubjects()
        ]);   
        console.log('Officer Data:', officerData);
        console.log('Subjects Data:', subjectsData);
        setOfficer(officerData);
        setSubjects(subjectsData);
      } catch (err) {
        setError('Không thể tải dữ liệu. Vui lòng thử lại.');
        console.error('Error fetching data:', err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [officerId]);

  const getCompletedSubjects = (): Set<number> => {
    if (!officer?.list_submission) return new Set();
    return new Set(officer.list_submission.map(sub => sub.subject_id));
  };

  const handleSubjectSelect = (subjectId: number): void => {
    navigate(`/test/${subjectId}`);
  };

  if (loading) {
    return (
      <div className="container">
        <div className="loading">Đang tải...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="container">
        <div className="error-message">{error}</div>
      </div>
    );
  }

  const completedSubjects = getCompletedSubjects();

  return (
    <div className="container">
      <div className="card">
        <h2>Chọn môn thi</h2>
        
        {/* Officer Information */}
        {officer && (
          <div className="officer-info">
            <div className="info-item">
              <div className="info-label">Họ tên</div>
              <div className="info-value">{officer.name}</div>
            </div>
            <div className="info-item">
              <div className="info-label">Cấp bậc</div>
              <div className="info-value">{officer.rank}</div>
            </div>
            <div className="info-item">
              <div className="info-label">Chức vụ</div>
              <div className="info-value">{officer.position}</div>
            </div>
            <div className="info-item">
              <div className="info-label">Đơn vị</div>
              <div className="info-value">{officer.unit?.name || 'N/A'}</div>
            </div>
            <div className="info-item">
              <div className="info-label">Điểm hiện tại</div>
              <div className="info-value">{officer.score || 0}</div>
            </div>
            <div className="info-item">
              <div className="info-label">Số bài thi đã hoàn thành</div>
              <div className="info-value">{officer.list_submission?.length || 0}</div>
            </div>
          </div>
        )}
      </div>

      <div className="card">
        <h3>Các môn thi có sẵn</h3>
        
        {subjects.length === 0 ? (
          <p>Không có môn thi nào.</p>
        ) : (
          <div className="subject-grid">
            {subjects.map((subject) => {
              const isCompleted = completedSubjects.has(subject.id);
              const submission = officer?.list_submission?.find(sub => sub.subject_id === subject.id);
              
              return (
                <div
                  key={subject.id}
                  className={`subject-card ${isCompleted ? 'disabled' : ''}`}
                  onClick={() => !isCompleted && handleSubjectSelect(subject.id)}
                >
                  <div className="subject-title">{subject.name}</div>
                  <div className="subject-description">
                    {subject.description || 'Không có mô tả'}
                  </div>
                  <div className="subject-meta">
                    <span>
                      {subject.num_question_test || 0} câu hỏi • {subject.test_time || 0} phút
                    </span>
                    <span 
                      className={`test-status ${isCompleted ? 'completed' : 'available'}`}
                    >
                      {isCompleted 
                        ? `Đã hoàn thành (Điểm: ${submission?.score || 0})` 
                        : 'Có thể thi'
                      }
                    </span>
                  </div>
                </div>
              );
            })}
          </div>
        )}
        
        {officer?.list_submission && officer.list_submission.length > 0 && (
          <div style={{ marginTop: '30px' }}>
            <h4>Lịch sử thi</h4>
            <table className="leaderboard-table">
              <thead>
                <tr>
                  <th>Môn thi</th>
                  <th>Điểm</th>
                  <th>Thời gian nộp bài</th>
                </tr>
              </thead>
              <tbody>
                {officer.list_submission.map((submission, index) => (
                  <tr key={index}>
                    <td>{submission.subject_name}</td>
                    <td>{submission.score}</td>
                    <td>{new Date(submission.submitted_at * 1000).toLocaleString('vi-VN')}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
};

export default SelectSubjectPage;
