import { create } from 'zustand';
import { persist } from 'zustand/middleware';

export interface Coupon {
  id: string;
  name: string;
  discount: number; // 折扣金额
  minSpend: number; // 最低消费
  scope: string; // 适用范围 'all' | category name
  endTime: string;
}

export interface UserCoupon {
  id: string;
  couponId: string;
  userId: string;
  status: 'unused' | 'used' | 'expired';
  claimedAt: string;
  usedAt?: string;
}

interface CouponState {
  availableCoupons: Coupon[];
  userCoupons: UserCoupon[];
  claimCoupon: (couponId: string) => boolean;
  useCoupon: (userCouponId: string) => void;
}

// Mock 优惠券数据
const mockCoupons: Coupon[] = [
  {
    id: '1',
    name: '新人专享优惠券',
    discount: 50,
    minSpend: 199,
    scope: 'all',
    endTime: '2024-12-31',
  },
  {
    id: '2',
    name: '服装品类优惠券',
    discount: 30,
    minSpend: 299,
    scope: '服装',
    endTime: '2024-12-31',
  },
  {
    id: '3',
    name: '数码产品优惠券',
    discount: 100,
    minSpend: 999,
    scope: '数码',
    endTime: '2024-12-31',
  },
  {
    id: '4',
    name: '美妆护肤优惠券',
    discount: 20,
    minSpend: 199,
    scope: '美妆',
    endTime: '2024-12-31',
  },
  {
    id: '5',
    name: '限时秒杀优惠券',
    discount: 200,
    minSpend: 2000,
    scope: 'all',
    endTime: '2024-01-31',
  },
  {
    id: '6',
    name: '满减优惠券',
    discount: 15,
    minSpend: 100,
    scope: 'all',
    endTime: '2024-12-31',
  },
];

export const useCouponStore = create<CouponState>()(
  persist(
    (set, get) => ({
      availableCoupons: mockCoupons,
      userCoupons: [],

      claimCoupon: (couponId) => {
        const { userCoupons } = get();
        
        // 检查是否已领取
        if (userCoupons.some((uc) => uc.couponId === couponId)) {
          return false;
        }

        const now = new Date();
        const claimedAt = `${now.getFullYear()}-${String(now.getMonth() + 1).padStart(2, '0')}-${String(now.getDate()).padStart(2, '0')}`;

        const newUserCoupon: UserCoupon = {
          id: `uc-${Date.now()}`,
          couponId,
          userId: 'current-user',
          status: 'unused',
          claimedAt,
        };

        set((state) => ({
          userCoupons: [...state.userCoupons, newUserCoupon],
        }));

        return true;
      },

      useCoupon: (userCouponId) => {
        const now = new Date();
        const usedAt = `${now.getFullYear()}-${String(now.getMonth() + 1).padStart(2, '0')}-${String(now.getDate()).padStart(2, '0')}`;

        set((state) => ({
          userCoupons: state.userCoupons.map((uc) =>
            uc.id === userCouponId ? { ...uc, status: 'used' as const, usedAt } : uc
          ),
        }));
      },
    }),
    {
      name: 'coupon-storage',
    }
  )
);
