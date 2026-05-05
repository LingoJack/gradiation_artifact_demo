import React, { useState, useEffect, useCallback } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import {
  ShoppingCart,
  Heart,
  Share2,
  Shield,
  Truck,
  RotateCcw,
  ChevronRight,
  Star,
  ThumbsUp
} from 'lucide-react';
import { productApi, cartApi, favoriteApi } from '../../api';
import { useCartStore } from '../../store/useCartStore';
import { useFavoriteStore } from '../../store/useFavoriteStore';
import { useBrowseHistoryStore } from '../../store/useBrowseHistoryStore';
import { showToast, copyToClipboard } from '../../utils/toast';

// 格式化数字
const formatNumber = (num: number) => {
  if (num >= 10000) {
    return `${(num / 10000).toFixed(1)}万`;
  }
  return num.toString();
};

// 分类数据
const categories: Record<number, string> = {
  1: '服装',
  2: '数码',
  3: '美妆',
  4: '家居',
  5: '食品',
  6: '运动',
  7: '母婴',
  8: '图书',
};

// 后端 SKU 类型
interface ApiSku {
  id: number;
  product_id: number;
  sku_code: string;
  spec_values: Record<string, string> | string;
  price: number;
  stock: number;
  image: string;
  status: number;
}

// 后端评价类型
interface ApiReview {
  id?: number;
  username?: string;
  user_name?: string;
  avatar?: string;
  rating?: number;
  content?: string;
  images?: string[] | string;
  specs?: string;
  time?: string;
  created_at?: string;
  likes?: number;
  reply?: string;
  badge?: string;
}

// 映射后的商品数据
interface ProductData {
  id: string;
  name: string;
  price: number;
  originalPrice?: number;
  mainImage: string;
  stock: number;
  sales: number;
  categoryId: string;
  status: number;
  description: string;
  images: string[];
  skus: ApiSku[];
  reviews: ApiReview[];
  shopName?: string;
}

// API 返回的商品详情类型
interface ApiProductDetail {
  id: number;
  name: string;
  price: number | string;
  main_image?: string;
  mainImage?: string;
  stock: number | string;
  sales: number | string;
  category_id?: number;
  categoryId?: number;
  status?: number;
  description?: string;
  images?: string[] | string;
  skus?: ApiSku[];
  reviews?: ApiReview[];
  shop_name?: string;
  shopName?: string;
}

// 解析 spec_values（可能是对象或 JSON 字符串）
const parseSpecValues = (specValues: Record<string, string> | string): Record<string, string> => {
  if (typeof specValues === 'string') {
    try {
      return JSON.parse(specValues);
    } catch {
      return {};
    }
  }
  return specValues;
};

// 解析 images（可能是数组或 JSON 字符串）
const parseImages = (images: string[] | string): string[] => {
  if (typeof images === 'string') {
    try {
      return JSON.parse(images);
    } catch {
      return [];
    }
  }
  return images || [];
};

// 从所有 SKU 中提取规格维度
const extractSpecDimensions = (skus: ApiSku[]): { key: string; values: string[] }[] => {
  const dimensionMap = new Map<string, Set<string>>();

  skus.forEach((sku) => {
    if (sku.status !== 1) return;
    const specValues = parseSpecValues(sku.spec_values);
    Object.entries(specValues).forEach(([key, value]) => {
      if (!dimensionMap.has(key)) {
        dimensionMap.set(key, new Set());
      }
      dimensionMap.get(key)!.add(value);
    });
  });

  return Array.from(dimensionMap.entries()).map(([key, values]) => ({
    key,
    values: Array.from(values),
  }));
};

// 格式化评价数据
const formatReviews = (reviews: ApiReview[]) => {
  return reviews.map((review) => ({
    id: review.id ?? 0,
    username: review.username || review.user_name || '匿名用户',
    avatar: review.avatar || '',
    badge: review.badge || null,
    rating: review.rating || 5,
    time: review.time || (review.created_at
      ? new Date(review.created_at).toLocaleDateString()
      : '未知时间'),
    content: review.content || '',
    specs: review.specs || '',
    images: parseImages(review.images || []),
    likes: review.likes || 0,
    reply: review.reply || null,
  }));
};

export const ProductDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const addItemToStore = useCartStore((s) => s.addItem);
  const addFavoriteToStore = useFavoriteStore((s) => s.addFavorite);
  const removeFavoriteFromStore = useFavoriteStore((s) => s.removeFavorite);
  const addBrowseHistory = useBrowseHistoryStore((s) => s.addItem);

  const [product, setProduct] = useState<ProductData | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [selectedSpecs, setSelectedSpecs] = useState<Record<string, string>>({});
  const [quantity, setQuantity] = useState(1);
  const [showAddedToCart, setShowAddedToCart] = useState(false);
  const [addingToCart, setAddingToCart] = useState(false);
  const [mainImage, setMainImage] = useState<string>('');
  const [isFav, setIsFav] = useState(false);

  // 获取商品详情
  const fetchProduct = useCallback(async () => {
    if (!id) return;
    setLoading(true);
    setError(null);
    try {
      const data = await productApi.getProductDetail(Number(id)) as ApiProductDetail;

      const mappedProduct: ProductData = {
        id: String(data.id ?? data.id ?? ''),
        name: data.name || '',
        price: Number(data.price) || 0,
        mainImage: data.main_image || data.mainImage || '',
        stock: Number(data.stock) || 0,
        sales: Number(data.sales) || 0,
        categoryId: String(data.category_id ?? data.categoryId ?? ''),
        status: data.status ?? 1,
        description: data.description || '',
        images: parseImages(data.images || []),
        skus: (data.skus || []).map((sku: ApiSku) => ({
          ...sku,
          spec_values: parseSpecValues(sku.spec_values),
        })),
        reviews: data.reviews || [],
        shopName: data.shop_name || data.shopName || '',
      };

      setProduct(mappedProduct);
      setMainImage(mappedProduct.mainImage);

      // 记录浏览历史
      addBrowseHistory({
        id: mappedProduct.id,
        name: mappedProduct.name,
        price: mappedProduct.price,
        mainImage: mappedProduct.mainImage,
      });

      // 检查收藏状态
      try {
        const favResult: any = await favoriteApi.checkFavorite(Number(id));
        setIsFav(favResult?.isFavorite ?? favResult?.is_favorite ?? false);
      } catch {
        // 未登录或其他错误，忽略
      }
    } catch (err: any) {
      setError(err.message || '获取商品详情失败');
    } finally {
      setLoading(false);
    }
  }, [id]);

  useEffect(() => {
    fetchProduct();
  }, [fetchProduct]);

  // 商品切换时重置规格和数量
  useEffect(() => {
    setSelectedSpecs({});
    setQuantity(1);
    setMainImage('');
  }, [id]);

  // 设置默认主图
  useEffect(() => {
    if (product && !mainImage) {
      setMainImage(product.mainImage);
    }
  }, [product, mainImage]);

  // 规格维度
  const specDimensions = product ? extractSpecDimensions(product.skus) : [];

  // 当前选中的 SKU（所有规格维度都匹配时）
  const selectedSku = product?.skus?.find((sku) => {
    if (sku.status !== 1) return false;
    const specValues = parseSpecValues(sku.spec_values);
    return Object.entries(selectedSpecs).every(
      ([key, value]) => specValues[key] === value
    );
  });

  // 规格是否全部选择完毕
  const allSpecsSelected = specDimensions.length > 0
    ? specDimensions.every((dim) => selectedSpecs[dim.key] !== undefined)
    : true;

  const currentPrice = selectedSku?.price || product?.price || 0;
  const currentStock = selectedSku?.stock || product?.stock || 0;

  // 处理规格选择
  const handleSpecSelect = (key: string, value: string) => {
    setSelectedSpecs((prev) => ({ ...prev, [key]: value }));
  };

  // 添加到购物车
  const handleAddToCart = async () => {
    if (!product || addingToCart) return;

    try {
      setAddingToCart(true);
      // 调用后端 API
      await cartApi.addItem({
        productId: Number(product.id),
        skuId: selectedSku ? Number(selectedSku.id) : undefined,
        quantity,
      });

      // 同步更新本地 store
      const specDisplay = selectedSku
        ? Object.entries(parseSpecValues(selectedSku.spec_values))
            .map(([, v]) => v)
            .join(' | ')
        : undefined;

      addItemToStore({
        id: `${product.id}-${selectedSku?.id || 'default'}`,
        userId: '1',
        productId: product.id,
        product: {
          id: product.id,
          name: product.name,
          price: currentPrice,
          mainImage: product.mainImage,
          stock: currentStock,
        },
        specId: selectedSku ? String(selectedSku.id) : undefined,
        spec: selectedSku
          ? {
              id: String(selectedSku.id),
              name: specDisplay || '',
              value: specDisplay || '',
              price: selectedSku.price,
              stock: selectedSku.stock,
            }
          : undefined,
        quantity,
        selected: true,
      });

      setShowAddedToCart(true);
      setTimeout(() => setShowAddedToCart(false), 2000);
      showToast('已添加到购物车', 'success');
    } catch (err: any) {
      showToast(err.message || '添加购物车失败', 'error');
    } finally {
      setAddingToCart(false);
    }
  };

  // 立即购买
  const handleBuyNow = async () => {
    try {
      await handleAddToCart();
      navigate('/cart');
    } catch (err: any) {
      // handleAddToCart 内部已 showToast，这里不再重复提示
      console.error('立即购买失败:', err);
    }
  };

  // 切换收藏
  const handleToggleFavorite = async () => {
    if (!product) return;
    try {
      if (isFav) {
        await favoriteApi.removeFavorite(Number(product.id));
        removeFavoriteFromStore(product.id);
        setIsFav(false);
        showToast('已取消收藏', 'info');
      } else {
        await favoriteApi.addFavorite(Number(product.id));
        addFavoriteToStore(product.id);
        setIsFav(true);
        showToast('收藏成功！', 'success');
      }
    } catch (err: any) {
      showToast(err.message || '操作失败', 'error');
    }
  };

  // Loading 状态
  if (loading) {
    return (
      <div className="min-h-screen bg-gray-50 dark:bg-gray-900 flex items-center justify-center">
        <div className="text-center">
          <div className="w-12 h-12 border-4 border-orange-500 border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
          <p className="text-gray-500 dark:text-gray-400">加载中...</p>
        </div>
      </div>
    );
  }

  // 错误或商品不存在
  if (error || !product) {
    return (
      <div className="min-h-screen bg-gray-50 dark:bg-gray-900 flex items-center justify-center">
        <div className="text-center">
          <p className="text-2xl text-gray-500 dark:text-gray-400 mb-4">
            {error || '商品不存在'}
          </p>
          <button
            onClick={() => navigate('/')}
            className="text-orange-500 hover:text-orange-600"
          >
            返回首页
          </button>
        </div>
      </div>
    );
  }

  const formattedReviews = formatReviews(product.reviews);

  return (
    <div className="min-h-screen bg-gray-50 dark:bg-gray-900">
      {/* 面包屑导航 */}
      <div className="glass-nav">
        <div className="max-w-[1200px] mx-auto px-4 py-3">
          <div className="flex items-center text-sm text-gray-600 dark:text-gray-400">
            <span
              className="hover:text-orange-500 cursor-pointer"
              onClick={() => navigate('/')}
            >
              首页
            </span>
            <ChevronRight className="w-4 h-4 mx-2" />
            <span className="hover:text-orange-500 cursor-pointer">
              {categories[Number(product.categoryId)] || '商品'}
            </span>
            <ChevronRight className="w-4 h-4 mx-2" />
            <span className="text-gray-900 dark:text-gray-100">{product.name}</span>
          </div>
        </div>
      </div>

      {/* 主要内容区 */}
      <div className="max-w-[1200px] mx-auto px-4 py-6">
        <div className="glass-card rounded-2xl p-6">
          <div className="grid grid-cols-12 gap-8">
            {/* 左侧图片区 */}
            <div className="col-span-5">
              {/* 主图 */}
              <div className="aspect-square bg-gray-100 dark:bg-gray-700 rounded-lg overflow-hidden mb-4">
                <img
                  src={mainImage || product.mainImage}
                  alt={product.name}
                  className="w-full h-full object-cover"
                  onError={(e) => {
                    const target = e.target as HTMLImageElement;
                    target.src = 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="400" height="400"%3E%3Crect fill="%23f3f4f6" width="400" height="400"/%3E%3Ctext fill="%239ca3af" x="50%25" y="50%25" text-anchor="middle" dy=".3em"%3E暂无图片%3C/text%3E%3C/svg%3E';
                  }}
                />
              </div>

              {/* 缩略图 */}
              {product.images.length > 0 && (
                <div className="grid grid-cols-5 gap-2">
                  {product.images.map((img, index) => (
                    <div
                      key={`${img}-${index}`}
                      onClick={() => setMainImage(img)}
                      className={`aspect-square bg-gray-100 dark:bg-gray-700 rounded overflow-hidden cursor-pointer border-2 transition-all ${
                        mainImage === img
                          ? 'border-orange-500'
                          : 'border-transparent hover:border-gray-300 dark:hover:border-gray-600'
                      }`}
                    >
                      <img
                        src={img}
                        alt={`${product.name} ${index + 1}`}
                        className="w-full h-full object-cover"
                        onError={(e) => {
                          const target = e.target as HTMLImageElement;
                          target.style.display = 'none';
                        }}
                      />
                    </div>
                  ))}
                </div>
              )}

              {/* 分享收藏 */}
              <div className="flex items-center justify-end gap-4 mt-4 pt-4 border-t border-gray-200 dark:border-gray-700">
                <button
                  onClick={handleToggleFavorite}
                  className={`flex items-center gap-1 text-sm ${
                    isFav ? 'text-red-500' : 'text-gray-600 dark:text-gray-400 hover:text-red-500'
                  }`}
                >
                  <Heart className={`w-4 h-4 ${isFav ? 'fill-current' : ''}`} />
                  收藏
                </button>
                <button
                  onClick={async () => {
                    const url = window.location.href;
                    const success = await copyToClipboard(url);
                    if (success) {
                      showToast('链接已复制到剪贴板！', 'success');
                    } else {
                      showToast('复制失败，请手动复制', 'error');
                    }
                  }}
                  className="flex items-center gap-1 text-sm text-gray-600 dark:text-gray-400 hover:text-blue-500"
                >
                  <Share2 className="w-4 h-4" />
                  分享
                </button>
              </div>
            </div>

            {/* 右侧信息区 */}
            <div className="col-span-7">
              {/* 标题 */}
              <h1 className="text-xl font-medium text-gray-900 dark:text-gray-100 mb-3 line-clamp-2">
                {product.name}
              </h1>

              {/* 描述 */}
              <p className="text-sm text-gray-600 dark:text-gray-400 mb-4">{product.description}</p>

              {/* 价格区域 */}
              <div className="price-glass rounded-xl p-5 mb-5">
                <div className="flex items-baseline gap-3 mb-2">
                  <span className="text-sm text-gray-600 dark:text-gray-400">促销价</span>
                  <span className="text-sm text-red-600 dark:text-red-400 font-medium">¥</span>
                  <span className="text-4xl font-bold text-red-600 dark:text-red-400">
                    {currentPrice}
                  </span>
                  {product.originalPrice && (
                    <span className="text-sm text-gray-400 line-through">
                      ¥{product.originalPrice}
                    </span>
                  )}
                  {product.originalPrice && (
                    <span className="bg-red-600 text-white text-xs px-2 py-1 rounded">
                      省¥{product.originalPrice - currentPrice}
                    </span>
                  )}
                </div>
                {/* 促销标签 */}
                <div className="flex items-center gap-2 mt-3">
                  <span className="bg-red-600 text-white text-xs px-2 py-1 rounded">
                    限时特惠
                  </span>
                  <span className="bg-orange-600 text-white text-xs px-2 py-1 rounded">
                    包邮
                  </span>
                  <span className="text-xs text-gray-600 dark:text-gray-400">
                    销量 {product.sales}+ 已售
                  </span>
                </div>
              </div>

              {/* 服务保障 */}
              <div className="flex items-center gap-4 mb-5 text-sm text-gray-600 dark:text-gray-400">
                <div className="flex items-center gap-1">
                  <Shield className="w-4 h-4 text-blue-500" />
                  <span>正品保障</span>
                </div>
                <div className="flex items-center gap-1">
                  <Truck className="w-4 h-4 text-blue-500" />
                  <span>极速发货</span>
                </div>
                <div className="flex items-center gap-1">
                  <RotateCcw className="w-4 h-4 text-blue-500" />
                  <span>7天无理由</span>
                </div>
              </div>

              {/* 规格选择 - 支持多维度规格 */}
              {specDimensions.length > 0 && (
                <div className="mb-5">
                  {specDimensions.map((dimension) => (
                    <div key={dimension.key} className="flex items-center gap-4 mb-3">
                      <span className="text-sm text-gray-600 dark:text-gray-400 w-16">
                        {dimension.key}
                      </span>
                      <div className="flex flex-wrap gap-2">
                        {dimension.values.map((value) => (
                          <button
                            key={value}
                            onClick={() => handleSpecSelect(dimension.key, value)}
                            className={`px-4 py-2 border-2 rounded-lg text-sm transition-all ${
                              selectedSpecs[dimension.key] === value
                                ? 'border-orange-500 bg-orange-50 dark:bg-orange-900/30 text-orange-600 dark:text-orange-400'
                                : 'border-gray-300 dark:border-gray-600 hover:border-orange-500 dark:hover:border-orange-500 text-gray-900 dark:text-gray-100'
                            }`}
                          >
                            {value}
                          </button>
                        ))}
                      </div>
                    </div>
                  ))}
                </div>
              )}

              {/* 数量选择 */}
              <div className="flex items-center gap-4 mb-5">
                <span className="text-sm text-gray-600 dark:text-gray-400 w-16">数量</span>
                <div className="flex items-center">
                  <button
                    onClick={() => setQuantity(Math.max(1, quantity - 1))}
                    disabled={currentStock === 0}
                    className="w-8 h-8 border border-gray-300 dark:border-gray-600 rounded-l flex items-center justify-center hover:bg-gray-100 dark:hover:bg-gray-700 text-gray-900 dark:text-gray-100 disabled:opacity-50 disabled:cursor-not-allowed"
                  >
                    -
                  </button>
                  <input
                    type="number"
                    value={quantity}
                    onChange={(e) => {
                      const value = parseInt(e.target.value) || 1;
                      setQuantity(Math.max(1, Math.min(currentStock > 0 ? currentStock : 1, value)));
                    }}
                    disabled={currentStock === 0}
                    className="w-16 h-8 border-t border-b border-gray-300 dark:border-gray-600 text-center outline-none bg-white dark:bg-gray-800 text-gray-900 dark:text-gray-100 disabled:opacity-50 disabled:cursor-not-allowed"
                  />
                  <button
                    onClick={() =>
                      setQuantity(Math.min(currentStock > 0 ? currentStock : 1, quantity + 1))
                    }
                    disabled={currentStock === 0}
                    className="w-8 h-8 border border-gray-300 dark:border-gray-600 rounded-r flex items-center justify-center hover:bg-gray-100 dark:hover:bg-gray-700 text-gray-900 dark:text-gray-100 disabled:opacity-50 disabled:cursor-not-allowed"
                  >
                    +
                  </button>
                  <span className="ml-3 text-sm text-gray-500 dark:text-gray-400">
                    {currentStock === 0 ? (
                      <span className="text-red-500 font-medium">已售罄</span>
                    ) : (
                      `库存 ${currentStock} 件`
                    )}
                  </span>
                </div>
              </div>

              {/* 购买按钮 */}
              <div className="flex gap-3 mb-6">
                <button
                  onClick={handleAddToCart}
                  disabled={addingToCart || currentStock === 0 || (!allSpecsSelected && specDimensions.length > 0)}
                  className={`flex-1 h-12 border-2 border-orange-500 text-orange-600 dark:text-orange-400 rounded-lg hover:bg-orange-50 dark:hover:bg-orange-900/20 flex items-center justify-center gap-2 font-medium transition-colors disabled:opacity-50 disabled:cursor-not-allowed`}
                >
                  <ShoppingCart className="w-5 h-5" />
                  {addingToCart ? '添加中...' : currentStock === 0 ? '已售罄' : '加入购物车'}
                </button>
                <button
                  onClick={handleBuyNow}
                  disabled={addingToCart || currentStock === 0 || (!allSpecsSelected && specDimensions.length > 0)}
                  className={`flex-1 h-12 bg-gradient-to-r from-orange-500 to-red-500 text-white rounded-lg hover:from-orange-600 hover:to-red-600 font-medium transition-all disabled:opacity-50 disabled:cursor-not-allowed`}
                >
                  {currentStock === 0 ? '已售罄' : '立即购买'}
                </button>
              </div>

              {showAddedToCart && (
                <div className="bg-green-100 dark:bg-green-900/30 text-green-800 dark:text-green-300 p-3 rounded-lg text-center mb-4 animate-pulse">
                  ✓ 已添加到购物车
                </div>
              )}

              {/* 店铺信息 */}
              <div className="border-t border-gray-200 dark:border-gray-700 pt-5">
                <div className="flex items-center justify-between">
                  <div className="flex items-center gap-3">
                    <div className="w-12 h-12 bg-gradient-to-br from-orange-400 to-red-400 rounded-lg flex items-center justify-center overflow-hidden">
                      <span className="text-white font-bold text-sm">
                        {(product.shopName || '淘宝店铺').slice(0, 2)}
                      </span>
                    </div>
                    <div>
                      <div className="font-medium text-gray-900 dark:text-gray-100">
                        {product.shopName || '淘宝官方店铺'}
                      </div>
                      <div className="flex items-center gap-2 text-xs text-gray-500 dark:text-gray-400 mt-1">
                        <div className="flex items-center gap-1">
                          <Star className="w-3 h-3 text-yellow-400 fill-current" />
                          <span>4.9</span>
                        </div>
                        <span>|</span>
                        <span>粉丝 {formatNumber(10000)}</span>
                      </div>
                    </div>
                  </div>
                  <div className="flex gap-2">
                    <button
                      onClick={() => showToast('进店功能开发中', 'info')}
                      className="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded text-sm hover:bg-gray-50 dark:hover:bg-gray-700 text-gray-900 dark:text-gray-100"
                    >
                      进店逛逛
                    </button>
                    <button
                      onClick={() => showToast('关注成功！', 'success')}
                      className="px-4 py-2 border border-orange-500 text-orange-600 dark:text-orange-400 rounded text-sm hover:bg-orange-50 dark:hover:bg-orange-900/20"
                    >
                      关注店铺
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        {/* 商品详情区 */}
        <div className="mt-6 glass-card rounded-2xl p-6">
          <div className="border-b border-gray-200/50 dark:border-gray-700/50 pb-4 mb-6">
            <h2 className="text-xl font-bold text-gray-900 dark:text-gray-100">商品详情</h2>
          </div>

          {/* 商品参数 */}
          <div className="mb-8">
            <h3 className="font-medium text-gray-900 dark:text-gray-100 mb-4">商品参数</h3>
            <div className="grid grid-cols-4 gap-4 text-sm">
              <div className="bg-gray-50 dark:bg-gray-700 p-3 rounded">
                <span className="text-gray-600 dark:text-gray-400">品牌：</span>
                <span className="text-gray-900 dark:text-gray-100">{(product.shopName || '淘宝店铺').split('官方')[0]}</span>
              </div>
              <div className="bg-gray-50 dark:bg-gray-700 p-3 rounded">
                <span className="text-gray-600 dark:text-gray-400">产地：</span>
                <span className="text-gray-900 dark:text-gray-100">中国</span>
              </div>
              <div className="bg-gray-50 dark:bg-gray-700 p-3 rounded">
                <span className="text-gray-600 dark:text-gray-400">分类：</span>
                <span className="text-gray-900 dark:text-gray-100">
                  {categories[Number(product.categoryId)] || '其他'}
                </span>
              </div>
              <div className="bg-gray-50 dark:bg-gray-700 p-3 rounded">
                <span className="text-gray-600 dark:text-gray-400">上市时间：</span>
                <span className="text-gray-900 dark:text-gray-100">2024年</span>
              </div>
            </div>
          </div>

          {/* 商品图片展示 */}
          {product.images.length > 0 && (
            <div>
              <h3 className="font-medium text-gray-900 dark:text-gray-100 mb-4">商品展示</h3>
              <div className="space-y-4">
                {product.images.map((img, index) => (
                  <div key={`${img}-${index}`} className="w-full bg-gray-100 dark:bg-gray-700 rounded-lg overflow-hidden">
                    <img
                      src={img}
                      alt={`${product.name} 详情图 ${index + 1}`}
                      className="w-full h-auto"
                      onError={(e) => {
                        const target = e.target as HTMLImageElement;
                        target.style.display = 'none';
                      }}
                    />
                  </div>
                ))}
              </div>
            </div>
          )}

          {/* 用户评价 */}
          <div className="mt-8">
            {/* 评价头部 */}
            <div className="flex items-center justify-between mb-6">
              <div className="flex items-center gap-3">
                <h3 className="text-xl font-semibold text-gray-900 dark:text-gray-100">用户评价</h3>
                <span className="text-sm text-gray-500 dark:text-gray-400">({formattedReviews.length})</span>
              </div>

              {/* 好评率卡片 */}
              <div className="flex items-center gap-4 bg-gradient-to-r from-orange-50 to-red-50 dark:from-orange-900/20 dark:to-red-900/20 px-6 py-3 rounded-xl border border-orange-200 dark:border-orange-800">
                <div className="text-center">
                  <div className="text-2xl font-bold text-orange-600 dark:text-orange-400">98%</div>
                  <div className="text-xs text-gray-600 dark:text-gray-400">好评率</div>
                </div>
              </div>
            </div>

            {/* 标签筛选 */}
            <div className="flex items-center gap-3 mb-6">
              {['全部评价', '有图评价', '好评', '中评', '差评'].map((tag, index) => (
                <button
                  key={tag}
                  className={`px-4 py-2 rounded-full text-sm font-medium transition-all ${
                    index === 0
                      ? 'bg-orange-500 text-white'
                      : 'bg-gray-100 dark:bg-gray-700 text-gray-700 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-600'
                  }`}
                >
                  {tag}
                </button>
              ))}
            </div>

            {/* 评价列表 */}
            {formattedReviews.length > 0 ? (
              <div className="space-y-4">
                {formattedReviews.map((review) => (
                  <div
                    key={review.id}
                    className="glass-card rounded-xl p-6 hover:shadow-xl transition-all duration-300"
                  >
                    {/* 用户信息 */}
                    <div className="flex items-start justify-between mb-4">
                      <div className="flex items-center gap-3">
                        {review.avatar ? (
                          <img
                            src={review.avatar}
                            alt="用户头像"
                            className="w-10 h-10 rounded-full object-cover"
                            onError={(e) => {
                              const target = e.target as HTMLImageElement;
                              target.src = `https://ui-avatars.com/api/?name=${review.username}&background=f97316&color=fff`;
                            }}
                          />
                        ) : (
                          <div className="w-10 h-10 rounded-full bg-gradient-to-br from-orange-400 to-red-400 flex items-center justify-center text-white text-sm font-bold">
                            {review.username.slice(0, 1)}
                          </div>
                        )}
                        <div>
                          <div className="flex items-center gap-2">
                            <span className="font-medium text-gray-900 dark:text-gray-100">{review.username}</span>
                            {review.badge && (
                              <span className="px-2 py-0.5 bg-gradient-to-r from-orange-400 to-red-400 text-white text-xs rounded-full">
                                {review.badge}
                              </span>
                            )}
                          </div>
                          <div className="flex items-center gap-2 mt-1">
                            <div className="flex items-center gap-0.5">
                              {[1, 2, 3, 4, 5].map((star) => (
                                <Star
                                  key={star}
                                  className={`w-3.5 h-3.5 ${
                                    star <= review.rating
                                      ? 'text-orange-400 fill-current'
                                      : 'text-gray-300 dark:text-gray-600'
                                  }`}
                                />
                              ))}
                            </div>
                            <span className="text-xs text-gray-400">{review.time}</span>
                          </div>
                        </div>
                      </div>
                    </div>

                    {/* 评价内容 */}
                    <div className="mb-4">
                      <p className="text-gray-700 dark:text-gray-300 text-sm leading-relaxed mb-2">
                        {review.content}
                      </p>
                      {review.specs && (
                        <div className="inline-block px-3 py-1 bg-gray-100 dark:bg-gray-700 rounded text-xs text-gray-600 dark:text-gray-400">
                          规格：{review.specs}
                        </div>
                      )}
                    </div>

                    {/* 评价图片 */}
                    {review.images.length > 0 && (
                      <div className="grid grid-cols-3 gap-2 mb-4">
                        {review.images.map((img: string, idx: number) => (
                          <div
                            key={idx}
                            className="aspect-square rounded-lg overflow-hidden bg-gray-100 dark:bg-gray-700"
                          >
                            <img
                              src={img}
                              alt={`评价图片 ${idx + 1}`}
                              className="w-full h-full object-cover hover:scale-105 transition-transform cursor-pointer"
                              onError={(e) => {
                                const target = e.target as HTMLImageElement;
                                target.style.display = 'none';
                              }}
                            />
                          </div>
                        ))}
                      </div>
                    )}

                    {/* 操作按钮 */}
                    <div className="flex items-center justify-between pt-4 border-t border-gray-100 dark:border-gray-700">
                      <div className="flex items-center gap-3">
                        <button
                          onClick={() => showToast('点赞成功！', 'success')}
                          className="flex items-center gap-1.5 text-sm text-gray-500 dark:text-gray-400 hover:text-orange-500 dark:hover:text-orange-400 transition-colors"
                        >
                          <ThumbsUp className="w-4 h-4" />
                          <span>{review.likes}</span>
                        </button>
                        <button
                          onClick={() => showToast('回复功能开发中', 'info')}
                          className="text-sm text-gray-500 dark:text-gray-400 hover:text-orange-500 dark:hover:text-orange-400 transition-colors"
                        >
                          回复
                        </button>
                      </div>
                    </div>

                    {/* 商家回复 */}
                    {review.reply && (
                      <div className="mt-4 bg-gray-50 dark:bg-gray-700/50 rounded-lg p-4">
                        <div className="flex items-center gap-2 mb-2">
                          <span className="text-sm font-medium text-orange-600 dark:text-orange-400">商家回复：</span>
                          <span className="text-xs text-gray-400">2天前</span>
                        </div>
                        <p className="text-sm text-gray-600 dark:text-gray-400">{review.reply}</p>
                      </div>
                    )}
                  </div>
                ))}
              </div>
            ) : (
              <div className="text-center py-12">
                <p className="text-gray-400 dark:text-gray-500">暂无评价</p>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};
