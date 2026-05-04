import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Check, Plus, Loader2 } from 'lucide-react';
import { userApi } from '../../api/user';
import { orderApi } from '../../api/order';
import { cartApi } from '../../api/cart';
import { showToast } from '../../utils/toast';

// 后端返回的原始购物车项
interface BackendCartItem {
  id: number;
  user_id: number;
  product_id: number;
  sku_id: number;
  quantity: number;
  selected: number;
  product: {
    id: number;
    name: string;
    price: number;
    main_image: string;
    stock: number;
  };
  sku: {
    id: number;
    price: number;
    stock: number;
    spec_values: string;
    image: string;
  } | null;
}

// 后端返回的原始地址
interface BackendAddress {
  id: number;
  receiver_name: string;
  receiver_phone: string;
  province: string;
  city: string;
  district: string;
  detail_address: string;
  is_default: number;
}

// 前端使用的标准化地址
interface Address {
  id: number;
  receiver: string;
  phone: string;
  province: string;
  city: string;
  district: string;
  detail: string;
  isDefault: boolean;
}

// 前端使用的标准化购物车项
interface CartItem {
  id: number;
  productId: number;
  skuId: number | null;
  quantity: number;
  selected: boolean;
  product: {
    id: number;
    name: string;
    price: number;
    mainImage: string;
    stock: number;
  };
  spec: {
    name: string;
    value: string;
  } | null;
}

function mapCartItem(raw: BackendCartItem): CartItem {
  let spec: CartItem['spec'] = null;
  if (raw.sku && raw.sku.spec_values) {
    const parts = raw.sku.spec_values.split(':');
    spec = {
      name: parts[0] || '规格',
      value: parts.slice(1).join(':') || raw.sku.spec_values,
    };
  }
  return {
    id: raw.id,
    productId: raw.product_id,
    skuId: raw.sku_id,
    quantity: raw.quantity,
    selected: raw.selected === 1,
    product: {
      id: raw.product.id,
      name: raw.product.name,
      price: raw.sku ? raw.sku.price : raw.product.price,
      mainImage: raw.product.main_image,
      stock: raw.sku ? raw.sku.stock : raw.product.stock,
    },
    spec,
  };
}

function mapAddress(raw: BackendAddress): Address {
  return {
    id: raw.id,
    receiver: raw.receiver_name,
    phone: raw.receiver_phone,
    province: raw.province,
    city: raw.city,
    district: raw.district,
    detail: raw.detail_address,
    isDefault: raw.is_default === 1,
  };
}

export const Checkout: React.FC = () => {
  const navigate = useNavigate();
  const [items, setItems] = useState<CartItem[]>([]);
  const [addresses, setAddresses] = useState<Address[]>([]);
  const [selectedAddressId, setSelectedAddressId] = useState<number | null>(null);
  const [loading, setLoading] = useState(true);
  const [submitting, setSubmitting] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        // 并行获取购物车和地址数据
        const [cartData, addressData] = await Promise.all([
          cartApi.getCart(),
          userApi.getAddresses(),
        ]);

        const mappedItems = (cartData as BackendCartItem[]).map(mapCartItem);
        const selectedItems = mappedItems.filter((i) => i.selected);

        // 如果没有选中商品，跳转回购物车
        if (selectedItems.length === 0) {
          navigate('/cart');
          return;
        }

        const mappedAddresses = (addressData as BackendAddress[]).map(mapAddress);
        const defaultAddress = mappedAddresses.find((a) => a.isDefault);

        setItems(selectedItems);
        setAddresses(mappedAddresses);
        setSelectedAddressId(defaultAddress?.id || mappedAddresses[0]?.id || null);
      } catch (err) {
        console.error('加载数据失败:', err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [navigate]);

  const selectedTotal = items.reduce((sum, i) => sum + i.product.price * i.quantity, 0);
  const selectedAddress = addresses.find((a) => a.id === selectedAddressId);

  const handleSubmit = async () => {
    if (!selectedAddressId) {
      showToast('请选择收货地址', 'warning');
      return;
    }

    try {
      setSubmitting(true);
      const orderData = {
        addressId: selectedAddressId,
        items: items.map((i) => ({
          productId: i.productId,
          skuId: i.skuId || undefined,
          quantity: i.quantity,
        })),
      };

      const result = await orderApi.createOrder(orderData) as { order_no: string };
      
      // 清空购物车
      await cartApi.clearCart();
      
      showToast(`订单创建成功！订单号：${result.order_no}`, 'success');
      navigate('/orders');
    } catch (err) {
      console.error('创建订单失败:', err);
      showToast('创建订单失败，请重试', 'error');
    } finally {
      setSubmitting(false);
    }
  };

  if (loading) {
    return (
      <div className="container py-8">
        <h1 className="text-2xl font-bold mb-6 dark:text-white">确认订单</h1>
        <div className="flex items-center justify-center py-20">
          <Loader2 className="w-8 h-8 animate-spin text-primary" />
          <span className="ml-3 text-gray-500 dark:text-gray-400">加载中...</span>
        </div>
      </div>
    );
  }

  return (
    <div className="container py-8">
      <h1 className="text-2xl font-bold mb-6 dark:text-white">确认订单</h1>

      <div className="grid grid-cols-3 gap-8">
        <div className="col-span-2 space-y-6">
          {/* 收货地址 */}
          <div className="glass-card rounded-xl p-6">
            <h2 className="text-lg font-bold mb-4 dark:text-white">收货地址</h2>
            {addresses.length === 0 ? (
              <div className="text-center py-8 text-gray-500 dark:text-gray-400">
                <p className="mb-2">暂无收货地址</p>
                <button
                  onClick={() => navigate('/addresses')}
                  className="text-primary hover:underline"
                >
                  添加新地址
                </button>
              </div>
            ) : (
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                {addresses.map((address) => (
                  <div
                    key={address.id}
                    onClick={() => setSelectedAddressId(address.id)}
                    className={`relative p-4 rounded-lg border-2 cursor-pointer transition-all ${
                      selectedAddressId === address.id
                        ? 'border-primary bg-primary/5'
                        : 'border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600'
                    }`}
                  >
                    {selectedAddressId === address.id && (
                      <div className="absolute top-2 right-2 w-5 h-5 bg-primary rounded-full flex items-center justify-center">
                        <Check className="w-3 h-3 text-white" />
                      </div>
                    )}
                    <div className="space-y-2 text-sm">
                      <div className="flex items-center space-x-2">
                        <span className="font-medium dark:text-white">{address.receiver}</span>
                        <span className="dark:text-gray-300">{address.phone}</span>
                        {address.isDefault && (
                          <span className="px-1.5 py-0.5 bg-primary/10 text-primary text-xs rounded">
                            默认
                          </span>
                        )}
                      </div>
                      <p className="text-gray-600 dark:text-gray-400">
                        {address.province} {address.city} {address.district} {address.detail}
                      </p>
                    </div>
                  </div>
                ))}
                {/* 新增地址按钮 */}
                <div
                  onClick={() => navigate('/addresses')}
                  className="p-4 rounded-lg border-2 border-dashed border-gray-300 dark:border-gray-600 cursor-pointer hover:border-primary dark:hover:border-primary transition-colors flex items-center justify-center"
                >
                  <div className="text-center text-gray-400 dark:text-gray-500">
                    <Plus className="w-6 h-6 mx-auto mb-1" />
                    <span className="text-sm">新增地址</span>
                  </div>
                </div>
              </div>
            )}
          </div>

          {/* 商品清单 */}
          <div className="glass-card rounded-xl p-6">
            <h2 className="text-lg font-bold mb-4 dark:text-white">商品清单</h2>
            <div className="space-y-4">
              {items.map((item) => (
                <div key={item.id} className="flex space-x-4">
                  <img
                    src={item.product.mainImage}
                    alt={item.product.name}
                    className="w-20 h-20 object-cover rounded bg-gray-100 dark:bg-gray-700"
                    onError={(e) => {
                      const target = e.target as HTMLImageElement;
                      target.src = 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="80" height="80"%3E%3Crect fill="%23f3f4f6" width="80" height="80"/%3E%3Ctext fill="%239ca3af" x="50%25" y="50%25" text-anchor="middle" dy=".3em" font-size="10"%3E暂无图片%3C/text%3E%3C/svg%3E';
                    }}
                  />
                  <div className="flex-1">
                    <p className="text-sm dark:text-white">{item.product.name}</p>
                    {item.spec && (
                      <p className="text-xs text-gray-500 dark:text-gray-400 mt-1">
                        {item.spec.name}: {item.spec.value}
                      </p>
                    )}
                    <div className="flex items-center justify-between mt-2">
                      <span className="text-primary font-bold">
                        ¥{item.product.price}
                      </span>
                      <span className="text-sm text-gray-500 dark:text-gray-400">
                        x{item.quantity}
                      </span>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>

        {/* 右侧结算信息 */}
        <div className="glass-card rounded-xl p-6 h-fit sticky top-24">
          <h2 className="text-lg font-bold mb-4 dark:text-white">订单信息</h2>
          
          {/* 已选地址摘要 */}
          {selectedAddress && (
            <div className="mb-4 p-3 bg-gray-50 dark:bg-gray-700/50 rounded-lg text-sm">
              <p className="font-medium dark:text-white">
                {selectedAddress.receiver} {selectedAddress.phone}
              </p>
              <p className="text-gray-500 dark:text-gray-400 mt-1 text-xs">
                {selectedAddress.province} {selectedAddress.city} {selectedAddress.district}
              </p>
            </div>
          )}
          
          <div className="space-y-3 text-sm dark:text-gray-300">
            <div className="flex justify-between">
              <span>商品总额：</span>
              <span>¥{selectedTotal.toFixed(2)}</span>
            </div>
            <div className="flex justify-between">
              <span>运费：</span>
              <span className="text-success">免运费</span>
            </div>
            <div className="flex justify-between">
              <span>优惠：</span>
              <span className="text-error">-¥0.00</span>
            </div>
            <div className="border-t border-gray-200 dark:border-gray-700 pt-3 flex justify-between font-bold text-base">
              <span>实付：</span>
              <span className="text-primary text-xl">¥{selectedTotal.toFixed(2)}</span>
            </div>
          </div>
          <button
            onClick={handleSubmit}
            disabled={submitting || !selectedAddressId}
            className="w-full mt-6 py-3 bg-primary text-white rounded-lg hover:bg-primary-hover disabled:bg-gray-300 dark:disabled:bg-gray-600 disabled:cursor-not-allowed flex items-center justify-center"
          >
            {submitting ? (
              <>
                <Loader2 className="w-5 h-5 animate-spin mr-2" />
                提交中...
              </>
            ) : (
              '提交订单'
            )}
          </button>
        </div>
      </div>
    </div>
  );
};
