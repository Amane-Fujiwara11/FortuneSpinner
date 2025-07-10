import { gachaApi } from '../infrastructure/api/gachaApi';
import { GachaResult, GachaHistory } from '../domain/Gacha';

export class GachaUsecase {
  async executeGacha(userId: number): Promise<GachaResult> {
    if (userId <= 0) {
      throw new Error('Invalid user ID');
    }
    return gachaApi.executeGacha(userId);
  }

  async getGachaHistory(userId: number, limit: number = 20): Promise<GachaHistory[]> {
    if (userId <= 0) {
      throw new Error('Invalid user ID');
    }
    return gachaApi.getGachaHistory(userId, limit);
  }
}

export const gachaUsecase = new GachaUsecase();