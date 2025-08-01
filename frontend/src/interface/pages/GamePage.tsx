import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { PointDisplay } from '../components/PointDisplay';
import { GachaSpinner } from '../components/GachaSpinner';
import { GachaHistory } from '../components/GachaHistory';
import { userUsecase } from '../../usecases/userUsecase';
import { User } from '../../domain/User';
import { GachaResult } from '../../domain/Gacha';

export const GamePage: React.FC = () => {
  const { userId } = useParams<{ userId: string }>();
  const navigate = useNavigate();
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [historyRefreshTrigger, setHistoryRefreshTrigger] = useState<number>(0);

  // ユーザー情報を取得
  useEffect(() => {
    const fetchUser = async () => {
      if (!userId) {
        navigate('/');
        return;
      }

      try {
        setIsLoading(true);
        setError(null);
        const userIdNum = parseInt(userId);
        if (isNaN(userIdNum) || userIdNum <= 0) {
          throw new Error('Invalid user ID');
        }
        const fetchedUser = await userUsecase.getUserById(userIdNum);
        setUser(fetchedUser);
      } catch (err) {
        console.error('Failed to fetch user:', err);
        const errorMessage = err instanceof Error ? err.message : 'Unknown error';
        setError(`User not found (${errorMessage}). Redirecting to home...`);
        // 3秒後にホームにリダイレクト
        setTimeout(() => navigate('/'), 3000);
      } finally {
        setIsLoading(false);
      }
    };

    fetchUser();
  }, [userId, navigate]);

  const handleGachaComplete = (result: GachaResult) => {
    // ガチャ完了後にポイント表示と履歴を更新
    setHistoryRefreshTrigger(prev => prev + 1);
  };

  if (isLoading) {
    return (
      <div className="gacha-page">
        <div className="loading">
          <h2>Loading...</h2>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="gacha-page">
        <div className="error">
          <h2>Error</h2>
          <p>{error}</p>
        </div>
      </div>
    );
  }

  if (!user) {
    return null;
  }

  return (
    <div className="gacha-page">
      <header className="page-header">
        <h1>Fortune Spinner</h1>
        <p>Welcome, {user.name}!</p>
        <button 
          className="new-game-button"
          onClick={() => navigate('/')}
        >
          New Game
        </button>
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