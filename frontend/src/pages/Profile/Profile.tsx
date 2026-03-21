import React, { useState } from 'react';
import { Camera, Save } from 'lucide-react';
import { useUserStore } from '../../store/useUserStore';
import { useSpotlight } from '../../hooks/useSpotlight';

export const Profile: React.FC = () => {
  const { user } = useUserStore();
  const cardSpotlight = useSpotlight();

  const [formData, setFormData] = useState({
    username: user?.username || '',
    email: user?.email || '',
    phone: user?.phone || '',
    nickname: user?.username || '',
    gender: 'male',
    birthday: '',
    bio: '',
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // TODO: 保存用户信息
    alert('保存成功！');
  };

  return (
    <div className="container py-8">
      <h1 className="text-2xl font-bold mb-6 dark:text-white">个人信息</h1>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        {/* 头像区域 */}
        <div
          ref={cardSpotlight.ref as React.RefObject<HTMLDivElement>}
          className="glass-liquid rounded-xl p-6 text-center overflow-hidden relative"
          style={cardSpotlight.spotlightStyle}
          {...cardSpotlight.handlers}
        >
          <div className="relative inline-block">
            <div className="w-32 h-32 bg-primary rounded-full flex items-center justify-center text-white text-4xl font-bold mx-auto">
              {formData.username.charAt(0).toUpperCase()}
            </div>
            <button className="absolute bottom-0 right-0 w-10 h-10 bg-white dark:bg-gray-700 rounded-full shadow-lg flex items-center justify-center hover:bg-gray-50 dark:hover:bg-gray-600">
              <Camera className="w-5 h-5 text-gray-600 dark:text-gray-300" />
            </button>
          </div>
          <p className="mt-4 text-lg font-medium dark:text-white">{formData.username}</p>
          <p className="text-sm text-gray-500 dark:text-gray-400">{formData.email}</p>
        </div>

        {/* 表单区域 */}
        <div className="lg:col-span-2">
          <div className="glass-card rounded-xl p-6">
            <form onSubmit={handleSubmit} className="space-y-6">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                    用户名
                  </label>
                  <input
                    type="text"
                    name="username"
                    value={formData.username}
                    onChange={handleChange}
                    className="glass-input w-full px-4 py-2 rounded-lg dark:text-white"
                    placeholder="请输入用户名"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                    昵称
                  </label>
                  <input
                    type="text"
                    name="nickname"
                    value={formData.nickname}
                    onChange={handleChange}
                    className="glass-input w-full px-4 py-2 rounded-lg dark:text-white"
                    placeholder="请输入昵称"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                    邮箱
                  </label>
                  <input
                    type="email"
                    name="email"
                    value={formData.email}
                    onChange={handleChange}
                    className="glass-input w-full px-4 py-2 rounded-lg dark:text-white"
                    placeholder="请输入邮箱"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                    手机号
                  </label>
                  <input
                    type="tel"
                    name="phone"
                    value={formData.phone}
                    onChange={handleChange}
                    className="glass-input w-full px-4 py-2 rounded-lg dark:text-white"
                    placeholder="请输入手机号"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                    性别
                  </label>
                  <select
                    name="gender"
                    value={formData.gender}
                    onChange={handleChange}
                    className="glass-input w-full px-4 py-2 rounded-lg dark:text-white"
                  >
                    <option value="male">男</option>
                    <option value="female">女</option>
                    <option value="other">保密</option>
                  </select>
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                    生日
                  </label>
                  <input
                    type="date"
                    name="birthday"
                    value={formData.birthday}
                    onChange={handleChange}
                    className="glass-input w-full px-4 py-2 rounded-lg dark:text-white"
                  />
                </div>
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                  个人简介
                </label>
                <textarea
                  name="bio"
                  value={formData.bio}
                  onChange={handleChange}
                  rows={4}
                  className="glass-input w-full px-4 py-2 rounded-lg dark:text-white"
                  placeholder="介绍一下自己吧~"
                />
              </div>
              <div className="flex justify-end">
                <button
                  type="submit"
                  className="flex items-center space-x-2 px-6 py-2 bg-primary text-white rounded-lg hover:bg-primary-hover"
                >
                  <Save className="w-4 h-4" />
                  <span>保存修改</span>
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
};
