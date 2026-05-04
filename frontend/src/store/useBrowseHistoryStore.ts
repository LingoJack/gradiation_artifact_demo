import { create } from 'zustand';
import { persist } from 'zustand/middleware';

export interface BrowseHistoryItem {
  id: string;
  name: string;
  price: number;
  mainImage: string;
  category?: string;
  visitedAt: number; // timestamp
}

interface BrowseHistoryState {
  items: BrowseHistoryItem[];
  addItem: (item: Omit<BrowseHistoryItem, 'visitedAt'>) => void;
  removeItem: (id: string) => void;
  clearAll: () => void;
  getRecent: (count?: number) => BrowseHistoryItem[];
}

const MAX_HISTORY_ITEMS = 100;

export const useBrowseHistoryStore = create<BrowseHistoryState>()(
  persist(
    (set, get) => ({
      items: [],

      addItem: (item) => {
        set((state) => {
          // Remove existing item with same id (dedup)
          const filtered = state.items.filter((i) => i.id !== item.id);
          // Add new item at the beginning
          const newItem: BrowseHistoryItem = {
            ...item,
            visitedAt: Date.now(),
          };
          const updated = [newItem, ...filtered];
          // Keep only MAX_HISTORY_ITEMS
          return { items: updated.slice(0, MAX_HISTORY_ITEMS) };
        });
      },

      removeItem: (id) => {
        set((state) => ({
          items: state.items.filter((i) => i.id !== id),
        }));
      },

      clearAll: () => {
        set({ items: [] });
      },

      getRecent: (count = 20) => {
        return get().items.slice(0, count);
      },
    }),
    {
      name: 'browse-history-storage',
    }
  )
);
