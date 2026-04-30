import { get, post, put, del } from './client'

export const userApi = {
  getProfile: () => get('/user/profile'),

  updateProfile: (data: {
    nickname?: string
    avatar?: string
    gender?: number
    birthday?: string
    bio?: string
  }) => put('/user/profile', data),

  getAddresses: () => get('/user/addresses'),

  createAddress: (data: any) => post('/user/addresses', data),

  updateAddress: (id: number, data: any) => put(`/user/addresses/${id}`, data),

  deleteAddress: (id: number) => del(`/user/addresses/${id}`),

  setDefaultAddress: (id: number) => put(`/user/addresses/${id}/default`),
}
