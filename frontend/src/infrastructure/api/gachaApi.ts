import { apiClient } from './client';
import { GachaResult, GachaHistory } from '../../domain/Gacha';

export interface ExecuteGachaRequest {
  user_id: number;
}

interface GachaHistoryResponse {
  id: number;
  item_name: string;
  rarity: string;
  points_earned: number;
  created_at: string;
}

export const gachaApi = {
  executeGacha: (userId: number): Promise<GachaResult> =>
    apiClient.post<GachaResult>('/gacha/execute', { user_id: userId }),

  getGachaHistory: async (userId: number, limit: number = 20): Promise<GachaHistory[]> => {
    const response = await apiClient.get<GachaHistoryResponse[]>(`/gacha/history?user_id=${userId}&limit=${limit}`);
    return response.map(item => ({
      id: item.id,
      itemName: item.item_name,
      rarity: item.rarity as any,
      pointsEarned: item.points_earned,
      createdAt: item.created_at,
    }));
  },
};