import { get, post, del } from './client'

export const favoriteApi = {
  getFavorites: (params?: { page?: number; pageSize?: number }) =>
    get('/favorites', params),

  addFavorite: (productId: number) =>
    post('/favorites', { productId }),

  removeFavorite: (productId: number) =>
    del(`/favorites/${productId}`),

  checkFavorite: (productId: number) =>
    get(`/favorites/${productId}/check`),
}
