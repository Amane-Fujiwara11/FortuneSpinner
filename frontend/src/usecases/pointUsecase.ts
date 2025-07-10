import { pointApi } from '../infrastructure/api/pointApi';
import { UserPoint, PointTransaction } from '../domain/Point';

export class PointUsecase {
  async getBalance(userId: number): Promise<UserPoint> {
    if (userId <= 0) {
      throw new Error('Invalid user ID');
    }
    return pointApi.getBalance(userId);
  }

  async getTransactionHistory(userId: number, limit: number = 20): Promise<PointTransaction[]> {
    if (userId <= 0) {
      throw new Error('Invalid user ID');
    }
    return pointApi.getTransactionHistory(userId, limit);
  }
}

export const pointUsecase = new PointUsecase();