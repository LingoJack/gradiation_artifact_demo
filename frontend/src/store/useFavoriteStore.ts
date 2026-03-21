import { create } from 'zustand';
import { persist } from 'zustand/middleware';

interface FavoriteState {
  favorites: string[]; // 商品 ID 列表
  addFavorite: (productId: string) => void;
  removeFavorite: (productId: string) => void;
  isFavorite: (productId: string) => boolean;
  toggleFavorite: (productId: string) => void;
}

export const useFavoriteStore = create<FavoriteState>()(
  persist(
    (set, get) => ({
      favorites: [],

      addFavorite: (productId) => {
        set((state) => ({
          favorites: state.favorites.includes(productId)
            ? state.favorites
            : [...state.favorites, productId],
        }));
      },

      removeFavorite: (productId) => {
        set((state) => ({
          favorites: state.favorites.filter((id) => id !== productId),
        }));
      },

      isFavorite: (productId) => {
        return get().favorites.includes(productId);
      },

      toggleFavorite: (productId) => {
        const { favorites, addFavorite, removeFavorite } = get();
        if (favorites.includes(productId)) {
          removeFavorite(productId);
        } else {
          addFavorite(productId);
        }
      },
    }),
    {
      name: 'favorite-storage',
    }
  )
);
