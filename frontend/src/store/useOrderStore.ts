import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import type { Order, OrderStatus } from '../types/order';

interface OrderState {
  orders: Order[];
  createOrder: (order: Omit<Order, 'id' | 'orderNo' | 'createdAt'>) => string;
  cancelOrder: (orderId: string) => void;
  payOrder: (orderId: string) => void;
  shipOrder: (orderId: string) => void;
  confirmReceive: (orderId: string) => void;
  getOrderById: (orderId: string) => Order | undefined;
}

// 生成订单号
const generateOrderNo = () => {
  const now = new Date();
  const year = now.getFullYear();
  const month = String(now.getMonth() + 1).padStart(2, '0');
  const day = String(now.getDate()).padStart(2, '0');
  const random = String(Math.floor(Math.random() * 10000)).padStart(4, '0');
  return `${year}${month}${day}${random}`;
};

// Mock 初始订单数据
const mockOrders: Order[] = [
  {
    id: '1',
    userId: '1',
    orderNo: '202401200001',
    totalAmount: 258,
    payAmount: 258,
    status: 'pending',
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
    userId: '1',
    orderNo: '202401190002',
    totalAmount: 9999,
    payAmount: 9999,
    status: 'shipped',
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
  {
    id: '3',
    userId: '1',
    orderNo: '202401180003',
    totalAmount: 599,
    payAmount: 549,
    status: 'completed',
    receiverName: '张三',
    receiverPhone: '13800138000',
    receiverAddress: '北京市朝阳区某某街道',
    items: [
      {
        id: 'i3',
        orderId: '3',
        productId: '7',
        productName: 'Nike Air Max 270 运动鞋',
        productImage: 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=200&h=200&fit=crop',
        specName: '黑白配色 42码',
        price: 599,
        quantity: 1,
      },
    ],
    createdAt: '2024-01-18 10:00:00',
    paidAt: '2024-01-18 10:01:00',
    shippedAt: '2024-01-18 14:00:00',
    completedAt: '2024-01-20 16:00:00',
  },
];

export const useOrderStore = create<OrderState>()(
  persist(
    (set, get) => ({
      orders: mockOrders,

      createOrder: (orderData) => {
        const id = String(Date.now());
        const orderNo = generateOrderNo();
        const now = new Date();
        const createdAt = `${now.getFullYear()}-${String(now.getMonth() + 1).padStart(2, '0')}-${String(now.getDate()).padStart(2, '0')} ${String(now.getHours()).padStart(2, '0')}:${String(now.getMinutes()).padStart(2, '0')}:${String(now.getSeconds()).padStart(2, '0')}`;
        
        const newOrder: Order = {
          ...orderData,
          id,
          orderNo,
          createdAt,
        };

        set((state) => ({
          orders: [newOrder, ...state.orders],
        }));

        return id;
      },

      cancelOrder: (orderId) => {
        set((state) => ({
          orders: state.orders.map((order) =>
            order.id === orderId ? { ...order, status: 'cancelled' as OrderStatus } : order
          ),
        }));
      },

      payOrder: (orderId) => {
        const now = new Date();
        const paidAt = `${now.getFullYear()}-${String(now.getMonth() + 1).padStart(2, '0')}-${String(now.getDate()).padStart(2, '0')} ${String(now.getHours()).padStart(2, '0')}:${String(now.getMinutes()).padStart(2, '0')}:${String(now.getSeconds()).padStart(2, '0')}`;
        
        set((state) => ({
          orders: state.orders.map((order) =>
            order.id === orderId ? { ...order, status: 'paid' as OrderStatus, paidAt } : order
          ),
        }));
      },

      shipOrder: (orderId) => {
        const now = new Date();
        const shippedAt = `${now.getFullYear()}-${String(now.getMonth() + 1).padStart(2, '0')}-${String(now.getDate()).padStart(2, '0')} ${String(now.getHours()).padStart(2, '0')}:${String(now.getMinutes()).padStart(2, '0')}:${String(now.getSeconds()).padStart(2, '0')}`;
        
        set((state) => ({
          orders: state.orders.map((order) =>
            order.id === orderId ? { ...order, status: 'shipped' as OrderStatus, shippedAt } : order
          ),
        }));
      },

      confirmReceive: (orderId) => {
        const now = new Date();
        const completedAt = `${now.getFullYear()}-${String(now.getMonth() + 1).padStart(2, '0')}-${String(now.getDate()).padStart(2, '0')} ${String(now.getHours()).padStart(2, '0')}:${String(now.getMinutes()).padStart(2, '0')}:${String(now.getSeconds()).padStart(2, '0')}`;
        
        set((state) => ({
          orders: state.orders.map((order) =>
            order.id === orderId ? { ...order, status: 'completed' as OrderStatus, completedAt } : order
          ),
        }));
      },

      getOrderById: (orderId) => {
        return get().orders.find((order) => order.id === orderId);
      },
    }),
    {
      name: 'order-storage',
    }
  )
);
