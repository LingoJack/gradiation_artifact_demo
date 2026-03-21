import React, { useState } from 'react';
import { Plus, Edit2, Trash2, MapPin, Check } from 'lucide-react';
import { useSpotlight } from '../../hooks/useSpotlight';

interface Address {
  id: string;
  receiver: string;
  phone: string;
  province: string;
  city: string;
  district: string;
  detail: string;
  isDefault: boolean;
}

// Mock 数据
const mockAddresses: Address[] = [
  {
    id: '1',
    receiver: '张三',
    phone: '13800138000',
    province: '北京市',
    city: '北京市',
    district: '朝阳区',
    detail: '某某街道某某小区1号楼101室',
    isDefault: true,
  },
  {
    id: '2',
    receiver: '李四',
    phone: '13900139000',
    province: '上海市',
    city: '上海市',
    district: '浦东新区',
    detail: '某某路某某大厦A座2001室',
    isDefault: false,
  },
];

export const Addresses: React.FC = () => {
  const [addresses, setAddresses] = useState<Address[]>(mockAddresses);
  const [showModal, setShowModal] = useState(false);
  const [editingAddress, setEditingAddress] = useState<Address | null>(null);
  const cardSpotlight = useSpotlight();
  const modalSpotlight = useSpotlight();

  const handleSetDefault = (id: string) => {
    setAddresses(
      addresses.map((addr) => ({
        ...addr,
        isDefault: addr.id === id,
      }))
    );
  };

  const handleDelete = (id: string) => {
    if (window.confirm('确定要删除这个地址吗？')) {
      setAddresses(addresses.filter((addr) => addr.id !== id));
    }
  };

  const handleEdit = (address: Address) => {
    setEditingAddress(address);
    setShowModal(true);
  };

  const handleAdd = () => {
    setEditingAddress(null);
    setShowModal(true);
  };

  return (
    <div className="container py-8">
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold dark:text-white">收货地址</h1>
        <button
          onClick={handleAdd}
          className="flex items-center space-x-2 px-4 py-2 bg-primary text-white rounded-lg hover:bg-primary-hover"
        >
          <Plus className="w-4 h-4" />
          <span>新增地址</span>
        </button>
      </div>

      <div className="space-y-4">
        {addresses.length === 0 ? (
          <div className="glass-card rounded-xl p-12 text-center">
            <MapPin className="w-16 h-16 text-gray-300 dark:text-gray-600 mx-auto mb-4" />
            <p className="text-gray-500 dark:text-gray-400">暂无收货地址</p>
            <button
              onClick={handleAdd}
              className="mt-4 text-primary hover:underline"
            >
              添加收货地址
            </button>
          </div>
        ) : (
          addresses.map((address) => (
            <div
              key={address.id}
              ref={cardSpotlight.ref as React.RefObject<HTMLDivElement>}
              className="glass-liquid rounded-xl p-6 relative overflow-hidden"
              style={cardSpotlight.spotlightStyle}
              {...cardSpotlight.handlers}
            >
              <div className="flex items-start justify-between">
                <div className="flex-1">
                  <div className="flex items-center space-x-4 mb-2">
                    <span className="font-bold dark:text-white">{address.receiver}</span>
                    <span className="text-gray-600 dark:text-gray-400">{address.phone}</span>
                    {address.isDefault && (
                      <span className="px-2 py-0.5 bg-primary/10 text-primary text-xs rounded">
                        默认
                      </span>
                    )}
                  </div>
                  <p className="text-gray-600 dark:text-gray-400">
                    {address.province} {address.city} {address.district} {address.detail}
                  </p>
                </div>
                <div className="flex items-center space-x-4">
                  {!address.isDefault && (
                    <button
                      onClick={() => handleSetDefault(address.id)}
                      className="flex items-center space-x-1 text-gray-500 dark:text-gray-400 hover:text-primary dark:hover:text-primary"
                    >
                      <Check className="w-4 h-4" />
                      <span className="text-sm">设为默认</span>
                    </button>
                  )}
                  <button
                    onClick={() => handleEdit(address)}
                    className="flex items-center space-x-1 text-gray-500 dark:text-gray-400 hover:text-primary dark:hover:text-primary"
                  >
                    <Edit2 className="w-4 h-4" />
                    <span className="text-sm">编辑</span>
                  </button>
                  <button
                    onClick={() => handleDelete(address.id)}
                    className="flex items-center space-x-1 text-gray-500 dark:text-gray-400 hover:text-error"
                  >
                    <Trash2 className="w-4 h-4" />
                    <span className="text-sm">删除</span>
                  </button>
                </div>
              </div>
            </div>
          ))
        )}
      </div>

      {/* 添加/编辑地址弹窗 */}
      {showModal && (
        <div className="fixed inset-0 z-50 flex items-center justify-center">
          <div
            className="absolute inset-0 bg-black/40 backdrop-blur-sm"
            onClick={() => setShowModal(false)}
          />
          <div
            ref={modalSpotlight.ref as React.RefObject<HTMLDivElement>}
            className="glass-modal rounded-2xl p-6 w-full max-w-lg mx-4 relative z-10 overflow-hidden"
            style={modalSpotlight.spotlightStyle}
            {...modalSpotlight.handlers}
          >
            <h2 className="text-xl font-bold mb-6 dark:text-white">
              {editingAddress ? '编辑地址' : '新增地址'}
            </h2>
            <form className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <label className="block text-sm text-gray-600 dark:text-gray-400 mb-1">
                    收货人
                  </label>
                  <input
                    type="text"
                    defaultValue={editingAddress?.receiver}
                    className="glass-input w-full px-3 py-2 rounded-lg dark:text-white"
                    placeholder="请输入收货人姓名"
                  />
                </div>
                <div>
                  <label className="block text-sm text-gray-600 dark:text-gray-400 mb-1">
                    手机号
                  </label>
                  <input
                    type="text"
                    defaultValue={editingAddress?.phone}
                    className="glass-input w-full px-3 py-2 rounded-lg dark:text-white"
                    placeholder="请输入手机号"
                  />
                </div>
              </div>
              <div>
                <label className="block text-sm text-gray-600 dark:text-gray-400 mb-1">
                  所在地区
                </label>
                <div className="grid grid-cols-3 gap-2">
                  <select className="glass-input px-3 py-2 rounded-lg dark:text-white">
                    <option>{editingAddress?.province || '请选择省份'}</option>
                  </select>
                  <select className="glass-input px-3 py-2 rounded-lg dark:text-white">
                    <option>{editingAddress?.city || '请选择城市'}</option>
                  </select>
                  <select className="glass-input px-3 py-2 rounded-lg dark:text-white">
                    <option>{editingAddress?.district || '请选择区县'}</option>
                  </select>
                </div>
              </div>
              <div>
                <label className="block text-sm text-gray-600 dark:text-gray-400 mb-1">
                  详细地址
                </label>
                <textarea
                  defaultValue={editingAddress?.detail}
                  className="glass-input w-full px-3 py-2 rounded-lg dark:text-white"
                  rows={3}
                  placeholder="请输入详细地址"
                />
              </div>
              <div className="flex items-center space-x-2">
                <input type="checkbox" id="setDefault" className="w-4 h-4" />
                <label htmlFor="setDefault" className="text-sm dark:text-gray-300">
                  设为默认地址
                </label>
              </div>
              <div className="flex justify-end space-x-4 pt-4">
                <button
                  type="button"
                  onClick={() => setShowModal(false)}
                  className="px-6 py-2 border border-gray-300 dark:border-gray-600 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700 dark:text-white"
                >
                  取消
                </button>
                <button
                  type="submit"
                  className="px-6 py-2 bg-primary text-white rounded-lg hover:bg-primary-hover"
                >
                  保存
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
};
