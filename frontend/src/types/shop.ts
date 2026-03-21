export interface Shop {
  id: string;
  name: string;
  description: string;
  avatar: string;
  coverImage: string;
  rating: number;
  sales: number;
  fans: number;
  createdAt: string;
  location: string;
}

export interface ShopProduct {
  id: string;
  shopId: string;
  name: string;
  price: number;
  originalPrice: number;
  mainImage: string;
  sales: number;
  stock: number;
}
