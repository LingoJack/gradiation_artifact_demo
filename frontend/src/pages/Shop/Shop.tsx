import React, { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';
import { MapPin, Star, Users, ShoppingBag, ChevronRight, Store } from 'lucide-react';
import { shopApi } from '../../api';
import { ProductCard } from '../../components/ProductCard/ProductCard';
import type { Product } from '../../types/product';

interface ShopData {
  id: number;
  name: string;
  description: string;
  avatar: string;
  cover_image: string;
  rating: number;
  sales: number;
  fans: number;
  location: string;
}

interface RawProduct {
  id: number;
  name: string;
  price: number;
  main_image: string;
  stock: number;
  sales: number;
  category_id: number;
  status: number;
  description: string;
  images: string;
}

function mapProduct(raw: RawProduct): Product {
  let images: string[] = [];
  try {
    if (raw.images) {
      const parsed = JSON.parse(raw.images);
      if (Array.isArray(parsed)) images = parsed;
    }
  } catch { /* ignore */ }

  return {
    id: String(raw.id),
    categoryId: String(raw.category_id),
    name: raw.name,
    description: raw.description,
    price: raw.price,
    stock: raw.stock,
    sales: raw.sales,
    mainImage: raw.main_image,
    image: raw.main_image,
    images,
    specs: [],
    status: raw.status === 1 ? 'active' : 'inactive',
    createdAt: '',
  };
}

export const Shop: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [shop, setShop] = useState<ShopData | null>(null);
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchShop = async () => {
      try {
        setLoading(true);
        const shopId = Number(id);
        const [shopRes, productsRes] = await Promise.all([
          shopApi.getShopDetail(shopId),
          shopApi.getShopProducts(shopId, { page: 1, pageSize: 20 }),
        ]);

        setShop(shopRes as ShopData);
        const prods = (productsRes as any).products || [];
        setProducts(prods.map(mapProduct));
      } catch (err) {
        console.error('Failed to fetch shop:', err);
      } finally {
        setLoading(false);
      }
    };

    if (id) fetchShop();
  }, [id]);

  const formatNumber = (num: number) => {
    if (num >= 10000) return `${(num / 10000).toFixed(1)}万`;
    return num.toString();
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center py-40">
        <div className="flex flex-col items-center gap-4">
          <div className="w-12 h-12 border-4 border-blue-200 border-t-blue-600 rounded-full animate-spin"></div>
          <p className="text-gray-500 dark:text-gray-400 text-sm">加载中...</p>
        </div>
      </div>
    );
  }

  if (!shop) {
    return (
      <div className="flex items-center justify-center py-40">
        <div className="text-center">
          <Store className="w-16 h-16 text-gray-300 mx-auto mb-4" />
          <p className="text-gray-500">店铺不存在</p>
          <Link to="/" className="text-primary hover:underline text-sm mt-2 inline-block">返回首页</Link>
        </div>
      </div>
    );
  }

  return (
    <div className="container py-6">
      {/* 店铺头部 */}
      <div className="glass-card rounded-xl overflow-hidden mb-6">
        {/* 封面图 */}
        <div 
          className="h-48 bg-cover bg-center relative"
          style={{ backgroundImage: `url(${shop.cover_image || 'https://images.unsplash.com/photo-1441986300917-64674bd600d8?w=800&h=300&fit=crop'})` }}
        >
          <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent" />
        </div>
        
        {/* 店铺信息 */}
        <div className="relative px-6 pb-6">
          {/* 头像 */}
          <div className="absolute -top-12 left-6">
            <img 
              src={shop.avatar} 
              alt={shop.name}
              className="w-24 h-24 rounded-full border-4 border-white dark:border-gray-800 object-cover"
            />
          </div>
          
          <div className="pt-14">
            <h1 className="text-2xl font-bold dark:text-white mb-2">{shop.name}</h1>
            <p className="text-gray-500 dark:text-gray-400 text-sm mb-4">{shop.description}</p>
            
            {/* 数据统计 */}
            <div className="flex items-center gap-6 text-sm">
              <div className="flex items-center gap-1">
                <Star className="w-4 h-4 text-yellow-500 fill-current" />
                <span className="font-medium dark:text-white">{shop.rating}</span>
                <span className="text-gray-400">评分</span>
              </div>
              <div className="flex items-center gap-1">
                <ShoppingBag className="w-4 h-4 text-primary" />
                <span className="font-medium dark:text-white">{formatNumber(shop.sales)}</span>
                <span className="text-gray-400">销量</span>
              </div>
              <div className="flex items-center gap-1">
                <Users className="w-4 h-4 text-blue-500" />
                <span className="font-medium dark:text-white">{formatNumber(shop.fans)}</span>
                <span className="text-gray-400">粉丝</span>
              </div>
              <div className="flex items-center gap-1">
                <MapPin className="w-4 h-4 text-gray-400" />
                <span className="text-gray-500 dark:text-gray-400">{shop.location || '未知'}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* 导航栏 */}
      <div className="glass-card rounded-xl p-4 mb-6">
        <div className="flex items-center gap-6">
          <Link 
            to="/"
            className="text-gray-500 dark:text-gray-400 hover:text-primary transition-colors text-sm"
          >
            首页
          </Link>
          <ChevronRight className="w-4 h-4 text-gray-300" />
          <span className="text-sm font-medium dark:text-white">{shop.name}</span>
        </div>
      </div>

      {/* 商品列表 */}
      <div className="glass-card rounded-xl p-6">
        <div className="flex items-center justify-between mb-6">
          <h2 className="text-lg font-bold dark:text-white flex items-center gap-2">
            <Store className="w-5 h-5" />
            全部商品
            <span className="text-sm font-normal text-gray-400">({products.length})</span>
          </h2>
        </div>

        {products.length > 0 ? (
          <div className="grid grid-cols-4 gap-4">
            {products.map((product) => (
              <ProductCard key={product.id} product={product as any} />
            ))}
          </div>
        ) : (
          <div className="text-center py-12">
            <Store className="w-16 h-16 text-gray-300 dark:text-gray-600 mx-auto mb-4" />
            <p className="text-gray-500 dark:text-gray-400">暂无商品</p>
          </div>
        )}
      </div>
    </div>
  );
};