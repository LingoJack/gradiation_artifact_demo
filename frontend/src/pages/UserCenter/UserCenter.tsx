import React from 'react';
import { Link } from 'react-router-dom';
import { User, MapPin, Heart, Package, Settings, LogOut } from 'lucide-react';
import { useUserStore } from '../../store/useUserStore';
import { useSpotlight } from '../../hooks/useSpotlight';

export const UserCenter: React.FC = () => {
  const { user, logout } = useUserStore();
  const menuSpotlight = useSpotlight();
  const mainSpotlight = useSpotlight();

  const menuItems = [
    { icon: Package, label: '我的订单', path: '/orders' },
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
                <Link
                  to="/settings"
                  className="bg-gray-50 dark:bg-gray-700/50 rounded-lg p-6 text-center hover:shadow-md transition dark:text-white"
                >
                  <Settings className="w-8 h-8 text-gray-600 dark:text-gray-400 mx-auto mb-2" />
                  <p className="text-sm">账户设置</p>
                </Link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
