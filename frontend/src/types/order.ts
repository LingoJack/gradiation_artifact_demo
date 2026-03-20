export interface Order {
  id: string;
  userId: string;
  orderNo: string;
  totalAmount: number;
  payAmount: number;
  status: OrderStatus;
  receiverName: string;
  receiverPhone: string;
  receiverAddress: string;
  items: OrderItem[];
  createdAt: string;
  paidAt?: string;
  shippedAt?: string;
  completedAt?: string;
}

export type OrderStatus = 
  | 'pending'      // 待付款
  | 'paid'         // 待发货
  | 'shipped'      // 待收货
  | 'completed'    // 已完成
  | 'cancelled';   // 已取消

export interface OrderItem {
  id: string;
  orderId: string;
  productId: string;
  productName: string;
  productImage: string;
  specName?: string;
  price: number;
  quantity: number;
}

export const ORDER_STATUS_TEXT: Record<OrderStatus, string> = {
  pending: '待付款',
  paid: '待发货',
  shipped: '待收货',
  completed: '已完成',
  cancelled: '已取消',
};
