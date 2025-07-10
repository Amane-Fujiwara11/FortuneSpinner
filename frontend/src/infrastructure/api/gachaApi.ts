import { apiClient } from './client';
import { GachaResult, GachaHistory } from '../../domain/Gacha';

export interface ExecuteGachaRequest {
  user_id: number;
}

export const gachaApi = {
  executeGacha: (userId: number): Promise<GachaResult> =>
    apiClient.post<GachaResult>('/gacha/execute', { user_id: userId }),

  getGachaHistory: (userId: number, limit: number = 20): Promise<GachaHistory[]> =>
    apiClient.get<GachaHistory[]>(`/gacha/history?user_id=${userId}&limit=${limit}`),
};