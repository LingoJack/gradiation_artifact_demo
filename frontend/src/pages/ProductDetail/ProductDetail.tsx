import React, { useState } from 'react';
import { useParams, useNavigate, Link } from 'react-router-dom';
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
import { mockProducts, getShopByName, getProductReviews } from '../../utils/mockData';
import { useCartStore } from '../../store/useCartStore';
import type { Product } from '../../types/product';

// 格式化数字
const formatNumber = (num: number) => {
  if (num >= 10000) {
    return `${(num / 10000).toFixed(1)}万`;
  }
  return num.toString();
};

// 分类数据
const mockCategories = [
  { id: '1', name: '服装' },
  { id: '2', name: '数码' },
  { id: '3', name: '美妆' },
  { id: '4', name: '家居' },
  { id: '5', name: '食品' },
  { id: '6', name: '运动' },
  { id: '7', name: '母婴' },
  { id: '8', name: '图书' },
];

export const ProductDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const { addItem } = useCartStore();
  const [selectedSpec, setSelectedSpec] = useState<string>('');
  const [quantity, setQuantity] = useState(1);
  const [showAddedToCart, setShowAddedToCart] = useState(false);
  const [mainImage, setMainImage] = useState<string>('');
  const [isFavorite, setIsFavorite] = useState(false);

  const product = mockProducts.find((p) => p.id === id) as Product | undefined;

  React.useEffect(() => {
    if (product && !mainImage) {
      setMainImage(product.mainImage);
    }
  }, [product, mainImage]);

  if (!product) {
    return (
      <div className="min-h-screen bg-gray-50 dark:bg-gray-900 flex items-center justify-center">
        <div className="text-center">
          <p className="text-2xl text-gray-500 dark:text-gray-400 mb-4">商品不存在</p>
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

  const handleAddToCart = () => {
    const spec = selectedSpec
      ? product.specs.find((s) => s.id === selectedSpec)
      : undefined;

    addItem({
      id: `${product.id}-${selectedSpec || 'default'}`,
      userId: '1',
      productId: product.id,
      product: {
        id: product.id,
        name: product.name,
        price: spec?.price || product.price,
        mainImage: product.mainImage,
        stock: spec?.stock || product.stock,
      },
      specId: selectedSpec,
      spec: spec
        ? {
            id: spec.id,
            name: spec.name,
            value: spec.value,
            price: spec.price,
            stock: spec.stock,
          }
        : undefined,
      quantity,
      selected: true,
    });

    setShowAddedToCart(true);
    setTimeout(() => setShowAddedToCart(false), 2000);
  };

  const handleBuyNow = () => {
    handleAddToCart();
    navigate('/cart');
  };

  const currentPrice = selectedSpec
    ? product.specs.find((s) => s.id === selectedSpec)?.price || product.price
    : product.price;

  const currentStock = selectedSpec
    ? product.specs.find((s) => s.id === selectedSpec)?.stock || product.stock
    : product.stock;

  return (
    <div className="min-h-screen bg-gray-50 dark:bg-gray-900">
      {/* 面包屑导航 */}
      <div className="glass-nav">
        <div className="max-w-[1200px] mx-auto px-4 py-3">
          <div className="flex items-center text-sm text-gray-600 dark:text-gray-400">
            <span className="hover:text-orange-500 cursor-pointer">首页</span>
            <ChevronRight className="w-4 h-4 mx-2" />
            <span className="hover:text-orange-500 cursor-pointer">
              {mockCategories.find(c => c.id === product.categoryId)?.name || '商品'}
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
              <div className="grid grid-cols-5 gap-2">
                {product.images.map((img, index) => (
                  <div
                    key={index}
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

              {/* 分享收藏 */}
              <div className="flex items-center justify-end gap-4 mt-4 pt-4 border-t border-gray-200 dark:border-gray-700">
                <button
                  onClick={() => setIsFavorite(!isFavorite)}
                  className={`flex items-center gap-1 text-sm ${
                    isFavorite ? 'text-red-500' : 'text-gray-600 dark:text-gray-400 hover:text-red-500'
                  }`}
                >
                  <Heart className={`w-4 h-4 ${isFavorite ? 'fill-current' : ''}`} />
                  收藏
                </button>
                <button className="flex items-center gap-1 text-sm text-gray-600 dark:text-gray-400 hover:text-blue-500">
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

              {/* 规格选择 */}
              {product.specs && product.specs.length > 0 && (
                <div className="mb-5">
                  <div className="flex items-center gap-4 mb-3">
                    <span className="text-sm text-gray-600 dark:text-gray-400 w-16">
                      {product.specs[0].name}
                    </span>
                    <div className="flex flex-wrap gap-2">
                      {product.specs.map((spec) => (
                        <button
                          key={spec.id}
                          onClick={() => setSelectedSpec(spec.id)}
                          className={`px-4 py-2 border-2 rounded-lg text-sm transition-all ${
                            selectedSpec === spec.id
                              ? 'border-orange-500 bg-orange-50 dark:bg-orange-900/30 text-orange-600 dark:text-orange-400'
                              : 'border-gray-300 dark:border-gray-600 hover:border-orange-500 dark:hover:border-orange-500 text-gray-900 dark:text-gray-100'
                          }`}
                        >
                          {spec.value}
                        </button>
                      ))}
                    </div>
                  </div>
                </div>
              )}

              {/* 数量选择 */}
              <div className="flex items-center gap-4 mb-5">
                <span className="text-sm text-gray-600 dark:text-gray-400 w-16">数量</span>
                <div className="flex items-center">
                  <button
                    onClick={() => setQuantity(Math.max(1, quantity - 1))}
                    className="w-8 h-8 border border-gray-300 dark:border-gray-600 rounded-l flex items-center justify-center hover:bg-gray-100 dark:hover:bg-gray-700 text-gray-900 dark:text-gray-100"
                  >
                    -
                  </button>
                  <input
                    type="number"
                    value={quantity}
                    onChange={(e) => {
                      const value = parseInt(e.target.value) || 1;
                      setQuantity(Math.max(1, Math.min(currentStock, value)));
                    }}
                    className="w-16 h-8 border-t border-b border-gray-300 dark:border-gray-600 text-center outline-none bg-white dark:bg-gray-800 text-gray-900 dark:text-gray-100"
                  />
                  <button
                    onClick={() =>
                      setQuantity(Math.min(currentStock, quantity + 1))
                    }
                    className="w-8 h-8 border border-gray-300 dark:border-gray-600 rounded-r flex items-center justify-center hover:bg-gray-100 dark:hover:bg-gray-700 text-gray-900 dark:text-gray-100"
                  >
                    +
                  </button>
                  <span className="ml-3 text-sm text-gray-500 dark:text-gray-400">
                    库存 {currentStock} 件
                  </span>
                </div>
              </div>

              {/* 购买按钮 */}
              <div className="flex gap-3 mb-6">
                <button
                  onClick={handleAddToCart}
                  className="flex-1 h-12 border-2 border-orange-500 text-orange-600 dark:text-orange-400 rounded-lg hover:bg-orange-50 dark:hover:bg-orange-900/20 flex items-center justify-center gap-2 font-medium transition-colors"
                >
                  <ShoppingCart className="w-5 h-5" />
                  加入购物车
                </button>
                <button
                  onClick={handleBuyNow}
                  className="flex-1 h-12 bg-gradient-to-r from-orange-500 to-red-500 text-white rounded-lg hover:from-orange-600 hover:to-red-600 font-medium transition-all"
                >
                  立即购买
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
                  {(() => {
                    const shop = getShopByName(product.shopName || '');
                    return (
                      <>
                        <Link 
                          to={`/shop/${shop.id}`}
                          className="flex items-center gap-3 hover:bg-gray-50 dark:hover:bg-gray-700/50 p-2 rounded-lg transition-colors -ml-2"
                        >
                          <div className="w-12 h-12 bg-gradient-to-br from-orange-400 to-red-400 rounded-lg flex items-center justify-center overflow-hidden">
                            <img 
                              src={shop.avatar} 
                              alt={product.shopName}
                              className="w-full h-full object-cover"
                              onError={(e) => {
                                const target = e.target as HTMLImageElement;
                                target.style.display = 'none';
                              }}
                            />
                          </div>
                          <div>
                            <div className="font-medium text-gray-900 dark:text-gray-100 group-hover:text-primary">
                              {product.shopName}
                            </div>
                            <div className="flex items-center gap-2 text-xs text-gray-500 dark:text-gray-400 mt-1">
                              <div className="flex items-center gap-1">
                                <Star className="w-3 h-3 text-yellow-400 fill-current" />
                                <span>{shop.rating}</span>
                              </div>
                              <span>|</span>
                              <span>粉丝 {formatNumber(shop.fans)}</span>
                            </div>
                          </div>
                          <ChevronRight className="w-5 h-5 text-gray-400" />
                        </Link>
                        <div className="flex gap-2">
                          <Link 
                            to={`/shop/${shop.id}`}
                            className="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded text-sm hover:bg-gray-50 dark:hover:bg-gray-700 text-gray-900 dark:text-gray-100"
                          >
                            进店逛逛
                          </Link>
                          <button className="px-4 py-2 border border-orange-500 text-orange-600 dark:text-orange-400 rounded text-sm hover:bg-orange-50 dark:hover:bg-orange-900/20">
                            关注店铺
                          </button>
                        </div>
                      </>
                    );
                  })()}
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
                  {mockCategories.find(c => c.id === product.categoryId)?.name || '其他'}
                </span>
              </div>
              <div className="bg-gray-50 dark:bg-gray-700 p-3 rounded">
                <span className="text-gray-600 dark:text-gray-400">上市时间：</span>
                <span className="text-gray-900 dark:text-gray-100">2024年</span>
              </div>
            </div>
          </div>

          {/* 商品图片展示 */}
          <div>
            <h3 className="font-medium text-gray-900 dark:text-gray-100 mb-4">商品展示</h3>
            <div className="space-y-4">
              {product.images.map((img, index) => (
                <div key={index} className="w-full bg-gray-100 dark:bg-gray-700 rounded-lg overflow-hidden">
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

          {/* 用户评价 */}
          <div className="mt-8">
            {/* 评价头部 */}
            <div className="flex items-center justify-between mb-6">
              <div className="flex items-center gap-3">
                <h3 className="text-xl font-semibold text-gray-900 dark:text-gray-100">用户评价</h3>
                <span className="text-sm text-gray-500 dark:text-gray-400">({getProductReviews(product?.id || '').length})</span>
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
            <div className="space-y-4">
              {getProductReviews(product?.id || '').map((review) => (
                <div
                  key={review.id}
                  className="glass-card rounded-xl p-6 hover:shadow-xl transition-all duration-300"
                >
                  {/* 用户信息 */}
                  <div className="flex items-start justify-between mb-4">
                    <div className="flex items-center gap-3">
                      <img
                        src={review.avatar}
                        alt="用户头像"
                        className="w-10 h-10 rounded-full object-cover"
                        onError={(e) => {
                          const target = e.target as HTMLImageElement;
                          target.src = `https://ui-avatars.com/api/?name=${review.username}&background=f97316&color=fff`;
                        }}
                      />
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
                      {review.images.map((img, idx) => (
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
                    <div className="flex items-center gap-4 opacity-0 group-hover:opacity-100 transition-opacity">
                      <button className="flex items-center gap-1.5 text-sm text-gray-500 dark:text-gray-400 hover:text-orange-500 dark:hover:text-orange-400 transition-colors">
                        <ThumbsUp className="w-4 h-4" />
                        <span>{review.likes}</span>
                      </button>
                      <button className="text-sm text-gray-500 dark:text-gray-400 hover:text-orange-500 dark:hover:text-orange-400 transition-colors">
                        回复
                      </button>
                    </div>
                    <div className="flex items-center gap-3">
                      <button className="flex items-center gap-1.5 text-sm text-gray-500 dark:text-gray-400 hover:text-orange-500 dark:hover:text-orange-400 transition-colors">
                        <ThumbsUp className="w-4 h-4" />
                        <span>{review.likes}</span>
                      </button>
                      <button className="text-sm text-gray-500 dark:text-gray-400 hover:text-orange-500 dark:hover:text-orange-400 transition-colors">
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
          </div>
        </div>
      </div>
    </div>
  );
};
