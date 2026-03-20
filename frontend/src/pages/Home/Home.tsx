import React from 'react';
import { Link } from 'react-router-dom';
import { ProductCard } from '../../components/ProductCard/ProductCard';
import { mockCategories, mockProducts } from '../../utils/mockData';
import type { Product } from '../../types/product';

export const Home: React.FC = () => {
  return (
    <div>
      {/* 轮播图区域 - 重新设计 */}
      <div className="bg-white dark:bg-gray-900">
        <div className="container py-5">
          <div className="flex gap-3" style={{ height: '420px' }}>
            {/* 左侧分类 - 浅色清爽设计 */}
            <div className="w-52 bg-gray-50 dark:bg-gray-800 rounded-xl overflow-hidden flex-shrink-0 border border-gray-200 dark:border-gray-700">
              <div className="p-3 border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800">
                <h3 className="text-gray-800 dark:text-gray-200 font-bold text-sm">
                  全部商品分类
                </h3>
              </div>
              <ul className="py-2">
                {mockCategories.map((cat, index) => {
                  const colors = [
                    'bg-orange-100 text-orange-600 dark:bg-orange-900/30 dark:text-orange-400',
                    'bg-blue-100 text-blue-600 dark:bg-blue-900/30 dark:text-blue-400',
                    'bg-pink-100 text-pink-600 dark:bg-pink-900/30 dark:text-pink-400',
                    'bg-teal-100 text-teal-600 dark:bg-teal-900/30 dark:text-teal-400',
                    'bg-green-100 text-green-600 dark:bg-green-900/30 dark:text-green-400',
                    'bg-purple-100 text-purple-600 dark:bg-purple-900/30 dark:text-purple-400',
                    'bg-red-100 text-red-600 dark:bg-red-900/30 dark:text-red-400',
                    'bg-yellow-100 text-yellow-700 dark:bg-yellow-900/30 dark:text-yellow-400',
                  ];
                  const colorClass = colors[index % colors.length];
                  
                  return (
                    <li
                      key={cat.id}
                      className="flex items-center px-3 py-2.5 hover:bg-white dark:hover:bg-gray-700 cursor-pointer transition-all group relative"
                    >
                      {/* 彩色图标方块 */}
                      <span className={`w-8 h-8 ${colorClass} rounded-lg flex items-center justify-center mr-3 transition-all group-hover:rounded-xl`}>
                        <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16M4 18h16" />
                        </svg>
                      </span>
                      {/* 分类名称 */}
                      <span className="text-sm text-gray-700 dark:text-gray-300 group-hover:text-orange-600 transition-colors font-medium">{cat.name}</span>
                      {/* 箭头 */}
                      <svg className="w-4 h-4 ml-auto text-gray-400 opacity-0 group-hover:opacity-100 transition-all group-hover:translate-x-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5l7 7-7 7" />
                      </svg>
                      {/* 悬停指示条 */}
                      <div className="absolute left-0 top-1/2 -translate-y-1/2 w-1 h-6 bg-orange-500 rounded-r opacity-0 group-hover:opacity-100 transition-opacity" />
                    </li>
                  );
                })}
              </ul>
            </div>

            {/* 中间轮播图 - 大图设计 */}
            <div className="flex-1 rounded-xl overflow-hidden shadow-lg relative group">
              <img 
                src="https://images.unsplash.com/photo-1607082348824-0a96f2a4b9da?w=1200&h=420&fit=crop" 
                alt="双11大促" 
                className="w-full h-full object-cover"
              />
              <div className="absolute inset-0 bg-gradient-to-r from-orange-600/80 via-red-500/70 to-transparent">
                <div className="relative h-full flex items-center px-12">
                  <div className="text-white">
                    <h2 className="text-6xl font-bold mb-3 drop-shadow-lg">双11大促</h2>
                    <p className="text-xl mb-6 drop-shadow opacity-95">全场商品5折起 | 限时特惠</p>
                    <button className="bg-white text-orange-600 px-8 py-3 rounded-full font-bold text-lg hover:bg-opacity-90 transition shadow-xl">
                      立即抢购
                    </button>
                  </div>
                </div>
              </div>
              {/* 轮播指示器 */}
              <div className="absolute bottom-4 left-1/2 transform -translate-x-1/2 flex space-x-2">
                <div className="w-8 h-1.5 bg-white rounded-full"></div>
                <div className="w-1.5 h-1.5 bg-white/50 rounded-full"></div>
                <div className="w-1.5 h-1.5 bg-white/50 rounded-full"></div>
                <div className="w-1.5 h-1.5 bg-white/50 rounded-full"></div>
              </div>
            </div>

            {/* 右侧推荐 - 卡片式设计 */}
            <div className="w-72 flex-shrink-0 flex flex-col gap-3">
              {/* 用户信息卡片 */}
              <div className="bg-gradient-to-br from-orange-50 to-red-50 dark:from-gray-800 dark:to-gray-800 rounded-xl p-5 border border-orange-100 dark:border-gray-700">
                <div className="flex items-center mb-4">
                  <div className="w-12 h-12 bg-orange-500 rounded-full flex items-center justify-center mr-3 overflow-hidden">
                    <img 
                      src="https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?w=100&h=100&fit=crop"
                      alt="用户头像"
                      className="w-full h-full object-cover"
                      onError={(e) => {
                        (e.target as HTMLImageElement).style.display = 'none';
                        const fallback = document.createElement('span');
                        fallback.className = 'text-white text-xl';
                        fallback.textContent = 'U';
                        (e.target as HTMLImageElement).parentNode?.appendChild(fallback);
                      }}
                    />
                  </div>
                  <div>
                    <p className="font-bold text-gray-800 dark:text-gray-200">Hi，欢迎来到淘宝</p>
                    <p className="text-xs text-gray-500 dark:text-gray-400">登录享更多优惠</p>
                  </div>
                </div>
                <div className="grid grid-cols-3 gap-2 text-center">
                  <div className="bg-white dark:bg-gray-700 rounded-lg py-2">
                    <p className="text-orange-500 font-bold">优惠券</p>
                    <p className="text-xs text-gray-400 dark:text-gray-500 mt-1">3张</p>
                  </div>
                  <div className="bg-white dark:bg-gray-700 rounded-lg py-2">
                    <p className="text-orange-500 font-bold">红包</p>
                    <p className="text-xs text-gray-400 dark:text-gray-500 mt-1">¥20</p>
                  </div>
                  <div className="bg-white dark:bg-gray-700 rounded-lg py-2">
                    <p className="text-orange-500 font-bold">积分</p>
                    <p className="text-xs text-gray-400 dark:text-gray-500 mt-1">1280</p>
                  </div>
                </div>
              </div>

              {/* 活动卡片 */}
              <div className="flex-1 rounded-xl overflow-hidden relative cursor-pointer hover:shadow-xl transition-shadow group">
                <img 
                  src="https://images.unsplash.com/photo-1441986300917-64674bd600d8?w=400&h=300&fit=crop" 
                  alt="新人专享" 
                  className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
                />
                 <div className="absolute inset-0 bg-gradient-to-br from-purple-600/85 to-pink-600/85 flex items-center justify-center">
                   <div className="text-white text-center">
                     <p className="font-bold text-2xl mb-2">新人专享</p>
                     <p className="text-base opacity-90">首单立减50元</p>
                   </div>
                 </div>
               </div>

               {/* 公告 */}
               <div className="bg-white dark:bg-gray-800 rounded-xl p-4 border border-gray-100 dark:border-gray-700">
                 <div className="flex items-center mb-2">
                   <svg className="w-5 h-5 text-orange-500 mr-2" fill="currentColor" viewBox="0 0 20 20">
                     <path fillRule="evenodd" d="M18 3a1 1 0 00-1.447-.894L8.763 6H5a3 3 0 000 6h.28l1.771 5.316A1 1 0 008 18h1a1 1 0 001-1v-4.382l6.553 3.276A1 1 0 0018 15V3z" clipRule="evenodd" />
                   </svg>
                   <span className="font-bold text-sm text-gray-700 dark:text-gray-200">淘宝头条</span>
                 </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* 快捷入口 - 横向滚动设计 */}
      <div className="bg-white dark:bg-gray-800 border-t border-gray-100 dark:border-gray-700">
        <div className="container py-6">
          <div className="grid grid-cols-10 gap-4">
            {[
              { icon: 'smartphone', name: '手机', color: 'text-blue-500', bgColor: 'bg-blue-100 dark:bg-blue-900/30' },
              { icon: 'shirt', name: '男装', color: 'text-purple-500', bgColor: 'bg-purple-100 dark:bg-purple-900/30' },
              { icon: 'sparkles', name: '女装', color: 'text-pink-500', bgColor: 'bg-pink-100 dark:bg-pink-900/30' },
              { icon: 'lightning-bolt', name: '运动', color: 'text-green-500', bgColor: 'bg-green-100 dark:bg-green-900/30' },
              { icon: 'computer-desktop', name: '电脑', color: 'text-indigo-500', bgColor: 'bg-indigo-100 dark:bg-indigo-900/30' },
              { icon: 'home', name: '家居', color: 'text-orange-500', bgColor: 'bg-orange-100 dark:bg-orange-900/30' },
              { icon: 'heart', name: '生鲜', color: 'text-red-500', bgColor: 'bg-red-100 dark:bg-red-900/30' },
              { icon: 'beaker', name: '美妆', color: 'text-rose-500', bgColor: 'bg-rose-100 dark:bg-rose-900/30' },
              { icon: 'face-smile', name: '母婴', color: 'text-yellow-500', bgColor: 'bg-yellow-100 dark:bg-yellow-900/30' },
              { icon: 'book-open', name: '图书', color: 'text-teal-500', bgColor: 'bg-teal-100 dark:bg-teal-900/30' },
            ].map((item, idx) => (
              <div key={idx} className="text-center cursor-pointer group">
                <div className={`w-14 h-14 ${item.bgColor} rounded-2xl flex items-center justify-center mx-auto mb-2 group-hover:scale-110 transition-transform shadow-sm`}>
                  <svg className={`w-7 h-7 ${item.color}`} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    {item.icon === 'smartphone' && <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z" />}
                    {item.icon === 'shirt' && <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6.5 6.5L8 5l4-1 4 1 1.5 1.5m-11 0l-2 4v1l2 1v5.5a2 2 0 002 2h7a2 2 0 002-2v-5.5l2-1v-1l-2-4m-11 0l3.5 2.5h5l3.5-2.5" />}
                    {item.icon === 'sparkles' && <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z" />}
                    {item.icon === 'lightning-bolt' && <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />}
                    {item.icon === 'computer-desktop' && <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />}
                    {item.icon === 'home' && <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />}
                    {item.icon === 'heart' && <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />}
                    {item.icon === 'beaker' && <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />}
                    {item.icon === 'face-smile' && <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15.182 15.182a4.5 4.5 0 01-6.364 0M21 12a9 9 0 11-18 0 9 9 0 0118 0zM9.75 9.75c0 .414-.168.75-.375.75S9 10.164 9 9.75 9.168 9 9.375 9s.375.336.375.75zm-.375 0h.008v.015h-.008V9.75zm5.625 0c0 .414-.168.75-.375.75s-.375-.336-.375-.75.168-.75.375-.75.375.336.375.75zm-.375 0h.008v.015h-.008V9.75z" />}
                    {item.icon === 'book-open' && <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />}
                  </svg>
                </div>
                <p className={`text-xs font-medium ${item.color}`}>{item.name}</p>
              </div>
            ))}
          </div>
        </div>
      </div>

      {/* 热门推荐 */}
      <div className="container mt-8">
        <div className="flex items-center justify-between mb-6">
          <h2 className="text-2xl font-bold flex items-center text-gray-900 dark:text-gray-100">
            <div className="w-8 h-8 bg-gradient-to-br from-orange-500 to-red-500 rounded-lg flex items-center justify-center mr-3">
              <svg className="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 20 20">
                <path fillRule="evenodd" d="M12.395 2.553a1 1 0 00-1.45-.385c-.345.23-.614.558-.822.88-.214.33-.403.713-.57 1.116-.334.804-.614 1.768-.84 2.734a31.365 31.365 0 00-.613 3.58 2.64 2.64 0 01-.945-1.067c-.328-.68-.398-1.534-.398-2.654A1 1 0 005.05 6.05 6.981 6.981 0 003 11a7 7 0 1011.95-4.95c-.592-.591-.98-.985-1.348-1.467-.363-.476-.724-1.063-1.207-2.03zM12.12 15.12A3 3 0 017 13s.879.5 2.5.5c0-1 .5-4 1.25-4.5.5 1 .786 1.293 1.371 1.879A2.99 2.99 0 0113 13a2.99 2.99 0 01-.879 2.121z" clipRule="evenodd" />
              </svg>
            </div>
            热门推荐
          </h2>
          <Link to="/products" className="text-primary hover:underline text-sm font-medium">
            查看更多 →
          </Link>
        </div>
        <div className="grid grid-cols-4 gap-5">
          {mockProducts.slice(0, 4).map((product) => (
            <ProductCard key={product.id} product={product as Product} />
          ))}
        </div>
      </div>

      {/* 分类推荐 */}
      <div className="container mt-12 mb-12">
        <div className="flex items-center justify-between mb-6">
          <h2 className="text-2xl font-bold flex items-center text-gray-900 dark:text-gray-100">
            <div className="w-8 h-8 bg-gradient-to-br from-pink-500 to-purple-500 rounded-lg flex items-center justify-center mr-3">
              <svg className="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 20 20">
                <path fillRule="evenodd" d="M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 17.657l-6.828-6.829a4 4 0 010-5.656z" clipRule="evenodd" />
              </svg>
            </div>
            猜你喜欢
          </h2>
          <Link to="/products" className="text-primary hover:underline text-sm font-medium">
            查看更多 →
          </Link>
        </div>
        <div className="grid grid-cols-5 gap-5">
          {mockProducts.slice(4, 9).map((product) => (
            <ProductCard key={product.id} product={product as Product} />
          ))}
        </div>
      </div>
    </div>
  );
};
