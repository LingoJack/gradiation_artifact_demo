import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useUserStore } from '../../store/useUserStore';

export const Login: React.FC = () => {
  const navigate = useNavigate();
  const { login } = useUserStore();
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = (e: React.FormEvent) => {
    e.preventDefault();
    // Mock 登录
    login(
      {
        id: '1',
        username,
        email: 'user@example.com',
        phone: '13800138000',
        createdAt: new Date().toISOString(),
      },
      'mock-token'
    );
    navigate('/');
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-orange-50 via-white to-red-50 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 px-4 relative overflow-hidden">
      {/* 动态背景装饰 - 液态玻璃风格 */}
      <div className="absolute inset-0 overflow-hidden">
        {/* 主要光晕 - 左上 */}
        <div 
          className="absolute -top-32 -left-32 w-[500px] h-[500px] rounded-full opacity-60"
          style={{
            background: 'radial-gradient(circle, rgba(251, 146, 60, 0.3) 0%, rgba(251, 146, 60, 0.1) 40%, transparent 70%)',
            animation: 'float 8s ease-in-out infinite',
          }}
        />
        {/* 主要光晕 - 右下 */}
        <div 
          className="absolute -bottom-32 -right-32 w-[600px] h-[600px] rounded-full opacity-60"
          style={{
            background: 'radial-gradient(circle, rgba(239, 68, 68, 0.25) 0%, rgba(239, 68, 68, 0.08) 40%, transparent 70%)',
            animation: 'float 10s ease-in-out infinite reverse',
          }}
        />
        {/* 辅助光晕 - 中上 */}
        <div 
          className="absolute top-20 left-1/2 -translate-x-1/2 w-[400px] h-[400px] rounded-full opacity-40"
          style={{
            background: 'radial-gradient(circle, rgba(251, 191, 36, 0.2) 0%, transparent 60%)',
            animation: 'float 12s ease-in-out infinite',
            animationDelay: '-2s',
          }}
        />
        {/* 液态玻璃光斑 */}
        <div 
          className="absolute top-1/4 right-1/4 w-[300px] h-[300px] rounded-full blur-3xl opacity-30"
          style={{
            background: 'linear-gradient(135deg, rgba(255, 255, 255, 0.8) 0%, rgba(251, 146, 60, 0.3) 50%, rgba(239, 68, 68, 0.2) 100%)',
            animation: 'morph 15s ease-in-out infinite',
          }}
        />
        <div 
          className="absolute bottom-1/4 left-1/4 w-[350px] h-[350px] rounded-full blur-3xl opacity-25"
          style={{
            background: 'linear-gradient(225deg, rgba(255, 255, 255, 0.6) 0%, rgba(239, 68, 68, 0.2) 50%, rgba(251, 146, 60, 0.2) 100%)',
            animation: 'morph 18s ease-in-out infinite reverse',
          }}
        />
      </div>
      
      <div className="glass-modal p-8 rounded-2xl w-full max-w-md relative z-10">
        {/* Logo */}
        <div className="text-center mb-8">
          <div className="inline-flex items-center justify-center w-16 h-16 rounded-full bg-gradient-to-r from-orange-500 to-red-500 mb-4">
            <svg className="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
            </svg>
          </div>
          <h2 className="text-3xl font-bold text-gray-900 dark:text-white">欢迎登录</h2>
          <p className="text-gray-500 dark:text-gray-400 mt-2">登录后享受更多优惠</p>
        </div>

        <form onSubmit={handleLogin}>
          <div className="mb-5">
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
              用户名
            </label>
            <input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="w-full px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent dark:bg-gray-700 dark:text-white transition-all"
              placeholder="请输入用户名"
              required
            />
          </div>
          <div className="mb-6">
            <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
              密码
            </label>
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent dark:bg-gray-700 dark:text-white transition-all"
              placeholder="请输入密码"
              required
            />
          </div>

          {/* 记住密码 & 忘记密码 */}
          <div className="flex items-center justify-between mb-6">
            <label className="flex items-center">
              <input type="checkbox" className="w-4 h-4 text-orange-500 border-gray-300 rounded focus:ring-orange-500" />
              <span className="ml-2 text-sm text-gray-600 dark:text-gray-400">记住密码</span>
            </label>
            <a href="#" className="text-sm text-orange-600 hover:text-orange-700 dark:text-orange-400 dark:hover:text-orange-300">
              忘记密码？
            </a>
          </div>

          <button
            type="submit"
            className="w-full bg-gradient-to-r from-orange-500 to-red-500 text-white py-3 rounded-lg hover:from-orange-600 hover:to-red-600 transition-all font-medium shadow-lg hover:shadow-xl"
          >
            登录
          </button>
        </form>

        {/* 第三方登录 */}
        <div className="mt-8">
          <div className="relative">
            <div className="absolute inset-0 flex items-center">
              <div className="w-full border-t border-gray-200 dark:border-gray-700"></div>
            </div>
            <div className="relative flex justify-center text-sm">
              <span className="px-4 bg-white dark:bg-gray-800 text-gray-500 dark:text-gray-400">或使用以下方式登录</span>
            </div>
          </div>

          <div className="grid grid-cols-3 gap-3 mt-6">
            <button className="flex items-center justify-center px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors">
              {/* WeChat icon from Bootstrap Icons */}
              <svg className="w-5 h-5 text-green-500" fill="currentColor" viewBox="0 0 16 16">
                <path d="M11.176 14.429c-2.665 0-4.826-1.8-4.826-4.018 0-2.22 2.159-4.02 4.824-4.02S16 8.191 16 10.411c0 1.21-.65 2.301-1.666 3.036a.32.32 0 0 0-.12.366l.218.81a.6.6 0 0 1 .029.117.166.166 0 0 1-.162.162.2.2 0 0 1-.092-.03l-1.057-.61a.5.5 0 0 0-.256-.074.5.5 0 0 0-.142.021 5.7 5.7 0 0 1-1.576.22M9.064 9.542a.647.647 0 1 0 .557-1 .645.645 0 0 0-.646.647.6.6 0 0 0 .09.353Zm3.232.001a.646.646 0 1 0 .546-1 .645.645 0 0 0-.644.644.63.63 0 0 0 .098.356"/>
                <path d="M0 6.826c0 1.455.781 2.765 2.001 3.656a.385.385 0 0 1 .143.439l-.161.6-.1.373a.5.5 0 0 0-.032.14.19.19 0 0 0 .193.193q.06 0 .111-.029l1.268-.733a.6.6 0 0 1 .308-.088q.088 0 .171.025a6.8 6.8 0 0 0 1.625.26 4.5 4.5 0 0 1-.177-1.251c0-2.936 2.785-5.02 5.824-5.02l.15.002C10.587 3.429 8.392 2 5.796 2 2.596 2 0 4.16 0 6.826m4.632-1.555a.77.77 0 1 1-1.54 0 .77.77 0 0 1 1.54 0m3.875 0a.77.77 0 1 1-1.54 0 .77.77 0 0 1 1.54 0"/>
              </svg>
            </button>
            <button className="flex items-center justify-center px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors">
              <svg className="w-5 h-5 text-blue-500" fill="currentColor" viewBox="0 0 24 24">
                <path d="M20.665 3.717l-17.73 6.837c-1.21.486-1.203 1.161-.222 1.462l4.552 1.42 10.532-6.645c.498-.303.953-.14.579.192l-8.533 7.701h-.002l.002.001-.314 4.692c.46 0 .663-.211.921-.46l2.211-2.15 4.599 3.397c.848.467 1.457.227 1.668-.785l3.019-14.228c.309-1.239-.473-1.8-1.282-1.434z"/>
              </svg>
            </button>
            <button className="flex items-center justify-center px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors">
              <svg className="w-5 h-5 text-gray-700 dark:text-gray-300" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12.017 0C5.396 0 .029 5.367.029 11.987c0 5.079 3.158 9.417 7.618 11.162-.105-.949-.199-2.403.041-3.439.219-.937 1.406-5.957 1.406-5.957s-.359-.72-.359-1.781c0-1.663.967-2.911 2.168-2.911 1.024 0 1.518.769 1.518 1.688 0 1.029-.653 2.567-.992 3.992-.285 1.193.6 2.165 1.775 2.165 2.128 0 3.768-2.245 3.768-5.487 0-2.861-2.063-4.869-5.008-4.869-3.41 0-5.409 2.562-5.409 5.199 0 1.033.394 2.143.889 2.741.099.12.112.225.085.345-.09.375-.293 1.199-.334 1.363-.053.225-.172.271-.401.165-1.495-.69-2.433-2.878-2.433-4.646 0-3.776 2.748-7.252 7.92-7.252 4.158 0 7.392 2.967 7.392 6.923 0 4.135-2.607 7.462-6.233 7.462-1.214 0-2.354-.629-2.758-1.379l-.749 2.848c-.269 1.045-1.004 2.352-1.498 3.146 1.123.345 2.306.535 3.55.535 6.607 0 11.985-5.365 11.985-11.987C23.97 5.39 18.592.026 11.985.026L12.017 0z"/>
              </svg>
            </button>
          </div>
        </div>

        <div className="mt-6 text-center text-sm text-gray-600 dark:text-gray-400">
          还没有账号？{' '}
          <Link to="/register" className="text-orange-600 hover:text-orange-700 dark:text-orange-400 dark:hover:text-orange-300 font-medium hover:underline">
            立即注册
          </Link>
        </div>
      </div>
    </div>
  );
};
