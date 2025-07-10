import React, { useState, useEffect } from 'react';
import { PointDisplay } from '../components/PointDisplay';
import { GachaSpinner } from '../components/GachaSpinner';
import { GachaHistory } from '../components/GachaHistory';
import { userUsecase } from '../../usecases/userUsecase';
import { User } from '../../domain/User';
import { GachaResult } from '../../domain/Gacha';

export const GachaPage: React.FC = () => {
  const [user, setUser] = useState<User | null>(null);
  const [userName, setUserName] = useState<string>('');
  const [isCreatingUser, setIsCreatingUser] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [historyRefreshTrigger, setHistoryRefreshTrigger] = useState<number>(0);

  const createUser = async () => {
    if (!userName.trim()) {
      setError('Please enter a name');
      return;
    }

    try {
      setIsCreatingUser(true);
      setError(null);
      const newUser = await userUsecase.createUser(userName);
      setUser(newUser);
      setUserName('');
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to create user');
    } finally {
      setIsCreatingUser(false);
    }
  };

  const handleGachaComplete = (result: GachaResult) => {
    // ガチャ完了後にポイント表示と履歴を更新
    setHistoryRefreshTrigger(prev => prev + 1);
  };

  if (!user) {
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
  }

  return (
    <div className="gacha-page">
      <header className="page-header">
        <h1>Fortune Spinner</h1>
        <p>Welcome, {user.name}!</p>
      </header>

      <div className="game-area">
        <div className="main-section">
          <PointDisplay
            userId={user.id}
            onBalanceUpdate={() => {}} // ポイント更新時の処理が必要な場合
          />
          <GachaSpinner
            userId={user.id}
            onGachaComplete={handleGachaComplete}
          />
        </div>

        <div className="sidebar">
          <GachaHistory
            userId={user.id}
            refreshTrigger={historyRefreshTrigger}
          />
        </div>
      </div>
    </div>
  );
};