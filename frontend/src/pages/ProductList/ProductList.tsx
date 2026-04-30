import React, { useState, useEffect, useCallback } from 'react';
import { useSearchParams } from 'react-router-dom';
import { Shirt, Smartphone, Sparkles, Home, Apple, Dumbbell, Baby, BookOpen, Loader2 } from 'lucide-react';
import { ProductCard } from '../../components/ProductCard/ProductCard';
import { productApi } from '../../api';
import type { Product, Category } from '../../types/product';

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

// 将后端返回的商品数据映射为前端 Product 类型
const mapApiProduct = (apiProduct: any): Product => ({
  id: String(apiProduct.id),
  categoryId: String(apiProduct.category_id),
  name: apiProduct.name || '',
  description: apiProduct.description || '',
  price: apiProduct.price || 0,
  stock: apiProduct.stock || 0,
  sales: apiProduct.sales || 0,
  mainImage: apiProduct.main_image || '',
  image: apiProduct.main_image || '',
  images: typeof apiProduct.images === 'string' ? JSON.parse(apiProduct.images || '[]') : (apiProduct.images || []),
  specs: [],
  status: apiProduct.status === 1 ? 'active' : 'inactive',
  createdAt: apiProduct.created_at || new Date().toISOString(),
});

// 将后端返回的分类数据映射为前端 Category 类型
const mapApiCategory = (apiCategory: any): Category => ({
  id: String(apiCategory.Id),
  name: apiCategory.name || '',
  parentId: apiCategory.parent_id ? String(apiCategory.parent_id) : undefined,
  icon: apiCategory.icon || '',
  sortOrder: apiCategory.sort_order || 0,
});

// 排序参数映射
const sortParamMap: Record<string, string> = {
  'default': '',
  'price-asc': 'price_asc',
  'price-desc': 'price_desc',
  'sales': 'sales',
};

export const ProductList: React.FC = () => {
  const [searchParams, setSearchParams] = useSearchParams();
  const [selectedCategory, setSelectedCategory] = useState<string>(searchParams.get('category') || '');
  const [searchQuery, setSearchQuery] = useState<string>(searchParams.get('search') || '');
  const [sortBy, setSortBy] = useState<'default' | 'price-asc' | 'price-desc' | 'sales'>('default');

  // API 数据状态
  const [products, setProducts] = useState<Product[]>([]);
  const [categories, setCategories] = useState<Category[]>([]);
  const [total, setTotal] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);
  const [page, setPage] = useState<number>(1);
  const [pageSize] = useState<number>(20);

  // 监听 URL 参数变化
  useEffect(() => {
    const category = searchParams.get('category');
    const search = searchParams.get('search');
    setSelectedCategory(category || '');
    setSearchQuery(search || '');
    // 筛选条件变化时重置页码
    setPage(1);
  }, [searchParams]);

  // 获取分类列表
  useEffect(() => {
    const fetchCategories = async () => {
      try {
        const res: any = await productApi.getCategories();
        if (res?.code === 0 && Array.isArray(res.data)) {
          setCategories(res.data.map(mapApiCategory));
        }
      } catch (error) {
        console.error('获取分类失败:', error);
      }
    };
    fetchCategories();
  }, []);

  // 获取商品列表
  const fetchProducts = useCallback(async () => {
    setLoading(true);
    try {
      const params: Record<string, any> = { page, pageSize };
      if (selectedCategory) params.categoryId = Number(selectedCategory);
      if (searchQuery) params.keyword = searchQuery;
      if (sortBy !== 'default' && sortParamMap[sortBy]) params.sort = sortParamMap[sortBy];

      const res: any = await productApi.getProducts(params);
      if (res?.code === 0 && res.data) {
        const apiProducts = res.data.products || [];
        setProducts(apiProducts.map(mapApiProduct));
        setTotal(res.data.total || 0);
      } else {
        setProducts([]);
        setTotal(0);
      }
    } catch (error) {
      console.error('获取商品列表失败:', error);
      setProducts([]);
      setTotal(0);
    } finally {
      setLoading(false);
    }
  }, [selectedCategory, searchQuery, sortBy, page, pageSize]);

  useEffect(() => {
    fetchProducts();
  }, [fetchProducts]);

  // 切换分类时更新 URL
  const handleCategoryChange = (categoryId: string) => {
    setSelectedCategory(categoryId);
    const params: Record<string, string> = {};
    if (categoryId) params.category = categoryId;
    if (searchQuery) params.search = searchQuery;
    setSearchParams(params);
  };

  // 排序切换
  const handleSortChange = (sort: 'default' | 'price-asc' | 'price-desc' | 'sales') => {
    setSortBy(sort);
    setPage(1);
  };

  // 分页：上一页/下一页
  const totalPages = Math.ceil(total / pageSize);
  const handlePrevPage = () => setPage((p) => Math.max(1, p - 1));
  const handleNextPage = () => setPage((p) => Math.min(totalPages, p + 1));

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
          {categories.map((cat) => (
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
            onClick={() => handleSortChange('default')}
            className={`px-4 py-1 rounded-lg transition-all ${
              sortBy === 'default'
                ? 'bg-primary text-white shadow-md'
                : 'bg-white/50 dark:bg-gray-700/50 hover:bg-white/80 dark:hover:bg-gray-700/80'
            }`}
          >
            综合
          </button>
          <button
            onClick={() => handleSortChange('sales')}
            className={`px-4 py-1 rounded-lg transition-all ${
              sortBy === 'sales'
                ? 'bg-primary text-white shadow-md'
                : 'bg-white/50 dark:bg-gray-700/50 hover:bg-white/80 dark:hover:bg-gray-700/80'
            }`}
          >
            销量
          </button>
          <button
            onClick={() => handleSortChange('price-asc')}
            className={`px-4 py-1 rounded-lg transition-all ${
              sortBy === 'price-asc'
                ? 'bg-primary text-white shadow-md'
                : 'bg-white/50 dark:bg-gray-700/50 hover:bg-white/80 dark:hover:bg-gray-700/80'
            }`}
          >
            价格升序
          </button>
          <button
            onClick={() => handleSortChange('price-desc')}
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

      {/* 搜索/分类标题 */}
      {(searchQuery || selectedCategory) && (
        <div className="mb-4">
          <h1 className="text-2xl font-bold text-gray-800 dark:text-white">
            {searchQuery ? `搜索"${searchQuery}"` : categories.find(c => c.id === selectedCategory)?.name || '商品列表'}
          </h1>
          <p className="text-sm text-gray-500 dark:text-gray-400 mt-1">
            共找到 {total} 件商品
          </p>
        </div>
      )}

      {/* 加载状态 */}
      {loading ? (
        <div className="glass-card rounded-xl p-12 text-center">
          <Loader2 className="w-8 h-8 animate-spin text-primary mx-auto mb-4" />
          <p className="text-gray-500 dark:text-gray-400">加载中...</p>
        </div>
      ) : products.length > 0 ? (
        <>
          {/* 商品列表 */}
          <div className="grid grid-cols-4 gap-6">
            {products.map((product) => (
              <ProductCard key={product.id} product={product} />
            ))}
          </div>

          {/* 分页 */}
          {totalPages > 1 && (
            <div className="flex items-center justify-center space-x-4 mt-8">
              <button
                onClick={handlePrevPage}
                disabled={page <= 1}
                className={`px-4 py-2 rounded-lg transition-all ${
                  page <= 1
                    ? 'bg-gray-200 dark:bg-gray-700 text-gray-400 cursor-not-allowed'
                    : 'bg-white/50 dark:bg-gray-700/50 hover:bg-white/80 dark:hover:bg-gray-700/80 text-gray-700 dark:text-gray-300'
                }`}
              >
                上一页
              </button>
              <span className="text-sm text-gray-600 dark:text-gray-400">
                第 {page} / {totalPages} 页
              </span>
              <button
                onClick={handleNextPage}
                disabled={page >= totalPages}
                className={`px-4 py-2 rounded-lg transition-all ${
                  page >= totalPages
                    ? 'bg-gray-200 dark:bg-gray-700 text-gray-400 cursor-not-allowed'
                    : 'bg-white/50 dark:bg-gray-700/50 hover:bg-white/80 dark:hover:bg-gray-700/80 text-gray-700 dark:text-gray-300'
                }`}
              >
                下一页
              </button>
            </div>
          )}
        </>
      ) : (
        <div className="glass-card rounded-xl p-12 text-center">
          <p className="text-gray-500 dark:text-gray-400 text-lg">
            {searchQuery ? `未找到与"${searchQuery}"相关的商品` : '暂无商品'}
          </p>
        </div>
      )}
    </div>
  );
};
