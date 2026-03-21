import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { ShoppingBag, User, Search, Heart, Sun, Moon, Monitor } from 'lucide-react';
import { useUserStore } from '../../store/useUserStore';
import { useCartStore } from '../../store/useCartStore';
import { useTheme } from '../../hooks/useTheme';

export const Header: React.FC = () => {
  const { isAuthenticated, user } = useUserStore();
  const { items } = useCartStore();
  const selectedCount = items.filter((i) => i.selected).length;
  const [searchQuery, setSearchQuery] = useState('');
  const [showThemeMenu, setShowThemeMenu] = useState(false);
  const navigate = useNavigate();
  const { theme, setTheme, resolvedTheme } = useTheme();

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    if (searchQuery.trim()) {
      navigate(`/products?search=${encodeURIComponent(searchQuery.trim())}`);
    }
  };

  return (
    <header className="bg-gradient-to-r from-orange-500/75 to-red-500/75 backdrop-blur-2xl shadow-lg sticky top-0 z-50 border-b border-white/20">
      <div className="container">
        <div className="flex items-center justify-between h-16">
          {/* Logo */}
          <Link to="/" className="flex items-center group">
            <div className="bg-white rounded-lg px-4 py-1.5 shadow-sm">
              <span className="text-2xl font-bold bg-gradient-to-r from-orange-500 to-red-500 bg-clip-text text-transparent">
                淘宝
              </span>
            </div>
          </Link>

          {/* Search Bar */}
          <form onSubmit={handleSearch} className="flex-1 max-w-2xl mx-8">
            <div className="relative flex items-center">
              <div className="absolute left-4 text-gray-400">
                <Search className="w-5 h-5" />
              </div>
              <input
                type="text"
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                placeholder="搜索商品、店铺..."
                className="w-full h-10 pl-12 pr-24 bg-white border-0 rounded-full focus:outline-none text-gray-700 shadow-sm transition-all duration-300 focus:shadow-[0_0_0_2px_rgba(59,130,246,0.5),0_0_20px_rgba(139,92,246,0.3),0_0_40px_rgba(59,130,246,0.2)]"
              />
              <button
                type="submit"
                className="absolute right-1 top-1 h-8 px-6 bg-gradient-to-r from-orange-500 to-red-500 rounded-full hover:from-orange-600 hover:to-red-600 transition text-white font-medium text-sm"
              >
                搜索
              </button>
            </div>
          </form>

          {/* User Menu */}
          <div className="flex items-center space-x-5">
            {/* Theme Toggle */}
            <div className="relative">
              <button
                onClick={() => setShowThemeMenu(!showThemeMenu)}
                className="text-white hover:text-orange-100 transition p-1.5 rounded-lg hover:bg-white/10"
                title="切换主题"
              >
                {resolvedTheme === 'dark' ? (
                  <Moon className="w-5 h-5" />
                ) : (
                  <Sun className="w-5 h-5" />
                )}
              </button>
              {showThemeMenu && (
                <div className="absolute right-0 top-full mt-2 glass-dropdown rounded-xl py-1 min-w-[140px] z-50">
                  <button
                    onClick={() => {
                      setTheme('light');
                      setShowThemeMenu(false);
                    }}
                    className={`w-full px-4 py-2 text-left text-sm flex items-center space-x-2 hover:bg-gray-100 dark:hover:bg-gray-700 ${
                      theme === 'light' ? 'text-orange-500 font-medium' : 'text-gray-700 dark:text-gray-300'
                    }`}
                  >
                    <Sun className="w-4 h-4" />
                    <span>亮色模式</span>
                  </button>
                  <button
                    onClick={() => {
                      setTheme('dark');
                      setShowThemeMenu(false);
                    }}
                    className={`w-full px-4 py-2 text-left text-sm flex items-center space-x-2 hover:bg-gray-100 dark:hover:bg-gray-700 ${
                      theme === 'dark' ? 'text-orange-500 font-medium' : 'text-gray-700 dark:text-gray-300'
                    }`}
                  >
                    <Moon className="w-4 h-4" />
                    <span>暗色模式</span>
                  </button>
                  <button
                    onClick={() => {
                      setTheme('system');
                      setShowThemeMenu(false);
                    }}
                    className={`w-full px-4 py-2 text-left text-sm flex items-center space-x-2 hover:bg-gray-100 dark:hover:bg-gray-700 ${
                      theme === 'system' ? 'text-orange-500 font-medium' : 'text-gray-700 dark:text-gray-300'
                    }`}
                  >
                    <Monitor className="w-4 h-4" />
                    <span>跟随系统</span>
                  </button>
                </div>
              )}
            </div>
            
            {isAuthenticated ? (
              <>
                <Link to="/user" className="flex items-center space-x-2 text-white hover:text-orange-100 transition">
                  <User className="w-5 h-5" />
                  <span className="text-sm font-medium">Hi, {user?.username}</span>
                </Link>
                <Link to="/orders" className="text-white hover:text-orange-100 transition text-sm font-medium">
                  我的订单
                </Link>
                <Link to="/cart" className="relative text-white hover:text-orange-100 transition">
                  <ShoppingBag className="w-5 h-5" />
                  {selectedCount > 0 && (
                    <span className="absolute -top-1.5 -right-1.5 bg-yellow-400 text-orange-900 text-xs w-4.5 h-4.5 rounded-full flex items-center justify-center font-bold">
                      {selectedCount}
                    </span>
                  )}
                </Link>
                <Link to="/user" className="text-white hover:text-orange-100 transition">
                  <Heart className="w-5 h-5" />
                </Link>
              </>
            ) : (
              <>
                <Link to="/login" className="text-white hover:text-orange-100 transition text-sm font-medium">
                  登录
                </Link>
                <Link to="/register" className="bg-white text-orange-500 px-4 py-1.5 rounded-full text-sm font-medium hover:bg-orange-50 transition">
                  注册
                </Link>
              </>
            )}
          </div>
        </div>
      </div>
    </header>
  );
};
