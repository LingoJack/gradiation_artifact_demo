import { get, post, put, del } from './client'

export const cartApi = {
  getCart: () => get('/cart'),

  addItem: (data: { productId: number; skuId?: number; quantity: number }) =>
    post('/cart', data),

  updateQuantity: (id: number, data: { quantity: number }) =>
    put(`/cart/${id}`, data),

  removeItem: (id: number) => del(`/cart/${id}`),

  updateSelected: (data: { itemIds: number[]; selected: boolean }) =>
    put('/cart/selected', data),

  selectAll: (data: { selected: boolean }) =>
    put('/cart/select-all', data),

  clearCart: () => del('/cart/clear'),
}
