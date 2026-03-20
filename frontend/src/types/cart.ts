export interface CartItem {
  id: string;
  userId: string;
  productId: string;
  product: Product;
  specId?: string;
  spec?: ProductSpec;
  quantity: number;
  selected: boolean;
}

export interface Product {
  id: string;
  name: string;
  price: number;
  mainImage: string;
  stock: number;
}

export interface ProductSpec {
  id: string;
  name: string;
  value: string;
  price: number;
  stock: number;
}
