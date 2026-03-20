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
      <div className="bg-white rounded-lg p-4 mb-6">
        <div className="flex items-center space-x-4">
          <span className="text-sm font-medium text-gray-700">分类：</span>
          <button
            onClick={() => setSelectedCategory('')}
            className={`px-4 py-1 rounded ${
              !selectedCategory
                ? 'bg-primary text-white'
                : 'bg-gray-100 hover:bg-gray-200'
            }`}
          >
            全部
          </button>
          {mockCategories.map((cat) => (
            <button
              key={cat.id}
              onClick={() => setSelectedCategory(cat.id)}
              className={`px-4 py-1 rounded ${
                selectedCategory === cat.id
                  ? 'bg-primary text-white'
                  : 'bg-gray-100 hover:bg-gray-200'
              }`}
            >
              {cat.icon} {cat.name}
            </button>
          ))}
        </div>
      </div>

      {/* 排序 */}
      <div className="bg-white rounded-lg p-4 mb-6">
        <div className="flex items-center space-x-4">
          <span className="text-sm font-medium text-gray-700">排序：</span>
          <button
            onClick={() => setSortBy('default')}
            className={`px-4 py-1 rounded ${
              sortBy === 'default'
                ? 'bg-primary text-white'
                : 'bg-gray-100 hover:bg-gray-200'
            }`}
          >
            综合
          </button>
          <button
            onClick={() => setSortBy('sales')}
            className={`px-4 py-1 rounded ${
              sortBy === 'sales'
                ? 'bg-primary text-white'
                : 'bg-gray-100 hover:bg-gray-200'
            }`}
          >
            销量
          </button>
          <button
            onClick={() => setSortBy('price-asc')}
            className={`px-4 py-1 rounded ${
              sortBy === 'price-asc'
                ? 'bg-primary text-white'
                : 'bg-gray-100 hover:bg-gray-200'
            }`}
          >
            价格升序
          </button>
          <button
            onClick={() => setSortBy('price-desc')}
            className={`px-4 py-1 rounded ${
              sortBy === 'price-desc'
                ? 'bg-primary text-white'
                : 'bg-gray-100 hover:bg-gray-200'
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
