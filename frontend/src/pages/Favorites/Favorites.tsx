import React, { useState, useEffect, useCallback } from 'react';
import { Link } from 'react-router-dom';
import { Heart, Trash2, ShoppingCart, Loader2 } from 'lucide-react';
import { useSpotlight } from '../../hooks/useSpotlight';
import { useCartStore } from '../../store/useCartStore';
import { favoriteApi } from '../../api/favorite';
import { cartApi } from '../../api/cart';
import { showToast } from '../../utils/toast';

interface FavoriteItem {
  id: string;
  productId: string;
  name: string;
  image: string;
  price: number;
  sales: number;
}

interface BackendFavoriteProduct {
  id: number;
  name: string;
  price: number;
  main_image: string;
  sales: number;
}

interface BackendFavorite {
  id: number;
  user_id: number;
  product_id: number;
  created_at: string;
  product: BackendFavoriteProduct;
}

// 映射后端数据到前端格式
const mapFavoriteFromBackend = (item: BackendFavorite): FavoriteItem => ({
  id: String(item.id),
  productId: String(item.product_id),
  name: item.product.name,
  image: item.product.main_image,
  price: item.product.price,
  sales: item.product.sales,
});

export const Favorites: React.FC = () => {
  const [favorites, setFavorites] = useState<FavoriteItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [removingId, setRemovingId] = useState<string | null>(null);
  const cardSpotlight = useSpotlight();
  const addItem = useCartStore((state) => state.addItem);

  // 获取收藏列表
  const fetchFavorites = useCallback(async () => {
    try {
      setLoading(true);
      const response = await favoriteApi.getFavorites({ page: 1, pageSize: 20 });
      if (response.code === 0 && response.data?.favorites) {
        setFavorites(response.data.favorites.map(mapFavoriteFromBackend));
      }
    } catch (error) {
      console.error('获取收藏列表失败:', error);
      showToast('获取收藏列表失败', 'error');
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    fetchFavorites();
  }, [fetchFavorites]);

  // 取消收藏
  const handleRemove = async (id: string, productId: string) => {
    try {
      setRemovingId(id);
      const response = await favoriteApi.removeFavorite(Number(productId));
      if (response.code === 0) {
        showToast('已取消收藏', 'success');
        await fetchFavorites();
      } else {
        showToast(response.message || '取消收藏失败', 'error');
      }
    } catch (error) {
      console.error('取消收藏失败:', error);
      showToast('取消收藏失败', 'error');
    } finally {
      setRemovingId(null);
    }
  };

  // 加入购物车
  const handleAddToCart = async (item: FavoriteItem) => {
    try {
      const response = await cartApi.addItem({
        productId: Number(item.productId),
        quantity: 1,
      });
      if (response.code === 0) {
        // 同步更新本地 cart store
        addItem({
          id: `cart-${item.productId}-${Date.now().toString(36)}`,
          userId: 'user-1',
          productId: item.productId,
          product: {
            id: item.productId,
            name: item.name,
            mainImage: item.image,
            price: item.price,
            stock: 999,
          },
          quantity: 1,
          selected: true,
        });
        showToast('已加入购物车', 'success');
      } else {
        showToast(response.message || '加入购物车失败', 'error');
      }
    } catch (error) {
      console.error('加入购物车失败:', error);
      showToast('加入购物车失败', 'error');
    }
  };

  return (
    <div className="container py-8">
      <h1 className="text-2xl font-bold mb-6 dark:text-white">我的收藏</h1>

      {loading ? (
        <div className="glass-card rounded-xl p-12 text-center">
          <Loader2 className="w-8 h-8 text-primary animate-spin mx-auto mb-4" />
          <p className="text-gray-500 dark:text-gray-400">加载中...</p>
        </div>
      ) : favorites.length === 0 ? (
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
                  {item.sales > 0 && (
                    <span className="text-sm text-gray-400">
                      已售 {item.sales}
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
                    onClick={() => handleRemove(item.id, item.productId)}
                    disabled={removingId === item.id}
                    className="p-2 text-gray-400 hover:text-error rounded-lg hover:bg-gray-100 dark:hover:bg-gray-700 disabled:opacity-50 disabled:cursor-not-allowed"
                  >
                    {removingId === item.id ? (
                      <Loader2 className="w-5 h-5 animate-spin" />
                    ) : (
                      <Trash2 className="w-5 h-5" />
                    )}
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
