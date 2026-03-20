import React from 'react';
import { Link } from 'react-router-dom';
import { User, MapPin, Heart, Package, Settings, LogOut } from 'lucide-react';
import { useUserStore } from '../../store/useUserStore';

export const UserCenter: React.FC = () => {
  const { user, logout } = useUserStore();

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
          <div className="bg-white rounded-lg p-6">
            {/* 用户信息 */}
            <div className="flex items-center space-x-4 mb-6 pb-6 border-b">
              <div className="w-16 h-16 bg-primary rounded-full flex items-center justify-center text-white text-2xl">
                {user?.username?.charAt(0).toUpperCase() || 'U'}
              </div>
              <div>
                <h2 className="font-bold">{user?.username || '用户'}</h2>
                <p className="text-sm text-gray-500">{user?.email}</p>
              </div>
            </div>

            {/* 菜单列表 */}
            <nav className="space-y-2">
              {menuItems.map((item) => (
                <Link
                  key={item.path}
                  to={item.path}
                  className="flex items-center space-x-3 px-4 py-3 rounded-lg hover:bg-gray-50 transition"
                >
                  <item.icon className="w-5 h-5 text-gray-400" />
                  <span>{item.label}</span>
                </Link>
              ))}
              <button
                onClick={handleLogout}
                className="w-full flex items-center space-x-3 px-4 py-3 rounded-lg hover:bg-gray-50 transition text-error"
              >
                <LogOut className="w-5 h-5" />
                <span>退出登录</span>
              </button>
            </nav>
          </div>
        </div>

        {/* 右侧内容 */}
        <div className="col-span-3">
          <div className="bg-white rounded-lg p-6">
            <h2 className="text-xl font-bold mb-6">个人信息</h2>
            <div className="space-y-4">
              <div className="flex items-center space-x-4">
                <div className="w-24 text-gray-600">用户名：</div>
                <div>{user?.username}</div>
              </div>
              <div className="flex items-center space-x-4">
                <div className="w-24 text-gray-600">邮箱：</div>
                <div>{user?.email}</div>
              </div>
              <div className="flex items-center space-x-4">
                <div className="w-24 text-gray-600">手机号：</div>
                <div>{user?.phone}</div>
              </div>
              <div className="flex items-center space-x-4">
                <div className="w-24 text-gray-600">注册时间：</div>
                <div>{user?.createdAt}</div>
              </div>
            </div>

            <div className="mt-8">
              <h3 className="text-lg font-bold mb-4">快捷入口</h3>
              <div className="grid grid-cols-4 gap-4">
                <Link
                  to="/orders"
                  className="bg-primary-light rounded-lg p-6 text-center hover:shadow-md transition"
                >
                  <Package className="w-8 h-8 text-primary mx-auto mb-2" />
                  <p className="text-sm">我的订单</p>
                </Link>
                <Link
                  to="/favorites"
                  className="bg-red-50 rounded-lg p-6 text-center hover:shadow-md transition"
                >
                  <Heart className="w-8 h-8 text-error mx-auto mb-2" />
                  <p className="text-sm">我的收藏</p>
                </Link>
                <Link
                  to="/addresses"
                  className="bg-blue-50 rounded-lg p-6 text-center hover:shadow-md transition"
                >
                  <MapPin className="w-8 h-8 text-info mx-auto mb-2" />
                  <p className="text-sm">收货地址</p>
                </Link>
                <Link
                  to="/settings"
                  className="bg-gray-50 rounded-lg p-6 text-center hover:shadow-md transition"
                >
                  <Settings className="w-8 h-8 text-gray-600 mx-auto mb-2" />
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
