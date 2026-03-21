export interface User {
  id: string;
  username: string;
  email: string;
  phone: string;
  avatar?: string;
  nickname?: string;
  gender?: 'male' | 'female' | 'other';
  birthday?: string;
  bio?: string;
  createdAt: string;
}

export interface Address {
  id: string;
  userId: string;
  receiver: string;
  phone: string;
  province: string;
  city: string;
  district: string;
  detail: string;
  isDefault: boolean;
}
