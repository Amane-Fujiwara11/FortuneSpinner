import React, { useState, useEffect } from 'react';
import { gachaUsecase } from '../../usecases/gachaUsecase';
import { GachaHistory as GachaHistoryType, Rarity } from '../../domain/Gacha';

interface GachaHistoryProps {
  userId: number;
  refreshTrigger?: number;
}

export const GachaHistory: React.FC<GachaHistoryProps> = ({ userId, refreshTrigger }) => {
  const [history, setHistory] = useState<GachaHistoryType[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  const fetchHistory = async () => {
    if (userId <= 0) return;

    try {
      setLoading(true);
      setError(null);
      const historyData = await gachaUsecase.getGachaHistory(userId, 10);
      setHistory(historyData || []);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch history');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchHistory();
  }, [userId, refreshTrigger]);

  const getRarityColor = (rarity: Rarity): string => {
    switch (rarity) {
      case Rarity.Common:
        return '#808080';
      case Rarity.Rare:
        return '#0066ff';
      case Rarity.Epic:
        return '#8b00ff';
      case Rarity.Legendary:
        return '#ff6600';
      default:
        return '#000000';
    }
  };

  const formatDate = (dateString: string): string => {
    return new Date(dateString).toLocaleString();
  };

  if (loading) {
    return <div className="gacha-history loading">Loading history...</div>;
  }

  if (error) {
    return (
      <div className="gacha-history error">
        <p>Error: {error}</p>
        <button onClick={fetchHistory}>Retry</button>
      </div>
    );
  }

  return (
    <div className="gacha-history">
      <h3>Recent Gacha Results</h3>
      {!history || history.length === 0 ? (
        <p className="no-history">No gacha history yet. Try spinning!</p>
      ) : (
        <div className="history-list">
          {history.map((item) => (
            <div key={item.id} className="history-item">
              <div className="item-info">
                <span
                  className="item-name"
                  style={{ color: getRarityColor(item.rarity) }}
                >
                  {item.itemName}
                </span>
                <span className="rarity">({item.rarity})</span>
              </div>
              <div className="points">+{item.pointsEarned}</div>
              <div className="date">{formatDate(item.createdAt)}</div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};