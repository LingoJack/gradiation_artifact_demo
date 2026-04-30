import { get } from './client'

export const productApi = {
  getProducts: (params?: { categoryId?: number; keyword?: string; page?: number; pageSize?: number; sort?: string }) =>
    get('/products', params),

  getProductDetail: (id: number) => get(`/products/${id}`),

  searchSuggestions: (keyword: string) => get('/products/search', { keyword }),

  getCategories: () => get('/categories'),

  getBanners: () => get('/banners'),
}
