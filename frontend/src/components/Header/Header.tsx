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
  const [isSearchFocused, setIsSearchFocused] = useState(false);
  const navigate = useNavigate();
  const { theme, setTheme, resolvedTheme } = useTheme();

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    if (searchQuery.trim()) {
      navigate(`/products?search=${encodeURIComponent(searchQuery.trim())}`);
    }
  };

  return (
    <header 
      className={`backdrop-blur-2xl shadow-lg sticky top-0 z-50 border-b transition-all duration-700 relative overflow-hidden ${
        isSearchFocused 
          ? 'bg-white/95 border-gray-200' 
          : 'bg-gradient-to-r from-orange-500/75 to-red-500/75 border-white/20'
      }`}
    >
      <div className="container">
        <div className="flex items-center justify-between h-16">
          {/* Logo */}
          <Link to="/" className="flex items-center group">
            <div className={`rounded-lg px-4 py-1.5 shadow-sm transition-all duration-700 ${
              isSearchFocused 
                ? 'bg-gradient-to-r from-orange-500 to-red-500' 
                : 'bg-white'
            }`}>
              <span className={`text-2xl font-bold transition-all duration-700 ${
                isSearchFocused 
                  ? 'text-white' 
                  : 'bg-gradient-to-r from-orange-500 to-red-500 bg-clip-text text-transparent'
              }`}>
                淘宝
              </span>
            </div>
          </Link>

          {/* Search Bar */}
          <form onSubmit={handleSearch} className="flex-1 max-w-2xl mx-8">
            <div className="relative flex items-center">
              {/* 流动光效边框 */}
              <div 
                className={`absolute inset-0 rounded-full transition-all duration-700 ${
                  isSearchFocused ? 'opacity-100 scale-[1.02]' : 'opacity-0 scale-100'
                }`}
                style={{
                  background: 'linear-gradient(90deg, #f97316, #ef4444, #fb923c, #f97316, #ef4444, #f97316)',
                  backgroundSize: '300% 100%',
                  animation: isSearchFocused ? 'gradient-flow 1.5s linear infinite' : 'none',
                  padding: '3px',
                }}
              >
                <div className="w-full h-full bg-white rounded-full" />
              </div>
              
              {/* 搜索图标 */}
              <div className="absolute left-4 text-gray-400 z-10">
                <Search className="w-5 h-5" />
              </div>
              
              {/* 输入框 */}
              <input
                type="text"
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                onFocus={() => setIsSearchFocused(true)}
                onBlur={() => setIsSearchFocused(false)}
                placeholder="搜索商品、店铺..."
                className="relative w-full h-10 pl-12 pr-24 bg-white border-0 rounded-full focus:outline-none text-gray-700 shadow-sm transition-all duration-300"
                style={{
                  boxShadow: isSearchFocused 
                    ? '0 0 30px rgba(249, 115, 22, 0.5), 0 0 60px rgba(249, 115, 22, 0.3)' 
                    : '0 1px 3px rgba(0, 0, 0, 0.1)',
                }}
              />
              
              {/* 搜索按钮 */}
              <button
                type="submit"
                className="absolute right-1 top-1 h-8 px-6 bg-gradient-to-r from-orange-500 to-red-500 rounded-full hover:from-orange-600 hover:to-red-600 transition text-white font-medium text-sm z-10"
              >
                搜索
              </button>
            </div>
          </form>

          {/* 全局动画样式 */}
          <style>{`
            @keyframes gradient-flow {
              0% {
                background-position: 0% 50%;
              }
              100% {
                background-position: 300% 50%;
              }
            }
          `}</style>

          {/* User Menu */}
          <div className="flex items-center space-x-5">
            {/* Theme Toggle */}
            <div className="relative">
              <button
                onClick={() => setShowThemeMenu(!showThemeMenu)}
                className={`transition p-1.5 rounded-lg ${
                  isSearchFocused 
                    ? 'text-gray-600 hover:text-orange-500 hover:bg-gray-100' 
                    : 'text-white hover:text-orange-100 hover:bg-white/10'
                }`}
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
                <Link to="/user" className={`flex items-center space-x-2 transition ${
                  isSearchFocused 
                    ? 'text-gray-700 hover:text-orange-500' 
                    : 'text-white hover:text-orange-100'
                }`}>
                  <User className="w-5 h-5" />
                  <span className="text-sm font-medium">Hi, {user?.username}</span>
                </Link>
                <Link to="/orders" className={`transition text-sm font-medium ${
                  isSearchFocused 
                    ? 'text-gray-700 hover:text-orange-500' 
                    : 'text-white hover:text-orange-100'
                }`}>
                  我的订单
                </Link>
                <Link to="/cart" className={`relative transition ${
                  isSearchFocused 
                    ? 'text-gray-700 hover:text-orange-500' 
                    : 'text-white hover:text-orange-100'
                }`}>
                  <ShoppingBag className="w-5 h-5" />
                  {selectedCount > 0 && (
                    <span className="absolute -top-1.5 -right-1.5 bg-yellow-400 text-orange-900 text-xs w-4.5 h-4.5 rounded-full flex items-center justify-center font-bold">
                      {selectedCount}
                    </span>
                  )}
                </Link>
                <Link to="/user" className={`transition ${
                  isSearchFocused 
                    ? 'text-gray-700 hover:text-orange-500' 
                    : 'text-white hover:text-orange-100'
                }`}>
                  <Heart className="w-5 h-5" />
                </Link>
              </>
            ) : (
              <>
                <Link to="/login" className={`transition text-sm font-medium ${
                  isSearchFocused 
                    ? 'text-gray-700 hover:text-orange-500' 
                    : 'text-white hover:text-orange-100'
                }`}>
                  登录
                </Link>
                <Link to="/register" className={`px-4 py-1.5 rounded-full text-sm font-medium transition ${
                  isSearchFocused 
                    ? 'bg-gradient-to-r from-orange-500 to-red-500 text-white hover:from-orange-600 hover:to-red-600' 
                    : 'bg-white text-orange-500 hover:bg-orange-50'
                }`}>
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
