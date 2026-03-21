import React, { useState } from 'react';
import { ProductCard } from '../../components/ProductCard/ProductCard';
import { mockProducts, mockCategories } from '../../utils/mockData';
import type { Product } from '../../types/product';

export const ProductList: React.FC = () => {
  const [selectedCategory, setSelectedCategory] = useState<string>('');
  const [sortBy, setSortBy] = useState<'default' | 'price-asc' | 'price-desc' | 'sales'>('default');

  const filteredProducts = selectedCategory
    ? mockProducts.filter((p) => p.categoryId === selectedCategory)
    : mockProducts;

  const sortedProducts = [...filteredProducts].sort((a, b) => {
    switch (sortBy) {
      case 'price-asc':
        return a.price - b.price;
      case 'price-desc':
        return b.price - a.price;
      case 'sales':
        return b.sales - a.sales;
      default:
        return 0;
    }
  });

  return (
    <div className="container py-6">
      {/* 分类筛选 */}
      <div className="glass-card rounded-xl p-4 mb-6">
        <div className="flex items-center space-x-4">
          <span className="text-sm font-medium text-gray-700 dark:text-gray-300">分类：</span>
          <button
            onClick={() => setSelectedCategory('')}
            className={`px-4 py-1 rounded-lg transition-all ${
              !selectedCategory
                ? 'bg-primary text-white shadow-md'
                : 'bg-white/50 dark:bg-gray-700/50 hover:bg-white/80 dark:hover:bg-gray-700/80'
            }`}
          >
            全部
          </button>
          {mockCategories.map((cat) => (
            <button
              key={cat.id}
              onClick={() => setSelectedCategory(cat.id)}
              className={`px-4 py-1 rounded-lg transition-all ${
                selectedCategory === cat.id
                  ? 'bg-primary text-white shadow-md'
                  : 'bg-white/50 dark:bg-gray-700/50 hover:bg-white/80 dark:hover:bg-gray-700/80'
              }`}
            >
              {cat.icon} {cat.name}
            </button>
          ))}
        </div>
      </div>

      {/* 排序 */}
      <div className="glass-card rounded-xl p-4 mb-6">
        <div className="flex items-center space-x-4">
          <span className="text-sm font-medium text-gray-700 dark:text-gray-300">排序：</span>
          <button
            onClick={() => setSortBy('default')}
            className={`px-4 py-1 rounded-lg transition-all ${
              sortBy === 'default'
                ? 'bg-primary text-white shadow-md'
                : 'bg-white/50 dark:bg-gray-700/50 hover:bg-white/80 dark:hover:bg-gray-700/80'
            }`}
          >
            综合
          </button>
          <button
            onClick={() => setSortBy('sales')}
            className={`px-4 py-1 rounded-lg transition-all ${
              sortBy === 'sales'
                ? 'bg-primary text-white shadow-md'
                : 'bg-white/50 dark:bg-gray-700/50 hover:bg-white/80 dark:hover:bg-gray-700/80'
            }`}
          >
            销量
          </button>
          <button
            onClick={() => setSortBy('price-asc')}
            className={`px-4 py-1 rounded-lg transition-all ${
              sortBy === 'price-asc'
                ? 'bg-primary text-white shadow-md'
                : 'bg-white/50 dark:bg-gray-700/50 hover:bg-white/80 dark:hover:bg-gray-700/80'
            }`}
          >
            价格升序
          </button>
          <button
            onClick={() => setSortBy('price-desc')}
            className={`px-4 py-1 rounded-lg transition-all ${
              sortBy === 'price-desc'
                ? 'bg-primary text-white shadow-md'
                : 'bg-white/50 dark:bg-gray-700/50 hover:bg-white/80 dark:hover:bg-gray-700/80'
            }`}
          >
            价格降序
          </button>
        </div>
      </div>

      {/* 商品列表 */}
      <div className="grid grid-cols-4 gap-6">
        {sortedProducts.map((product) => (
          <ProductCard key={product.id} product={product as Product} />
        ))}
      </div>
    </div>
  );
};
