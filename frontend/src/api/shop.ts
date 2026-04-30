import { get, post } from './client'

export const shopApi = {
  getShopDetail: (id: number) => get(`/shops/${id}`),

  getShopProducts: (id: number, params?: { page?: number; pageSize?: number }) =>
    get(`/shops/${id}/products`, params),

  toggleFollow: (id: number, follow: boolean) =>
    post(`/shops/${id}/follow`, { follow }),

  checkFollow: (id: number) => get(`/shops/${id}/follow/check`),
}
