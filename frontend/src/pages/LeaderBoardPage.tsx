import React, { useState, useEffect } from 'react';
import { getOfficers } from '../api/api';
import { Officer } from '../types/api';

type SortBy = 'score' | 'name' | 'unit' | 'tests';
type SortOrder = 'asc' | 'desc';

const LeaderBoardPage: React.FC = () => {
  const [officers, setOfficers] = useState<Officer[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>('');
  const [sortBy, setSortBy] = useState<SortBy>('score');
  const [sortOrder, setSortOrder] = useState<SortOrder>('desc');

  useEffect(() => {
    const fetchOfficers = async (): Promise<void> => {
      try {
        setLoading(true);
        const officersData = await getOfficers();
        setOfficers(officersData);
      } catch (err) {
        setError('Không thể tải bảng xếp hạng. Vui lòng thử lại.');
        console.error('Error fetching officers:', err);
      } finally {
        setLoading(false);
      }
    };

    fetchOfficers();
  }, []);

  const sortedOfficers = React.useMemo(() => {
    if (!officers) return [];

    return [...officers].sort((a, b) => {
      let aValue: string | number, bValue: string | number;

      switch (sortBy) {
        case 'score':
          aValue = a.score || 0;
          bValue = b.score || 0;
          break;
        case 'name':
          aValue = a.name || '';
          bValue = b.name || '';
          break;
        case 'unit':
          aValue = a.unit?.name || '';
          bValue = b.unit?.name || '';
          break;
        case 'tests':
          aValue = a.list_submission?.length || 0;
          bValue = b.list_submission?.length || 0;
          break;
        default:
          aValue = a.score || 0;
          bValue = b.score || 0;
      }

      if (typeof aValue === 'string') {
        const comparison = aValue.localeCompare(bValue as string);
        return sortOrder === 'asc' ? comparison : -comparison;
      }

      return sortOrder === 'asc' ? (aValue as number) - (bValue as number) : (bValue as number) - (aValue as number);
    });
  }, [officers, sortBy, sortOrder]);

  const handleSort = (column: SortBy): void => {
    if (sortBy === column) {
      setSortOrder(sortOrder === 'asc' ? 'desc' : 'asc');
    } else {
      setSortBy(column);
      setSortOrder(column === 'score' || column === 'tests' ? 'desc' : 'asc');
    }
  };

  const getRankIcon = (index: number): string => {
    switch (index) {
      case 0: return '🥇';
      case 1: return '🥈';
      case 2: return '🥉';
      default: return `#${index + 1}`;
    }
  };

  const getScoreColor = (score: number): string => {
    if (score >= 80) return '#28a745';
    if (score >= 60) return '#ffc107';
    return '#dc3545';
  };

  if (loading) {
    return (
      <div className="container">
        <div className="loading">Đang tải bảng xếp hạng...</div>
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

  return (
    <div className="container">
      <div className="card">
        <h2>Bảng xếp hạng</h2>
        <p>Xếp hạng tất cả cán bộ dựa trên kết quả thi</p>
        
        {/* Statistics */}
        <div style={{ 
          display: 'grid', 
          gridTemplateColumns: 'repeat(auto-fit, minmax(200px, 1fr))', 
          gap: '16px', 
          margin: '20px 0',
          padding: '20px',
          background: '#f8f9fa',
          borderRadius: '8px'
        }}>
          <div style={{ textAlign: 'center' }}>
            <div style={{ fontSize: '24px', fontWeight: 'bold', color: '#007bff' }}>
              {officers.length}
            </div>
            <div style={{ color: '#666' }}>Tổng số cán bộ</div>
          </div>
          
          <div style={{ textAlign: 'center' }}>
            <div style={{ fontSize: '24px', fontWeight: 'bold', color: '#28a745' }}>
              {officers.filter(o => (o.score || 0) >= 80).length}
            </div>
            <div style={{ color: '#666' }}>Điểm cao (≥80%)</div>
          </div>
          
          <div style={{ textAlign: 'center' }}>
            <div style={{ fontSize: '24px', fontWeight: 'bold', color: '#ffc107' }}>
              {officers.reduce((sum, o) => sum + (o.list_submission?.length || 0), 0)}
            </div>
            <div style={{ color: '#666' }}>Tổng bài thi đã làm</div>
          </div>
          
          <div style={{ textAlign: 'center' }}>
            <div style={{ fontSize: '24px', fontWeight: 'bold', color: '#17a2b8' }}>
              {officers.length > 0 ? 
                Math.round(officers.reduce((sum, o) => sum + (o.score || 0), 0) / officers.length) 
                : 0}%
            </div>
            <div style={{ color: '#666' }}>Điểm trung bình</div>
          </div>
        </div>
      </div>

      <div className="card">
        {officers.length === 0 ? (
          <p>Không tìm thấy cán bộ nào.</p>
        ) : (
          <div style={{ overflowX: 'auto' }}>
            <table className="leaderboard-table">
              <thead>
                <tr>
                  <th>Hạng</th>
                  <th 
                    onClick={() => handleSort('name')}
                    style={{ cursor: 'pointer', userSelect: 'none' }}
                  >
                    Họ tên {sortBy === 'name' && (sortOrder === 'asc' ? '↑' : '↓')}
                  </th>
                  <th>Cấp bậc</th>
                  <th>Chức vụ</th>
                  <th 
                    onClick={() => handleSort('unit')}
                    style={{ cursor: 'pointer', userSelect: 'none' }}
                  >
                    Đơn vị {sortBy === 'unit' && (sortOrder === 'asc' ? '↑' : '↓')}
                  </th>
                  <th 
                    onClick={() => handleSort('score')}
                    style={{ cursor: 'pointer', userSelect: 'none' }}
                  >
                    Điểm {sortBy === 'score' && (sortOrder === 'asc' ? '↑' : '↓')}
                  </th>
                  <th 
                    onClick={() => handleSort('tests')}
                    style={{ cursor: 'pointer', userSelect: 'none' }}
                  >
                    Bài thi đã làm {sortBy === 'tests' && (sortOrder === 'asc' ? '↑' : '↓')}
                  </th>
                </tr>
              </thead>
              <tbody>
                {sortedOfficers.map((officer, index) => (
                  <tr key={officer.id}>
                    <td style={{ fontWeight: 'bold', fontSize: '18px' }}>
                      {getRankIcon(index)}
                    </td>
                    <td style={{ fontWeight: 'bold' }}>{officer.name}</td>
                    <td>{officer.rank}</td>
                    <td>{officer.position}</td>
                    <td>{officer.unit?.name || 'N/A'}</td>
                    <td style={{ 
                      fontWeight: 'bold', 
                      color: getScoreColor(officer.score || 0) 
                    }}>
                      {officer.score || 0}%
                    </td>
                    <td>
                      {officer.list_submission?.length || 0}
                      {officer.list_submission && officer.list_submission.length > 0 && (
                        <div style={{ fontSize: '12px', color: '#666', marginTop: '4px' }}>
                          Lần cuối: {new Date(
                            Math.max(...officer.list_submission.map(s => s.submitted_at)) * 1000
                          ).toLocaleDateString('vi-VN')}
                        </div>
                      )}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
        
        <div style={{ marginTop: '20px', padding: '16px', background: '#f8f9fa', borderRadius: '4px' }}>
          <h4 style={{ marginBottom: '12px' }}>Thông tin bảng xếp hạng</h4>
          <ul style={{ margin: 0, paddingLeft: '20px', color: '#666' }}>
            <li>Xếp hạng dựa trên điểm số cao nhất mà mỗi cán bộ đạt được</li>
            <li>Nhấp vào tiêu đề cột để sắp xếp bảng</li>
            <li>🥇🥈🥉 đại diện cho 3 người có thành tích cao nhất</li>
            <li>Màu điểm: <span style={{ color: '#28a745' }}>Xanh lá (≥80%)</span>, <span style={{ color: '#ffc107' }}>Vàng (60-79%)</span>, <span style={{ color: '#dc3545' }}>Đỏ (&lt;60%)</span></li>
          </ul>
        </div>
      </div>
    </div>
  );
};

export default LeaderBoardPage;
