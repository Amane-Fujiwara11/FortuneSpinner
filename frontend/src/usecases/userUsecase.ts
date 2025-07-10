import { userApi } from '../infrastructure/api/userApi';
import { User } from '../domain/User';

export class UserUsecase {
  async createUser(name: string): Promise<User> {
    if (!name || name.trim().length === 0) {
      throw new Error('Name is required');
    }
    return userApi.createUser(name.trim());
  }
}

export const userUsecase = new UserUsecase();