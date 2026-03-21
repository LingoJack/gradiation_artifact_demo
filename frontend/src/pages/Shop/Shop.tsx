import React, { useMemo } from 'react';
import { useParams, Link } from 'react-router-dom';
import { MapPin, Star, Users, ShoppingBag, ChevronRight, Store } from 'lucide-react';
import { mockProducts, getShopByName, getShopIdByName } from '../../utils/mockData';
import { ProductCard } from '../../components/ProductCard/ProductCard';

export const Shop: React.FC = () => {
  const { id } = useParams<{ id: string }>();

  // 根据 URL 中的 id 找到对应的店铺名称
  const shop = useMemo(() => {
    // 遍历所有商品，找到 shopId 匹配的店铺
    for (const product of mockProducts) {
      if (getShopIdByName(product.shopName || '') === id) {
        return getShopByName(product.shopName || '');
      }
    }
    // 默认返回第一个商品的店铺
    return getShopByName(mockProducts[0]?.shopName || '');
  }, [id]);

  const shopProducts = useMemo(() => {
    // 只显示店铺名称完全匹配的商品
    return mockProducts.filter((p) => p.shopName === shop.name);
  }, [shop.name]);

  const formatNumber = (num: number) => {
    if (num >= 10000) {
      return `${(num / 10000).toFixed(1)}万`;
    }
    return num.toString();
  };

  return (
    <div className="container py-6">
      {/* 店铺头部 */}
      <div className="glass-card rounded-xl overflow-hidden mb-6">
        {/* 封面图 */}
        <div 
          className="h-48 bg-cover bg-center relative"
          style={{ backgroundImage: `url(${shop.coverImage})` }}
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
                <span className="text-gray-500 dark:text-gray-400">{shop.location}</span>
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
            <span className="text-sm font-normal text-gray-400">({shopProducts.length})</span>
          </h2>
        </div>

        {shopProducts.length > 0 ? (
          <div className="grid grid-cols-4 gap-4">
            {shopProducts.map((product) => (
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
