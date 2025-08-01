import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { userUsecase } from '../../usecases/userUsecase';

export const LandingPage: React.FC = () => {
  const [userName, setUserName] = useState<string>('');
  const [isCreatingUser, setIsCreatingUser] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  const createUser = async () => {
    if (!userName.trim()) {
      setError('Please enter a name');
      return;
    }

    try {
      setIsCreatingUser(true);
      setError(null);
      const newUser = await userUsecase.createUser(userName);
      
      // ユーザー作成後、ゲーム画面にリダイレクト
      navigate(`/user/${newUser.id}`);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to create user');
    } finally {
      setIsCreatingUser(false);
    }
  };

  return (
    <div className="gacha-page">
      <div className="user-creation">
        <h1>Welcome to Fortune Spinner!</h1>
        <p>Please enter your name to start playing:</p>
        <div className="input-group">
          <input
            type="text"
            value={userName}
            onChange={(e) => setUserName(e.target.value)}
            placeholder="Enter your name"
            onKeyPress={(e) => e.key === 'Enter' && createUser()}
          />
          <button
            onClick={createUser}
            disabled={isCreatingUser || !userName.trim()}
          >
            {isCreatingUser ? 'Creating...' : 'Start Playing'}
          </button>
        </div>
        {error && <p className="error">{error}</p>}
      </div>
    </div>
  );
};