export interface Product {
  id: string;
  categoryId: string;
  name: string;
  description: string;
  price: number;
  originalPrice?: number;
  stock: number;
  sales: number;
  mainImage: string;
  image?: string;
  images: string[];
  specs: ProductSpec[];
  status: 'active' | 'inactive';
  createdAt: string;
  shopName?: string;
}

export interface ProductSpec {
  id: string;
  productId: string;
  name: string;
  value: string;
  stock: number;
  price: number;
}

export interface Category {
  id: string;
  name: string;
  parentId?: string;
  icon?: string;
  sortOrder: number;
}

export interface Review {
  id: string;
  userId: string;
  userName: string;
  productId: string;
  orderId: string;
  rating: number;
  content: string;
  images: string[];
  createdAt: string;
}
