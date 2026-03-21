import React, { useState } from 'react';
import { Camera, User, Mail, Phone, Calendar, Edit3, Check, Crown, Star } from 'lucide-react';
import { useUserStore } from '../../store/useUserStore';
import { useSpotlight } from '../../hooks/useSpotlight';

// 预设头像列表
const avatarOptions = [
  'https://images.unsplash.com/photo-1494790108377-be9c29b29330?w=200&h=200&fit=crop&crop=face',
  'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=200&h=200&fit=crop&crop=face',
  'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=200&h=200&fit=crop&crop=face',
  'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=200&h=200&fit=crop&crop=face',
  'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?w=200&h=200&fit=crop&crop=face',
  'https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=200&h=200&fit=crop&crop=face',
  'https://images.unsplash.com/photo-1544005313-94ddf0286df2?w=200&h=200&fit=crop&crop=face',
  'https://images.unsplash.com/photo-1517841905240-472988babdf9?w=200&h=200&fit=crop&crop=face',
];

export const Profile: React.FC = () => {
  const { user, updateProfile } = useUserStore();
  const cardSpotlight = useSpotlight();

  const [formData, setFormData] = useState({
    username: user?.username || '',
    email: user?.email || '',
    phone: user?.phone || '',
    nickname: user?.username || '',
    gender: user?.gender || 'male',
    birthday: user?.birthday || '',
    bio: user?.bio || '',
  });

  const [selectedAvatar, setSelectedAvatar] = useState(user?.avatar || avatarOptions[0]);
  const [showAvatarPicker, setShowAvatarPicker] = useState(false);
  const [isEditing, setIsEditing] = useState(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    updateProfile({
      ...formData,
      avatar: selectedAvatar,
    });
    setIsEditing(false);
    // 显示成功提示
    const toast = document.createElement('div');
    toast.className = 'fixed bottom-8 left-1/2 transform -translate-x-1/2 bg-green-500 text-white px-6 py-3 rounded-full shadow-lg z-50 flex items-center space-x-2 animate-fade-in';
    toast.innerHTML = '<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg><span>保存成功！</span>';
    document.body.appendChild(toast);
    setTimeout(() => toast.remove(), 3000);
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-orange-50 via-white to-red-50 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900">
      {/* 顶部背景装饰 */}
      <div className="h-48 bg-gradient-to-r from-orange-400 via-red-400 to-pink-400 relative overflow-hidden">
        <div className="absolute inset-0 opacity-30">
          <div className="absolute top-10 left-10 w-32 h-32 bg-white rounded-full blur-3xl"></div>
          <div className="absolute bottom-10 right-20 w-40 h-40 bg-yellow-200 rounded-full blur-3xl"></div>
        </div>
      </div>

      <div className="container -mt-32 relative z-10 pb-12">
        <div className="grid grid-cols-1 lg:grid-cols-12 gap-6">
          {/* 左侧卡片 - 用户信息 */}
          <div className="lg:col-span-4">
            <div
              ref={cardSpotlight.ref as React.RefObject<HTMLDivElement>}
              className="glass-card rounded-2xl p-6 text-center overflow-hidden relative shadow-xl"
              style={cardSpotlight.spotlightStyle}
              {...cardSpotlight.handlers}
            >
              {/* 头像 */}
              <div className="relative inline-block mb-4">
                <div className="w-28 h-28 rounded-full overflow-hidden ring-4 ring-white dark:ring-gray-700 shadow-xl mx-auto">
                  <img
                    src={selectedAvatar}
                    alt="用户头像"
                    className="w-full h-full object-cover"
                  />
                </div>
                <button
                  onClick={() => setShowAvatarPicker(!showAvatarPicker)}
                  className="absolute bottom-0 right-0 w-9 h-9 bg-gradient-to-r from-orange-500 to-red-500 rounded-full shadow-lg flex items-center justify-center hover:scale-110 transition-transform ring-2 ring-white dark:ring-gray-800"
                >
                  <Camera className="w-4 h-4 text-white" />
                </button>
              </div>

              {/* 用户名和等级 */}
              <h2 className="text-xl font-bold text-gray-900 dark:text-white mb-1">{formData.nickname || formData.username}</h2>
              <p className="text-sm text-gray-500 dark:text-gray-400 mb-4">@{formData.username}</p>

              {/* 会员标签 */}
              <div className="inline-flex items-center space-x-1 bg-gradient-to-r from-amber-400 to-orange-400 text-white px-3 py-1 rounded-full text-xs font-medium mb-4">
                <Crown className="w-3 h-3" />
                <span>金牌会员</span>
              </div>

              {/* 统计数据 */}
              <div className="grid grid-cols-3 gap-2 pt-4 border-t border-gray-100 dark:border-gray-700">
                <div className="text-center">
                  <div className="text-xl font-bold text-gray-900 dark:text-white">128</div>
                  <div className="text-xs text-gray-500 dark:text-gray-400">订单</div>
                </div>
                <div className="text-center">
                  <div className="text-xl font-bold text-gray-900 dark:text-white">36</div>
                  <div className="text-xs text-gray-500 dark:text-gray-400">收藏</div>
                </div>
                <div className="text-center">
                  <div className="text-xl font-bold text-gray-900 dark:text-white">520</div>
                  <div className="text-xs text-gray-500 dark:text-gray-400">积分</div>
                </div>
              </div>
            </div>

            {/* 头像选择器 */}
            {showAvatarPicker && (
              <div className="glass-card rounded-2xl p-4 mt-4 animate-fade-in">
                <h3 className="text-sm font-medium text-gray-700 dark:text-gray-300 mb-3">选择头像</h3>
                <div className="grid grid-cols-4 gap-2">
                  {avatarOptions.map((avatar, index) => (
                    <button
                      key={index}
                      onClick={() => {
                        setSelectedAvatar(avatar);
                        setShowAvatarPicker(false);
                      }}
                      className={`w-14 h-14 rounded-full overflow-hidden ring-2 transition-all hover:scale-105 ${
                        selectedAvatar === avatar
                          ? 'ring-orange-500 ring-offset-2'
                          : 'ring-transparent hover:ring-gray-300'
                      }`}
                    >
                      <img src={avatar} alt={`头像 ${index + 1}`} className="w-full h-full object-cover" />
                    </button>
                  ))}
                </div>
              </div>
            )}
          </div>

          {/* 右侧 - 表单区域 */}
          <div className="lg:col-span-8">
            <div className="glass-card rounded-2xl overflow-hidden shadow-xl">
              {/* 标题栏 */}
              <div className="px-6 py-4 border-b border-gray-100 dark:border-gray-700 flex items-center justify-between bg-gradient-to-r from-orange-50 to-red-50 dark:from-gray-800 dark:to-gray-800">
                <div className="flex items-center space-x-3">
                  <div className="w-10 h-10 bg-gradient-to-r from-orange-500 to-red-500 rounded-xl flex items-center justify-center">
                    <User className="w-5 h-5 text-white" />
                  </div>
                  <div>
                    <h1 className="text-lg font-bold text-gray-900 dark:text-white">个人信息</h1>
                    <p className="text-xs text-gray-500 dark:text-gray-400">管理您的账户信息</p>
                  </div>
                </div>
                {!isEditing && (
                  <button
                    onClick={() => setIsEditing(true)}
                    className="flex items-center space-x-1 text-primary hover:text-primary-hover transition-colors"
                  >
                    <Edit3 className="w-4 h-4" />
                    <span className="text-sm">编辑</span>
                  </button>
                )}
              </div>

              {/* 表单内容 */}
              <form onSubmit={handleSubmit} className="p-6">
                <div className="grid grid-cols-1 md:grid-cols-2 gap-x-8 gap-y-5">
                  {/* 用户名 */}
                  <div className="group">
                    <label className="flex items-center space-x-2 text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                      <User className="w-4 h-4 text-gray-400" />
                      <span>用户名</span>
                    </label>
                    <input
                      type="text"
                      name="username"
                      value={formData.username}
                      onChange={handleChange}
                      disabled={!isEditing}
                      className={`w-full px-4 py-3 rounded-xl border-2 transition-all ${
                        isEditing
                          ? 'border-gray-200 dark:border-gray-600 focus:border-primary focus:ring-4 focus:ring-primary/10 bg-white dark:bg-gray-800'
                          : 'border-transparent bg-gray-50 dark:bg-gray-800/50'
                      } dark:text-white`}
                      placeholder="请输入用户名"
                    />
                  </div>

                  {/* 昵称 */}
                  <div className="group">
                    <label className="flex items-center space-x-2 text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                      <Star className="w-4 h-4 text-gray-400" />
                      <span>昵称</span>
                    </label>
                    <input
                      type="text"
                      name="nickname"
                      value={formData.nickname}
                      onChange={handleChange}
                      disabled={!isEditing}
                      className={`w-full px-4 py-3 rounded-xl border-2 transition-all ${
                        isEditing
                          ? 'border-gray-200 dark:border-gray-600 focus:border-primary focus:ring-4 focus:ring-primary/10 bg-white dark:bg-gray-800'
                          : 'border-transparent bg-gray-50 dark:bg-gray-800/50'
                      } dark:text-white`}
                      placeholder="请输入昵称"
                    />
                  </div>

                  {/* 邮箱 */}
                  <div className="group">
                    <label className="flex items-center space-x-2 text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                      <Mail className="w-4 h-4 text-gray-400" />
                      <span>邮箱</span>
                    </label>
                    <input
                      type="email"
                      name="email"
                      value={formData.email}
                      onChange={handleChange}
                      disabled={!isEditing}
                      className={`w-full px-4 py-3 rounded-xl border-2 transition-all ${
                        isEditing
                          ? 'border-gray-200 dark:border-gray-600 focus:border-primary focus:ring-4 focus:ring-primary/10 bg-white dark:bg-gray-800'
                          : 'border-transparent bg-gray-50 dark:bg-gray-800/50'
                      } dark:text-white`}
                      placeholder="请输入邮箱"
                    />
                  </div>

                  {/* 手机号 */}
                  <div className="group">
                    <label className="flex items-center space-x-2 text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                      <Phone className="w-4 h-4 text-gray-400" />
                      <span>手机号</span>
                    </label>
                    <input
                      type="tel"
                      name="phone"
                      value={formData.phone}
                      onChange={handleChange}
                      disabled={!isEditing}
                      className={`w-full px-4 py-3 rounded-xl border-2 transition-all ${
                        isEditing
                          ? 'border-gray-200 dark:border-gray-600 focus:border-primary focus:ring-4 focus:ring-primary/10 bg-white dark:bg-gray-800'
                          : 'border-transparent bg-gray-50 dark:bg-gray-800/50'
                      } dark:text-white`}
                      placeholder="请输入手机号"
                    />
                  </div>

                  {/* 性别 */}
                  <div className="group">
                    <label className="flex items-center space-x-2 text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                      <User className="w-4 h-4 text-gray-400" />
                      <span>性别</span>
                    </label>
                    <select
                      name="gender"
                      value={formData.gender}
                      onChange={handleChange}
                      disabled={!isEditing}
                      className={`w-full px-4 py-3 rounded-xl border-2 transition-all appearance-none ${
                        isEditing
                          ? 'border-gray-200 dark:border-gray-600 focus:border-primary focus:ring-4 focus:ring-primary/10 bg-white dark:bg-gray-800'
                          : 'border-transparent bg-gray-50 dark:bg-gray-800/50'
                      } dark:text-white`}
                    >
                      <option value="male">男</option>
                      <option value="female">女</option>
                      <option value="other">保密</option>
                    </select>
                  </div>

                  {/* 生日 */}
                  <div className="group">
                    <label className="flex items-center space-x-2 text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                      <Calendar className="w-4 h-4 text-gray-400" />
                      <span>生日</span>
                    </label>
                    <input
                      type="date"
                      name="birthday"
                      value={formData.birthday}
                      onChange={handleChange}
                      disabled={!isEditing}
                      className={`w-full px-4 py-3 rounded-xl border-2 transition-all ${
                        isEditing
                          ? 'border-gray-200 dark:border-gray-600 focus:border-primary focus:ring-4 focus:ring-primary/10 bg-white dark:bg-gray-800'
                          : 'border-transparent bg-gray-50 dark:bg-gray-800/50'
                      } dark:text-white`}
                    />
                  </div>
                </div>

                {/* 个人简介 */}
                <div className="mt-5">
                  <label className="flex items-center space-x-2 text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                    <Edit3 className="w-4 h-4 text-gray-400" />
                    <span>个人简介</span>
                  </label>
                  <textarea
                    name="bio"
                    value={formData.bio}
                    onChange={handleChange}
                    disabled={!isEditing}
                    rows={3}
                    className={`w-full px-4 py-3 rounded-xl border-2 transition-all resize-none ${
                      isEditing
                        ? 'border-gray-200 dark:border-gray-600 focus:border-primary focus:ring-4 focus:ring-primary/10 bg-white dark:bg-gray-800'
                        : 'border-transparent bg-gray-50 dark:bg-gray-800/50'
                    } dark:text-white`}
                    placeholder="介绍一下自己吧~"
                  />
                </div>

                {/* 保存按钮 */}
                {isEditing && (
                  <div className="flex items-center justify-end space-x-3 mt-6 pt-6 border-t border-gray-100 dark:border-gray-700">
                    <button
                      type="button"
                      onClick={() => setIsEditing(false)}
                      className="px-6 py-2.5 text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-xl transition-colors"
                    >
                      取消
                    </button>
                    <button
                      type="submit"
                      className="flex items-center space-x-2 px-8 py-2.5 bg-gradient-to-r from-orange-500 to-red-500 text-white rounded-xl hover:from-orange-600 hover:to-red-600 transition-all shadow-lg hover:shadow-xl hover:-translate-y-0.5"
                    >
                      <Check className="w-4 h-4" />
                      <span>保存修改</span>
                    </button>
                  </div>
                )}
              </form>
            </div>
          </div>
        </div>
      </div>

      {/* 动画样式 */}
      <style>{`
        @keyframes fade-in {
          from { opacity: 0; transform: translateY(10px); }
          to { opacity: 1; transform: translateY(0); }
        }
        .animate-fade-in {
          animation: fade-in 0.3s ease-out;
        }
      `}</style>
    </div>
  );
};
