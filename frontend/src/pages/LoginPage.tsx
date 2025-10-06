import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { getOfficerById } from '../api/api';
import { LoginPageProps } from '../types/api';
import logo from '../logo.jpg';

const LoginPage: React.FC<LoginPageProps> = ({ setOfficerId }) => {
  const [inputOfficerId, setInputOfficerId] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>('');
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!inputOfficerId.trim()) {
      setError('Vui lòng nhập mã cán bộ');
      return;
    }

    setLoading(true);
    setError('');

    try {
      // Validate officer exists
      await getOfficerById(inputOfficerId);
      
      // Set officer ID and navigate
      setOfficerId(inputOfficerId);
      navigate('/select-subject');
    } catch (err: any) {
      if (err.response?.status === 404) {
        setError('Không tìm thấy cán bộ. Vui lòng kiểm tra lại mã cán bộ.');
      } else {
        setError('Đăng nhập thất bại. Vui lòng thử lại.');
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container">
      <div className="card" style={{ maxWidth: '400px', margin: '100px auto' }}>
        {/* Logo */}
        <div style={{ textAlign: 'center', marginBottom: '24px' }}>
          <img 
            src={logo} 
            alt="Logo" 
            style={{ 
              width: '120px', 
              height: 'auto', 
              borderRadius: '8px',
              marginBottom: '16px'
            }} 
          />
        </div>
        
        <h2 style={{ textAlign: 'center', marginBottom: '24px' }}>Đăng nhập</h2>
        
        {error && <div className="error-message">{error}</div>}
        
        <form onSubmit={handleSubmit}>
          <div className="input-group">
            <label htmlFor="officerId">Mã thí sinh</label>
            <input
              type="number"
              id="officerId"
              value={inputOfficerId}
              onChange={(e) => setInputOfficerId(e.target.value)}
              placeholder="Nhập mã thí sinh của bạn"
              disabled={loading}
              required
            />
          </div>
          
          <button 
            type="submit" 
            className="btn"
            style={{ width: '100%' }}
            disabled={loading}
          >
            {loading ? 'Đang đăng nhập...' : 'Đăng nhập'}
          </button>
        </form>
        
        <div style={{ textAlign: 'center', marginTop: '20px', color: '#666' }}>
          <p>Nhập mã thí sinh để truy cập hệ thống thi</p>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
