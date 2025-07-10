import React, { useState, useEffect } from 'react';
import { pointUsecase } from '../../usecases/pointUsecase';
import { UserPoint } from '../../domain/Point';

interface PointDisplayProps {
  userId: number;
  onBalanceUpdate?: (balance: number) => void;
}

export const PointDisplay: React.FC<PointDisplayProps> = ({ userId, onBalanceUpdate }) => {
  const [balance, setBalance] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  const fetchBalance = async () => {
    try {
      setLoading(true);
      setError(null);
      const userPoint: UserPoint = await pointUsecase.getBalance(userId);
      setBalance(userPoint.balance);
      onBalanceUpdate?.(userPoint.balance);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch balance');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (userId > 0) {
      fetchBalance();
    }
  }, [userId]);

  const refreshBalance = () => {
    fetchBalance();
  };

  if (loading) {
    return <div className="point-display loading">Loading...</div>;
  }

  if (error) {
    return (
      <div className="point-display error">
        <p>Error: {error}</p>
        <button onClick={refreshBalance}>Retry</button>
      </div>
    );
  }

  return (
    <div className="point-display">
      <div className="balance">
        <span className="label">Points:</span>
        <span className="value">{balance.toLocaleString()}</span>
      </div>
      <button className="refresh-btn" onClick={refreshBalance}>
        Refresh
      </button>
    </div>
  );
};