import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import type { CartItem } from '../types/cart';

interface CartState {
  items: CartItem[];
  total: number;
  addItem: (item: CartItem) => void;
  removeItem: (id: string) => void;
  updateQuantity: (id: string, quantity: number) => void;
  toggleSelect: (id: string) => void;
  toggleSelectAll: () => void;
  clearCart: () => void;
  calculateTotal: () => void;
}

export const useCartStore = create<CartState>()(
  persist(
    (set, get) => ({
      items: [],
      total: 0,
      addItem: (item) => {
        const items = get().items;
        const existItem = items.find(
          (i) => i.productId === item.productId && i.specId === item.specId
        );
        
        if (existItem) {
          const newItems = items.map((i) =>
            i.id === existItem.id
              ? { ...i, quantity: i.quantity + item.quantity }
              : i
          );
          set({ items: newItems });
        } else {
          set({ items: [...items, item] });
        }
        get().calculateTotal();
      },
      removeItem: (id) => {
        const items = get().items.filter((i) => i.id !== id);
        set({ items });
        get().calculateTotal();
      },
      updateQuantity: (id, quantity) => {
        const items = get().items.map((i) =>
          i.id === id ? { ...i, quantity } : i
        );
        set({ items });
        get().calculateTotal();
      },
      toggleSelect: (id) => {
        const items = get().items.map((i) =>
          i.id === id ? { ...i, selected: !i.selected } : i
        );
        set({ items });
        get().calculateTotal();
      },
      toggleSelectAll: () => {
        const items = get().items;
        const allSelected = items.every((i) => i.selected);
        const newItems = items.map((i) => ({ ...i, selected: !allSelected }));
        set({ items: newItems });
        get().calculateTotal();
      },
      clearCart: () => set({ items: [], total: 0 }),
      calculateTotal: () => {
        const total = get().items
          .filter((i) => i.selected)
          .reduce((sum, i) => sum + i.product.price * i.quantity, 0);
        set({ total });
      },
    }),
    {
      name: 'cart-storage',
    }
  )
);
