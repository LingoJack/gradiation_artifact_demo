import React from 'react';
import { Link } from 'react-router-dom';
import { User, MapPin, Heart, Package, Settings, LogOut, Ticket } from 'lucide-react';
import { useUserStore } from '../../store/useUserStore';
import { useCouponStore } from '../../store/useCouponStore';
import { useSpotlight } from '../../hooks/useSpotlight';

export const UserCenter: React.FC = () => {
  const { user, logout } = useUserStore();
  const { coupons, points, redPacket } = useCouponStore();
  const menuSpotlight = useSpotlight();
  const mainSpotlight = useSpotlight();

  const availableCoupons = coupons.filter((c) => c.status === 'available').length;

  const menuItems = [
    { icon: Package, label: '我的订单', path: '/orders' },
    { icon: Ticket, label: '我的优惠', path: '/coupons' },
    { icon: Heart, label: '我的收藏', path: '/favorites' },
    { icon: MapPin, label: '收货地址', path: '/addresses' },
    { icon: User, label: '个人信息', path: '/profile' },
    { icon: Settings, label: '账户设置', path: '/settings' },
  ];

  const handleLogout = () => {
    if (confirm('确定要退出登录吗？')) {
      logout();
      window.location.href = '/';
    }
  };

  return (
    <div className="container py-8">
      <div className="grid grid-cols-4 gap-8">
        {/* 左侧菜单 */}
        <div className="col-span-1">
          <div 
            ref={menuSpotlight.ref as React.RefObject<HTMLDivElement>}
            className="glass-user-card rounded-xl p-6 overflow-hidden relative"
            style={menuSpotlight.spotlightStyle}
            {...menuSpotlight.handlers}
          >
            {/* 用户信息 */}
            <div className="flex items-center space-x-4 mb-6 pb-6 border-b border-gray-200 dark:border-gray-700">
              <div className="w-16 h-16 bg-primary rounded-full flex items-center justify-center text-white text-2xl">
                {user?.username?.charAt(0).toUpperCase() || 'U'}
              </div>
              <div>
                <h2 className="font-bold dark:text-white">{user?.username || '用户'}</h2>
                <p className="text-sm text-gray-500 dark:text-gray-400">{user?.email}</p>
              </div>
            </div>

            {/* 优惠资产卡片 */}
            <div className="relative mb-6 pb-6 border-b border-gray-200 dark:border-gray-700">
              {/* 背景装饰 */}
              <div className="absolute -inset-1 bg-gradient-to-r from-orange-400 via-red-400 to-amber-400 rounded-xl opacity-15 blur-sm" />
              
              <div className="relative bg-gradient-to-br from-orange-500 via-red-500 to-amber-500 rounded-xl p-4 text-white overflow-hidden">
                {/* 装饰圆圈 */}
                <div className="absolute -right-4 -top-4 w-20 h-20 bg-white/10 rounded-full" />
                <div className="absolute -right-2 -bottom-2 w-12 h-12 bg-white/10 rounded-full" />
                
                <div className="relative flex items-center justify-between">
                  <div>
                    <p className="text-white/80 text-xs mb-1">我的优惠资产</p>
                    <p className="text-2xl font-bold">¥{(availableCoupons * 30 + redPacket).toFixed(0)}<span className="text-sm font-normal ml-1">可用</span></p>
                  </div>
                  <Link 
                    to="/coupons"
                    className="flex items-center gap-2 bg-white/20 hover:bg-white/30 backdrop-blur-sm px-4 py-2 rounded-full transition-all text-sm font-medium"
                  >
                    <span>查看全部</span>
                    <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5l7 7-7 7" />
                    </svg>
                  </Link>
                </div>
                
                <div className="relative grid grid-cols-3 gap-2 mt-4">
                  <div className="bg-white/15 backdrop-blur-sm rounded-lg py-2 text-center">
                    <div className="flex items-center justify-center gap-1">
                      <svg className="w-4 h-4 text-yellow-300" fill="currentColor" viewBox="0 0 20 20">
                        <path d="M10 2a8 8 0 100 16 8 8 0 000-16zm1 11H9v-2h2v2zm0-4H9V5h2v4z"/>
                      </svg>
                      <span className="text-lg font-bold">{availableCoupons}</span>
                    </div>
                    <p className="text-xs text-white/70">优惠券</p>
                  </div>
                  <div className="bg-white/15 backdrop-blur-sm rounded-lg py-2 text-center">
                    <div className="flex items-center justify-center gap-1">
                      <svg className="w-4 h-4 text-red-200" fill="currentColor" viewBox="0 0 20 20">
                        <path fillRule="evenodd" d="M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 17.657l-6.828-6.829a4 4 0 010-5.656z" clipRule="evenodd"/>
                      </svg>
                      <span className="text-lg font-bold">¥{redPacket}</span>
                    </div>
                    <p className="text-xs text-white/70">红包</p>
                  </div>
                  <div className="bg-white/15 backdrop-blur-sm rounded-lg py-2 text-center">
                    <div className="flex items-center justify-center gap-1">
                      <svg className="w-4 h-4 text-yellow-200" fill="currentColor" viewBox="0 0 20 20">
                        <path d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.736 6.979C9.208 6.193 9.696 6 10 6c.304 0 .792.193 1.264.979a1 1 0 001.715-1.029C12.279 4.784 11.232 4 10 4s-2.279.784-2.979 1.95c-.285.475-.507 1-.67 1.55H6a1 1 0 000 2h.013a9.358 9.358 0 000 1H6a1 1 0 100 2h.351c.163.55.385 1.075.67 1.55C7.721 15.216 8.768 16 10 16s2.279-.784 2.979-1.95a1 1 0 10-1.715-1.029c-.472.786-.96.979-1.264.979-.304 0-.792-.193-1.264-.979a4.265 4.265 0 01-.264-.521H10a1 1 0 100-2H8.017a7.36 7.36 0 010-1H10a1 1 0 100-2H8.472a4.265 4.265 0 01.264-.521z"/>
                      </svg>
                      <span className="text-lg font-bold">{points}</span>
                    </div>
                    <p className="text-xs text-white/70">积分</p>
                  </div>
                </div>
              </div>
            </div>

            {/* 菜单列表 */}
            <nav className="space-y-2">
              {menuItems.map((item) => (
                <Link
                  key={item.path}
                  to={item.path}
                  className="flex items-center space-x-3 px-4 py-3 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-700/50 transition dark:text-gray-300"
                >
                  <item.icon className="w-5 h-5 text-gray-400 dark:text-gray-500" />
                  <span>{item.label}</span>
                </Link>
              ))}
              <button
                onClick={handleLogout}
                className="w-full flex items-center space-x-3 px-4 py-3 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-700/50 transition text-error"
              >
                <LogOut className="w-5 h-5" />
                <span>退出登录</span>
              </button>
            </nav>
          </div>
        </div>

        {/* 右侧内容 */}
        <div className="col-span-3">
          <div 
            ref={mainSpotlight.ref as React.RefObject<HTMLDivElement>}
            className="glass-liquid rounded-xl p-6 overflow-hidden relative"
            style={mainSpotlight.spotlightStyle}
            {...mainSpotlight.handlers}
          >
            <h2 className="text-xl font-bold mb-6 dark:text-white">个人信息</h2>
            <div className="space-y-4">
              <div className="flex items-center space-x-4">
                <div className="w-24 text-gray-600 dark:text-gray-400">用户名：</div>
                <div className="dark:text-white">{user?.username}</div>
              </div>
              <div className="flex items-center space-x-4">
                <div className="w-24 text-gray-600 dark:text-gray-400">邮箱：</div>
                <div className="dark:text-white">{user?.email}</div>
              </div>
              <div className="flex items-center space-x-4">
                <div className="w-24 text-gray-600 dark:text-gray-400">手机号：</div>
                <div className="dark:text-white">{user?.phone}</div>
              </div>
              <div className="flex items-center space-x-4">
                <div className="w-24 text-gray-600 dark:text-gray-400">注册时间：</div>
                <div className="dark:text-white">{user?.createdAt}</div>
              </div>
            </div>

            <div className="mt-8">
              <h3 className="text-lg font-bold mb-4 dark:text-white">快捷入口</h3>
              <div className="grid grid-cols-4 gap-4">
                <Link
                  to="/orders"
                  className="bg-primary-light/50 dark:bg-primary/20 rounded-lg p-6 text-center hover:shadow-md transition dark:text-white"
                >
                  <Package className="w-8 h-8 text-primary mx-auto mb-2" />
                  <p className="text-sm">我的订单</p>
                </Link>
                <Link
                  to="/coupons"
                  className="bg-amber-50 dark:bg-amber-900/20 rounded-lg p-6 text-center hover:shadow-md transition dark:text-white"
                >
                  <Ticket className="w-8 h-8 text-amber-500 mx-auto mb-2" />
                  <p className="text-sm">我的优惠</p>
                </Link>
                <Link
                  to="/favorites"
                  className="bg-red-50 dark:bg-red-900/20 rounded-lg p-6 text-center hover:shadow-md transition dark:text-white"
                >
                  <Heart className="w-8 h-8 text-error mx-auto mb-2" />
                  <p className="text-sm">我的收藏</p>
                </Link>
                <Link
                  to="/addresses"
                  className="bg-blue-50 dark:bg-blue-900/20 rounded-lg p-6 text-center hover:shadow-md transition dark:text-white"
                >
                  <MapPin className="w-8 h-8 text-info mx-auto mb-2" />
                  <p className="text-sm">收货地址</p>
                </Link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
