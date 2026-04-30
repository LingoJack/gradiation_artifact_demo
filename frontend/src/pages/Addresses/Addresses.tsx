import React, { useState, useEffect } from 'react';
import { Plus, Edit2, Trash2, MapPin, Check, Loader2 } from 'lucide-react';
import { userApi } from '../../api';
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

// 映射后端数据到前端格式
function mapAddress(raw: any): Address {
  return {
    id: String(raw.id),
    receiver: raw.receiver_name || raw.receiver || '',
    phone: raw.receiver_phone || raw.phone || '',
    province: raw.province || '',
    city: raw.city || '',
    district: raw.district || '',
    detail: raw.detail_address || raw.detail || '',
    isDefault: raw.is_default === 1 || raw.isDefault === true,
  };
}

// 独立的卡片组件，每个实例有自己的 spotlight
const AddressCard: React.FC<{
  address: Address;
  onSetDefault: (id: string) => void;
  onEdit: (address: Address) => void;
  onDelete: (id: string) => void;
  operating?: boolean;
}> = ({ address, onSetDefault, onEdit, onDelete, operating }) => {
  const cardSpotlight = useSpotlight();

  return (
    <div
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
              onClick={() => onSetDefault(address.id)}
              disabled={operating}
              className="flex items-center space-x-1 text-gray-500 dark:text-gray-400 hover:text-primary dark:hover:text-primary disabled:opacity-50"
            >
              <Check className="w-4 h-4" />
              <span className="text-sm">设为默认</span>
            </button>
          )}
          <button
            onClick={() => onEdit(address)}
            disabled={operating}
            className="flex items-center space-x-1 text-gray-500 dark:text-gray-400 hover:text-primary dark:hover:text-primary disabled:opacity-50"
          >
            <Edit2 className="w-4 h-4" />
            <span className="text-sm">编辑</span>
          </button>
          <button
            onClick={() => onDelete(address.id)}
            disabled={operating}
            className="flex items-center space-x-1 text-gray-500 dark:text-gray-400 hover:text-error disabled:opacity-50"
          >
            <Trash2 className="w-4 h-4" />
            <span className="text-sm">删除</span>
          </button>
        </div>
      </div>
    </div>
  );
};

export const Addresses: React.FC = () => {
  const [addresses, setAddresses] = useState<Address[]>([]);
  const [loading, setLoading] = useState(true);
  const [operating, setOperating] = useState<string | null>(null);
  const [showModal, setShowModal] = useState(false);
  const [editingAddress, setEditingAddress] = useState<Address | null>(null);
  const [formData, setFormData] = useState({
    receiver: '',
    phone: '',
    province: '',
    city: '',
    district: '',
    detail: '',
    isDefault: false,
  });
  const modalSpotlight = useSpotlight();

  // 获取地址列表
  useEffect(() => {
    fetchAddresses();
  }, []);

  const fetchAddresses = async () => {
    try {
      setLoading(true);
      const data = await userApi.getAddresses() as any[];
      setAddresses(data.map(mapAddress));
    } catch (err) {
      console.error('Failed to fetch addresses:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleSetDefault = async (id: string) => {
    try {
      setOperating(id);
      await userApi.setDefaultAddress(Number(id));
      await fetchAddresses();
    } catch (err) {
      console.error('Failed to set default:', err);
      alert('设置默认地址失败');
    } finally {
      setOperating(null);
    }
  };

  const handleDelete = async (id: string) => {
    if (!window.confirm('确定要删除这个地址吗？')) return;
    try {
      setOperating(id);
      await userApi.deleteAddress(Number(id));
      await fetchAddresses();
    } catch (err) {
      console.error('Failed to delete:', err);
      alert('删除地址失败');
    } finally {
      setOperating(null);
    }
  };

  const handleEdit = (address: Address) => {
    setEditingAddress(address);
    setFormData({
      receiver: address.receiver,
      phone: address.phone,
      province: address.province,
      city: address.city,
      district: address.district,
      detail: address.detail,
      isDefault: address.isDefault,
    });
    setShowModal(true);
  };

  const handleAdd = () => {
    setEditingAddress(null);
    setFormData({
      receiver: '',
      phone: '',
      province: '',
      city: '',
      district: '',
      detail: '',
      isDefault: false,
    });
    setShowModal(true);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      setOperating('form');
      const data = {
        receiver: formData.receiver,
        phone: formData.phone,
        province: formData.province,
        city: formData.city,
        district: formData.district,
        detail: formData.detail,
        isDefault: formData.isDefault ? 1 : 0,
      };

      if (editingAddress) {
        await userApi.updateAddress(Number(editingAddress.id), data);
      } else {
        await userApi.createAddress(data);
      }

      setShowModal(false);
      await fetchAddresses();
    } catch (err) {
      console.error('Failed to save address:', err);
      alert('保存地址失败');
    } finally {
      setOperating(null);
    }
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center py-40">
        <Loader2 className="w-8 h-8 animate-spin text-primary" />
      </div>
    );
  }

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
            <AddressCard
              key={address.id}
              address={address}
              onSetDefault={handleSetDefault}
              onEdit={handleEdit}
              onDelete={handleDelete}
              operating={operating === address.id}
            />
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
            <form onSubmit={handleSubmit} className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <label className="block text-sm text-gray-600 dark:text-gray-400 mb-1">
                    收货人
                  </label>
                  <input
                    type="text"
                    value={formData.receiver}
                    onChange={(e) => setFormData({ ...formData, receiver: e.target.value })}
                    className="glass-input w-full px-3 py-2 rounded-lg dark:text-white"
                    placeholder="请输入收货人姓名"
                    required
                  />
                </div>
                <div>
                  <label className="block text-sm text-gray-600 dark:text-gray-400 mb-1">
                    手机号
                  </label>
                  <input
                    type="text"
                    value={formData.phone}
                    onChange={(e) => setFormData({ ...formData, phone: e.target.value })}
                    className="glass-input w-full px-3 py-2 rounded-lg dark:text-white"
                    placeholder="请输入手机号"
                    required
                  />
                </div>
              </div>
              <div>
                <label className="block text-sm text-gray-600 dark:text-gray-400 mb-1">
                  所在地区
                </label>
                <div className="grid grid-cols-3 gap-2">
                  <input
                    type="text"
                    value={formData.province}
                    onChange={(e) => setFormData({ ...formData, province: e.target.value })}
                    className="glass-input px-3 py-2 rounded-lg dark:text-white"
                    placeholder="省份"
                    required
                  />
                  <input
                    type="text"
                    value={formData.city}
                    onChange={(e) => setFormData({ ...formData, city: e.target.value })}
                    className="glass-input px-3 py-2 rounded-lg dark:text-white"
                    placeholder="城市"
                    required
                  />
                  <input
                    type="text"
                    value={formData.district}
                    onChange={(e) => setFormData({ ...formData, district: e.target.value })}
                    className="glass-input px-3 py-2 rounded-lg dark:text-white"
                    placeholder="区县"
                    required
                  />
                </div>
              </div>
              <div>
                <label className="block text-sm text-gray-600 dark:text-gray-400 mb-1">
                  详细地址
                </label>
                <textarea
                  value={formData.detail}
                  onChange={(e) => setFormData({ ...formData, detail: e.target.value })}
                  className="glass-input w-full px-3 py-2 rounded-lg dark:text-white"
                  rows={3}
                  placeholder="请输入详细地址"
                  required
                />
              </div>
              <div className="flex items-center space-x-2">
                <input 
                  type="checkbox" 
                  id="setDefault" 
                  className="w-4 h-4"
                  checked={formData.isDefault}
                  onChange={(e) => setFormData({ ...formData, isDefault: e.target.checked })}
                />
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
                  disabled={operating === 'form'}
                  className="px-6 py-2 bg-primary text-white rounded-lg hover:bg-primary-hover disabled:opacity-50 flex items-center gap-2"
                >
                  {operating === 'form' && <Loader2 className="w-4 h-4 animate-spin" />}
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