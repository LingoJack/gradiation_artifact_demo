// Toast 提示工具
type ToastType = 'success' | 'error' | 'info' | 'warning';

// Toast 队列项
interface ToastItem {
  message: string;
  type: ToastType;
}

// 队列控制
let currentToast: HTMLDivElement | null = null;
let toastTimer: ReturnType<typeof setTimeout> | null = null;
let isShowing = false; // 是否正在显示
let toastQueue: ToastItem[] = []; // Toast 队列

// 处理队列中的下一个 Toast
const processQueue = () => {
  // 如果队列不为空且当前没有正在显示的 Toast
  if (toastQueue.length > 0 && !isShowing) {
    const nextToast = toastQueue.shift();
    if (nextToast) {
      displayToast(nextToast.message, nextToast.type);
    }
  }
};

// 实际显示 Toast 的函数
const displayToast = (message: string, type: ToastType) => {
  isShowing = true;
  
  // 如果已有 Toast，先移除
  if (currentToast) {
    currentToast.style.opacity = '0';
    currentToast.style.transform = 'translateX(100%)';
    setTimeout(() => currentToast?.remove(), 150);
    currentToast = null;
  }
  
  // 清除之前的定时器
  if (toastTimer) {
    clearTimeout(toastTimer);
    toastTimer = null;
  }
  
  const toast = document.createElement('div');
  
  const icons = {
    success: '<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>',
    error: '<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>',
    info: '<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>',
    warning: '<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"></path>',
  };
  
  toast.className = `fixed top-24 right-8 glass-dropdown text-gray-900 dark:text-white px-6 py-3 rounded-xl z-50 flex items-center space-x-3 animate-slide-in-right border border-white/40 dark:border-gray-600/30`;
  
  const iconColors = {
    success: 'text-green-500',
    error: 'text-red-500',
    info: 'text-blue-500',
    warning: 'text-amber-500',
  };
  
  toast.innerHTML = `<svg class="w-5 h-5 ${iconColors[type]}" fill="none" stroke="currentColor" viewBox="0 0 24 24">${icons[type]}</svg><span class="font-medium">${message}</span>`;
  document.body.appendChild(toast);
  
  // 记录当前 Toast
  currentToast = toast;
  
  // 添加动画样式
  if (!document.querySelector('#toast-animation-style')) {
    const style = document.createElement('style');
    style.id = 'toast-animation-style';
    style.textContent = `
      @keyframes slide-in-right {
        from {
          opacity: 0;
          transform: translateX(100%);
        }
        to {
          opacity: 1;
          transform: translateX(0);
        }
      }
    `;
    document.head.appendChild(style);
  }
  
  // 设置自动消失定时器
  toastTimer = setTimeout(() => {
    toast.style.opacity = '0';
    toast.style.transform = 'translateX(100%)';
    toast.style.transition = 'all 0.3s ease-in-out';
    setTimeout(() => {
      toast.remove();
      if (currentToast === toast) {
        currentToast = null;
      }
      // 重置显示标记
      isShowing = false;
      // 处理队列中的下一个 Toast
      processQueue();
    }, 300);
  }, 3000);
};

// 导出的 showToast 函数 - 将 Toast 加入队列
export const showToast = (message: string, type: ToastType = 'success') => {
  // 将新的 Toast 加入队列
  toastQueue.push({ message, type });
  // 尝试处理队列
  processQueue();
};

// 复制到剪贴板
export const copyToClipboard = async (text: string): Promise<boolean> => {
  try {
    await navigator.clipboard.writeText(text);
    return true;
  } catch {
    // 降级方案
    const textArea = document.createElement('textarea');
    textArea.value = text;
    textArea.style.position = 'fixed';
    textArea.style.left = '-9999px';
    document.body.appendChild(textArea);
    textArea.select();
    try {
      document.execCommand('copy');
      document.body.removeChild(textArea);
      return true;
    } catch {
      document.body.removeChild(textArea);
      return false;
    }
  }
};
