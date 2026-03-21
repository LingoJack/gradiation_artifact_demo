import React, { useState, useCallback } from 'react';
import { Link } from 'react-router-dom';
import type { Product } from '../../types/product';

interface ProductCardProps {
  product: Product;
}

export const ProductCard: React.FC<ProductCardProps> = ({ product }) => {
  const [mousePos, setMousePos] = useState({ x: 0, y: 0 });
  const [isHovered, setIsHovered] = useState(false);

  const handleMouseMove = useCallback((e: React.MouseEvent<HTMLDivElement>) => {
    const rect = e.currentTarget.getBoundingClientRect();
    setMousePos({
      x: e.clientX - rect.left,
      y: e.clientY - rect.top,
    });
  }, []);

  return (
    <Link to={`/products/${product.id}`} className="group block">
      <div 
        className="glass-card rounded-xl overflow-hidden hover:shadow-2xl hover:scale-[1.02] transition-all duration-300 relative"
        onMouseMove={handleMouseMove}
        onMouseEnter={() => setIsHovered(true)}
        onMouseLeave={() => setIsHovered(false)}
      >
        {/* 鼠标跟随光效 */}
        <div 
          className="pointer-events-none absolute inset-0 z-10 transition-opacity duration-300"
          style={{
            opacity: isHovered ? 1 : 0,
            background: `radial-gradient(200px circle at ${mousePos.x}px ${mousePos.y}px, rgba(251, 146, 60, 0.08), transparent 50%)`,
          }}
        />
        
        {/* 商品图片 */}
        <div className="aspect-square overflow-hidden relative bg-gray-50 dark:bg-gray-900">
          <img 
            src={product.mainImage || product.image} 
            alt={product.name}
            className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
            onError={(e) => {
              const target = e.target as HTMLImageElement;
              target.style.display = 'none';
              const parent = target.parentElement;
              if (parent) {
                parent.style.background = 'linear-gradient(135deg, #f3f4f6 0%, #e5e7eb 100%)';
                parent.innerHTML = `<div class="w-full h-full flex items-center justify-center text-gray-300"><svg class="w-16 h-16" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path></svg></div>`;
              }
            }}
          />
          {/* 促销标签 */}
          {product.originalPrice && (
            <div className="absolute top-2 left-2 bg-red-500 text-white text-xs px-2 py-1 rounded">
              限时折扣
            </div>
          )}
        </div>
        
        {/* 商品信息 */}
        <div className="p-4">
          {/* 标题 */}
          <h3 className="text-sm text-gray-800 dark:text-gray-200 line-clamp-2 h-10 mb-2 leading-5 group-hover:text-orange-600 transition-colors">
            {product.name}
          </h3>
          
          {/* 价格和销量 */}
          <div className="flex items-end justify-between">
            <div className="flex items-baseline">
              <span className="text-xs text-gray-400 mr-1">¥</span>
              <span className="text-xl font-bold text-orange-600">{product.price}</span>
            </div>
            <span className="text-xs text-gray-400 dark:text-gray-500">{product.sales.toLocaleString()}人付款</span>
          </div>
          
          {/* 原价 */}
          {product.originalPrice && (
            <div className="mt-1">
              <span className="text-xs text-gray-400 line-through">¥{product.originalPrice}</span>
            </div>
          )}
          
          {/* 店铺信息 */}
          <div className="mt-2 pt-2 border-t border-gray-50 dark:border-gray-700 flex items-center text-xs text-gray-400 dark:text-gray-500">
            <span className="truncate">{product.shopName || '官方旗舰店'}</span>
          </div>
        </div>
      </div>
    </Link>
  );
};
