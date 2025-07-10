import React, { useState } from 'react';
import { gachaUsecase } from '../../usecases/gachaUsecase';
import { GachaResult, Rarity } from '../../domain/Gacha';

interface GachaSpinnerProps {
  userId: number;
  onGachaComplete?: (result: GachaResult) => void;
}

export const GachaSpinner: React.FC<GachaSpinnerProps> = ({ userId, onGachaComplete }) => {
  const [isSpinning, setIsSpinning] = useState<boolean>(false);
  const [result, setResult] = useState<GachaResult | null>(null);
  const [error, setError] = useState<string | null>(null);

  const executeGacha = async () => {
    if (isSpinning || userId <= 0) return;

    try {
      setIsSpinning(true);
      setError(null);
      setResult(null);

      // アニメーション効果のため少し待機
      await new Promise(resolve => setTimeout(resolve, 1500));

      const gachaResult = await gachaUsecase.executeGacha(userId);
      setResult(gachaResult);
      onGachaComplete?.(gachaResult);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to execute gacha');
    } finally {
      setIsSpinning(false);
    }
  };

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

  return (
    <div className="gacha-spinner">
      <div className="spinner-container">
        <div className={`spinner ${isSpinning ? 'spinning' : ''}`}>
          {isSpinning ? (
            <div className="spinning-content">
              <div className="loading-circle"></div>
              <p>Spinning...</p>
            </div>
          ) : result ? (
            <div className="result-content">
              <h3 style={{ color: getRarityColor(result.rarity) }}>
                {result.itemName}
              </h3>
              <p className="rarity">{result.rarity}</p>
              <p className="points">+{result.pointsEarned} Points</p>
            </div>
          ) : (
            <div className="idle-content">
              <p>Ready to spin!</p>
            </div>
          )}
        </div>
      </div>

      <button
        className="spin-button"
        onClick={executeGacha}
        disabled={isSpinning || userId <= 0}
      >
        {isSpinning ? 'Spinning...' : 'Spin Gacha'}
      </button>

      {error && (
        <div className="error-message">
          <p>Error: {error}</p>
        </div>
      )}
    </div>
  );
};