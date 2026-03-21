import React, { useState } from 'react';
import { ORDER_STATUS_TEXT, type OrderStatus } from '../../types/order';

// Mock 订单数据
const mockOrders = [
  {
    id: '1',
    orderNo: '202401200001',
    totalAmount: 258,
    payAmount: 258,
    status: 'pending' as OrderStatus,
    receiverName: '张三',
    receiverPhone: '13800138000',
    receiverAddress: '北京市朝阳区某某街道',
    items: [
      {
        id: 'i1',
        orderId: '1',
        productId: '1',
        productName: '时尚休闲连帽卫衣 男士秋季新款',
        productImage: 'https://images.unsplash.com/photo-1556821840-3a63f95609a7?w=200&h=200&fit=crop',
        specName: '黑色',
        price: 129,
        quantity: 2,
      },
    ],
    createdAt: '2024-01-20 10:30:00',
  },
  {
    id: '2',
    orderNo: '202401190002',
    totalAmount: 9999,
    payAmount: 9999,
    status: 'shipped' as OrderStatus,
    receiverName: '张三',
    receiverPhone: '13800138000',
    receiverAddress: '北京市朝阳区某某街道',
    items: [
      {
        id: 'i2',
        orderId: '2',
        productId: '2',
        productName: 'Apple iPhone 15 Pro Max 256GB',
        productImage: 'https://images.unsplash.com/photo-1695048133142-1a20484d2569?w=200&h=200&fit=crop',
        specName: '深空黑',
        price: 9999,
        quantity: 1,
      },
    ],
    createdAt: '2024-01-19 15:20:00',
    paidAt: '2024-01-19 15:21:00',
    shippedAt: '2024-01-20 09:00:00',
  },
];

export const OrderList: React.FC = () => {
  const [activeTab, setActiveTab] = useState<OrderStatus | 'all'>('all');

  const tabs: Array<{ key: OrderStatus | 'all'; label: string }> = [
    { key: 'all', label: '全部订单' },
    { key: 'pending', label: '待付款' },
    { key: 'paid', label: '待发货' },
    { key: 'shipped', label: '待收货' },
    { key: 'completed', label: '已完成' },
  ];

  const filteredOrders =
    activeTab === 'all'
      ? mockOrders
      : mockOrders.filter((o) => o.status === activeTab);

  return (
    <div className="container py-8">
      <h1 className="text-2xl font-bold mb-6 dark:text-white">我的订单</h1>

      {/* Tabs */}
      <div className="glass-card rounded-lg mb-6">
        <div className="flex border-b border-gray-200 dark:border-gray-700">
          {tabs.map((tab) => (
            <button
              key={tab.key}
              onClick={() => setActiveTab(tab.key)}
              className={`px-6 py-4 font-medium transition-colors ${
                activeTab === tab.key
                  ? 'text-primary border-b-2 border-primary'
                  : 'text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-white'
              }`}
            >
              {tab.label}
            </button>
          ))}
        </div>
      </div>

      {/* 订单列表 */}
      <div className="space-y-4">
        {filteredOrders.length === 0 ? (
          <div className="glass-card rounded-xl p-12 text-center">
            <p className="text-gray-500 dark:text-gray-400">暂无订单</p>
          </div>
        ) : (
          filteredOrders.map((order) => (
            <div key={order.id} className="glass-card rounded-xl overflow-hidden">
              {/* 订单头部 */}
              <div className="bg-gray-50 dark:bg-gray-800/50 px-6 py-3 flex items-center justify-between text-sm">
                <div className="flex items-center space-x-6 text-gray-600 dark:text-gray-400">
                  <span>订单号：{order.orderNo}</span>
                  <span>{order.createdAt}</span>
                </div>
                <span className="font-bold text-primary">
                  {ORDER_STATUS_TEXT[order.status]}
                </span>
              </div>

              {/* 商品列表 */}
              <div className="p-6">
                {order.items.map((item) => (
                  <div key={item.id} className="flex space-x-4">
                    <img
                      src={item.productImage}
                      alt={item.productName}
                      className="w-20 h-20 object-cover rounded bg-gray-100 dark:bg-gray-700"
                      onError={(e) => {
                        const target = e.target as HTMLImageElement;
                        target.src = 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="80" height="80"%3E%3Crect fill="%23f3f4f6" width="80" height="80"/%3E%3Ctext fill="%239ca3af" x="50%25" y="50%25" text-anchor="middle" dy=".3em" font-size="10"%3E暂无图片%3C/text%3E%3C/svg%3E';
                      }}
                    />
                    <div className="flex-1">
                      <p className="text-sm dark:text-white">{item.productName}</p>
                      {item.specName && (
                        <p className="text-xs text-gray-500 dark:text-gray-400 mt-1">
                          规格：{item.specName}
                        </p>
                      )}
                      <div className="flex items-center justify-between mt-2">
                        <span className="text-primary font-bold">
                          ¥{item.price}
                        </span>
                        <span className="text-sm text-gray-500 dark:text-gray-400">
                          x{item.quantity}
                        </span>
                      </div>
                    </div>
                  </div>
                ))}

                <div className="flex items-center justify-between mt-4 pt-4 border-t border-gray-200 dark:border-gray-700">
                  <div className="text-sm text-gray-600 dark:text-gray-400">
                    共 {order.items.length} 件商品 合计：
                    <span className="text-lg font-bold text-primary">
                      ¥{order.payAmount}
                    </span>
                  </div>
                  <div className="flex space-x-2">
                    {order.status === 'pending' && (
                      <>
                        <button className="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded hover:bg-gray-50 dark:hover:bg-gray-700 dark:text-white transition-colors">
                          取消订单
                        </button>
                        <button className="px-4 py-2 bg-primary text-white rounded hover:bg-primary-hover">
                          立即支付
                        </button>
                      </>
                    )}
                    {order.status === 'shipped' && (
                      <button className="px-4 py-2 bg-primary text-white rounded hover:bg-primary-hover">
                        确认收货
                      </button>
                    )}
                    {order.status === 'completed' && (
                      <button className="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded hover:bg-gray-50 dark:hover:bg-gray-700 dark:text-white transition-colors">
                        评价
                      </button>
                    )}
                  </div>
                </div>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  );
};
