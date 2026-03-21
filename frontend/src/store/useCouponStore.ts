import { create } from 'zustand';
import { persist } from 'zustand/middleware';

export interface Coupon {
  id: string;
  name: string;
  type: 'fixed' | 'percent'; // 固定金额 or 百分比
  value: number; // 减免金额 or 折扣百分比
  minAmount: number; // 最低消费门槛
  startTime: string;
  endTime: string;
  status: 'available' | 'used' | 'expired';
}

interface CouponState {
  coupons: Coupon[];
  points: number;
  redPacket: number;
  
  // Actions
  addCoupon: (coupon: Coupon) => void;
  useCoupon: (id: string) => void;
  addPoints: (amount: number) => void;
  usePoints: (amount: number) => void;
  addRedPacket: (amount: number) => void;
  useRedPacket: (amount: number) => void;
  getAvailableCoupons: () => Coupon[];
  reset: () => void;
}

// Mock 优惠券数据（新用户注册时赠送）
const defaultCoupons: Coupon[] = [
  {
    id: '1',
    name: '新人专享券',
    type: 'fixed',
    value: 50,
    minAmount: 100,
    startTime: '2024-01-01',
    endTime: '2024-12-31',
    status: 'available',
  },
  {
    id: '2',
    name: '满200减30',
    type: 'fixed',
    value: 30,
    minAmount: 200,
    startTime: '2024-01-01',
    endTime: '2024-06-30',
    status: 'available',
  },
  {
    id: '3',
    name: '全场9折券',
    type: 'percent',
    value: 10,
    minAmount: 50,
    startTime: '2024-01-01',
    endTime: '2024-12-31',
    status: 'available',
  },
];

const defaultPoints = 100; // 新用户初始积分
const defaultRedPacket = 10; // 新用户初始红包

export const useCouponStore = create<CouponState>()(
  persist(
    (set, get) => ({
      coupons: [],
      points: 0,
      redPacket: 0,

      addCoupon: (coupon) =>
        set((state) => ({
          coupons: [...state.coupons, coupon],
        })),

      useCoupon: (id) =>
        set((state) => ({
          coupons: state.coupons.map((c) =>
            c.id === id ? { ...c, status: 'used' as const } : c
          ),
        })),

      addPoints: (amount) =>
        set((state) => ({
          points: state.points + amount,
        })),

      usePoints: (amount) =>
        set((state) => ({
          points: Math.max(0, state.points - amount),
        })),

      addRedPacket: (amount) =>
        set((state) => ({
          redPacket: state.redPacket + amount,
        })),

      useRedPacket: (amount) =>
        set((state) => ({
          redPacket: Math.max(0, state.redPacket - amount),
        })),

      getAvailableCoupons: () => {
        const { coupons } = get();
        return coupons.filter((c) => c.status === 'available');
      },

      reset: () =>
        set({
          coupons: defaultCoupons,
          points: defaultPoints,
          redPacket: defaultRedPacket,
        }),
    }),
    {
      name: 'coupon-storage',
    }
  )
);

// 初始化新用户优惠数据的方法
export const initNewUserBenefits = () => {
  const store = useCouponStore.getState();
  if (store.coupons.length === 0) {
    store.reset();
  }
};
