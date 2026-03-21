import React, { useState, useEffect } from 'react';
import { useSearchParams } from 'react-router-dom';
import { Shirt, Smartphone, Sparkles, Home, Apple, Dumbbell, Baby, BookOpen } from 'lucide-react';
import { ProductCard } from '../../components/ProductCard/ProductCard';
import { mockProducts, mockCategories } from '../../utils/mockData';
import type { Product } from '../../types/product';

// 分类图标映射
const categoryIcons: Record<string, React.ReactNode> = {
  '服装': <Shirt className="w-4 h-4" />,
  '数码': <Smartphone className="w-4 h-4" />,
  '美妆': <Sparkles className="w-4 h-4" />,
  '家居': <Home className="w-4 h-4" />,
  '食品': <Apple className="w-4 h-4" />,
  '运动': <Dumbbell className="w-4 h-4" />,
  '母婴': <Baby className="w-4 h-4" />,
  '图书': <BookOpen className="w-4 h-4" />,
};

export const ProductList: React.FC = () => {
  const [searchParams, setSearchParams] = useSearchParams();
  const [selectedCategory, setSelectedCategory] = useState<string>(searchParams.get('category') || '');
  const [sortBy, setSortBy] = useState<'default' | 'price-asc' | 'price-desc' | 'sales'>('default');

  // 监听 URL 参数变化
  useEffect(() => {
    const category = searchParams.get('category');
    if (category) {
      setSelectedCategory(category);
    }
  }, [searchParams]);

  // 切换分类时更新 URL
  const handleCategoryChange = (categoryId: string) => {
    setSelectedCategory(categoryId);
    if (categoryId) {
      setSearchParams({ category: categoryId });
    } else {
      setSearchParams({});
    }
  };

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
            onClick={() => handleCategoryChange('')}
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
              onClick={() => handleCategoryChange(cat.id)}
              className={`px-4 py-1.5 rounded-lg transition-all flex items-center space-x-1.5 ${
                selectedCategory === cat.id
                  ? 'bg-primary text-white shadow-md'
                  : 'bg-white/50 dark:bg-gray-700/50 hover:bg-white/80 dark:hover:bg-gray-700/80 text-gray-700 dark:text-gray-300'
              }`}
            >
              {categoryIcons[cat.name]}
              <span>{cat.name}</span>
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

      {/* 分类标题 */}
      {selectedCategory && (
        <div className="mb-4">
          <h1 className="text-2xl font-bold text-gray-800 dark:text-white">
            {mockCategories.find(c => c.id === selectedCategory)?.name || '商品列表'}
          </h1>
          <p className="text-sm text-gray-500 dark:text-gray-400 mt-1">
            共找到 {sortedProducts.length} 件商品
          </p>
        </div>
      )}

      {/* 商品列表 */}
      <div className="grid grid-cols-4 gap-6">
        {sortedProducts.map((product) => (
          <ProductCard key={product.id} product={product as Product} />
        ))}
      </div>
    </div>
  );
};
