import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { Heart, Trash2, ShoppingCart } from 'lucide-react';
import { useSpotlight } from '../../hooks/useSpotlight';
import { useCartStore } from '../../store/useCartStore';

interface FavoriteItem {
  id: string;
  productId: string;
  name: string;
  image: string;
  price: number;
  originalPrice?: number;
  stock: number;
}

// Mock 数据
const mockFavorites: FavoriteItem[] = [
  {
    id: '1',
    productId: '1',
    name: '时尚休闲连帽卫衣 男士秋季新款纯棉舒适',
    image: 'https://images.unsplash.com/photo-1556821840-3a63f95609a7?w=200&h=200&fit=crop',
    price: 129,
    originalPrice: 199,
    stock: 100,
  },
  {
    id: '2',
    productId: '2',
    name: 'Apple iPhone 15 Pro Max 256GB 原色钛金属',
    image: 'https://images.unsplash.com/photo-1695048133142-1a20484d2569?w=200&h=200&fit=crop',
    price: 9999,
    stock: 50,
  },
  {
    id: '3',
    productId: '3',
    name: '北欧简约落地灯 客厅卧室书房装饰台灯',
    image: 'https://images.unsplash.com/photo-1507473885765-e6ed057f782c?w=200&h=200&fit=crop',
    price: 299,
    originalPrice: 399,
    stock: 30,
  },
];

export const Favorites: React.FC = () => {
  const [favorites, setFavorites] = useState<FavoriteItem[]>(mockFavorites);
  const cardSpotlight = useSpotlight();
  const addItem = useCartStore((state) => state.addItem);

  const handleRemove = (id: string) => {
    setFavorites(favorites.filter((item) => item.id !== id));
  };

  const handleAddToCart = (item: FavoriteItem) => {
    // 模拟后端返回的购物车项 ID
    const cartItemId = `cart-${item.productId}-${Date.now().toString(36)}`;
    addItem({
      id: cartItemId,
      userId: 'user-1',
      productId: item.productId,
      product: {
        id: item.productId,
        name: item.name,
        mainImage: item.image,
        price: item.price,
        stock: item.stock,
      },
      quantity: 1,
      selected: true,
    });
  };

  return (
    <div className="container py-8">
      <h1 className="text-2xl font-bold mb-6 dark:text-white">我的收藏</h1>

      {favorites.length === 0 ? (
        <div className="glass-card rounded-xl p-12 text-center">
          <Heart className="w-16 h-16 text-gray-300 dark:text-gray-600 mx-auto mb-4" />
          <p className="text-gray-500 dark:text-gray-400">暂无收藏商品</p>
          <Link to="/products" className="mt-4 inline-block text-primary hover:underline">
            去逛逛
          </Link>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
          {favorites.map((item) => (
            <div
              key={item.id}
              ref={cardSpotlight.ref as React.RefObject<HTMLDivElement>}
              className="glass-liquid rounded-xl overflow-hidden relative group"
              style={cardSpotlight.spotlightStyle}
              {...cardSpotlight.handlers}
            >
              <Link to={`/products/${item.productId}`} className="block">
                <img
                  src={item.image}
                  alt={item.name}
                  className="w-full aspect-square object-cover bg-gray-100 dark:bg-gray-700"
                  onError={(e) => {
                    const target = e.target as HTMLImageElement;
                    target.src = 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="200" height="200"%3E%3Crect fill="%23f3f4f6" width="200" height="200"/%3E%3Ctext fill="%239ca3af" x="50%25" y="50%25" text-anchor="middle" dy=".3em" font-size="12"%3E暂无图片%3C/text%3E%3C/svg%3E';
                  }}
                />
              </Link>
              <div className="p-4">
                <Link
                  to={`/products/${item.productId}`}
                  className="text-sm line-clamp-2 hover:text-primary dark:text-gray-200 dark:hover:text-primary"
                >
                  {item.name}
                </Link>
                <div className="flex items-baseline space-x-2 mt-2">
                  <span className="text-lg font-bold text-primary">¥{item.price}</span>
                  {item.originalPrice && (
                    <span className="text-sm text-gray-400 line-through">
                      ¥{item.originalPrice}
                    </span>
                  )}
                </div>
                <div className="flex items-center space-x-2 mt-3">
                  <button
                    onClick={() => handleAddToCart(item)}
                    className="flex-1 flex items-center justify-center space-x-1 py-2 bg-primary text-white rounded-lg hover:bg-primary-hover text-sm"
                  >
                    <ShoppingCart className="w-4 h-4" />
                    <span>加入购物车</span>
                  </button>
                  <button
                    onClick={() => handleRemove(item.id)}
                    className="p-2 text-gray-400 hover:text-error rounded-lg hover:bg-gray-100 dark:hover:bg-gray-700"
                  >
                    <Trash2 className="w-5 h-5" />
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};
