import { apiClient } from './client';
import { UserPoint, PointTransaction } from '../../domain/Point';

export const pointApi = {
  getBalance: (userId: number): Promise<UserPoint> =>
    apiClient.get<UserPoint>(`/points/balance?user_id=${userId}`),

  getTransactionHistory: (userId: number, limit: number = 20): Promise<PointTransaction[]> =>
    apiClient.get<PointTransaction[]>(`/points/transactions?user_id=${userId}&limit=${limit}`),
};