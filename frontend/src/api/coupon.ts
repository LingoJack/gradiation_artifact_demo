import { get, post, put } from './client'

export const couponApi = {
  getAvailable: () => get('/coupons/available'),

  getUserCoupons: () => get('/coupons/mine'),

  claimCoupon: (id: number) => post(`/coupons/${id}/claim`),

  useCoupon: (id: number) => put(`/coupons/${id}/use`),
}
