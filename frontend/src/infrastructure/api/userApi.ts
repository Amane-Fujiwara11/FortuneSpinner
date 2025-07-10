import { apiClient } from './client';
import { User } from '../../domain/User';

export interface CreateUserRequest {
  name: string;
}

export const userApi = {
  createUser: (name: string): Promise<User> =>
    apiClient.post<User>('/users', { name }),
};