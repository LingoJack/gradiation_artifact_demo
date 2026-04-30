import { get, post, put, del } from './client'

export const orderApi = {
  createOrder: (data: {
    addressId: number
    couponId?: number
    remark?: string
    items: Array<{ productId: number; skuId?: number; quantity: number }>
  }) => post('/orders', data),

  getOrders: (params?: { status?: string; page?: number; pageSize?: number }) =>
    get('/orders', params),

  getOrderDetail: (id: number) => get(`/orders/${id}`),

  cancelOrder: (id: number) => put(`/orders/${id}/cancel`),

  payOrder: (id: number) => put(`/orders/${id}/pay`),

  confirmReceive: (id: number) => put(`/orders/${id}/confirm`),

  deleteOrder: (id: number) => del(`/orders/${id}`),
}
