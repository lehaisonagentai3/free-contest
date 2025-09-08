import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { NavigationProps } from '../types/api';

const Navigation: React.FC<NavigationProps> = ({ officerId, setOfficerId }) => {
  const navigate = useNavigate();

  const handleLogout = (): void => {
    setOfficerId('');
    localStorage.removeItem('officerId');
    localStorage.removeItem('lastSubmission');
    navigate('/login');
  };

  return (
    <nav className="nav-bar">
      <div className="nav-content">
        <h1 className="nav-title">Thi Trắc Nghiệm</h1>
        <div className="nav-links">
          <Link to="/leaderboard" className="nav-link">
            Bảng xếp hạng
          </Link>
          {officerId && (
            <>
              <Link to="/select-subject" className="nav-link">
                Chọn môn thi
              </Link>
              <span className="nav-link">Mã cán bộ: {officerId}</span>
              <button 
                onClick={handleLogout}
                className="btn btn-danger"
                style={{ padding: '6px 12px', fontSize: '14px' }}
              >
                Đăng xuất
              </button>
            </>
          )}
          {!officerId && (
            <Link to="/login" className="nav-link">
              Đăng nhập
            </Link>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navigation;
