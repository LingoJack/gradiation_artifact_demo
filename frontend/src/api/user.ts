import { get, post, put, del } from './client'

// 地址数据接口
export interface AddressData {
  receiver_name?: string
  receiver_phone?: string
  province?: string
  city?: string
  district?: string
  detail_address?: string
  is_default?: boolean
}

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

  createAddress: (data: AddressData) => post('/user/addresses', data),

  updateAddress: (id: number, data: AddressData) => put(`/user/addresses/${id}`, data),

  deleteAddress: (id: number) => del(`/user/addresses/${id}`),

  setDefaultAddress: (id: number) => put(`/user/addresses/${id}/default`),
}
