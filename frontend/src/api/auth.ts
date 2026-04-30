import { post } from './client'

export const authApi = {
  register: (data: { username: string; password: string; phone: string; email?: string }) =>
    post('/auth/register', data),

  login: (data: { username: string; password: string }) =>
    post('/auth/login', data),

  logout: () => post('/auth/logout'),
}
